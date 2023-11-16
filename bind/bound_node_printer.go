package bind

import (
	"fmt"
	"io"
	"strconv"

	"github.com/Ranxy/looper/symbol"
	"github.com/Ranxy/looper/syntax"
	"github.com/Ranxy/looper/write"
)

func WriteTo(w io.Writer, node BoundNode) {
	switch v := w.(type) {
	case *write.Writer:
		writeTo(v, node)
	default:
		writeTo(write.NewWriteIndent(w, false), node)
	}
}

func writeTo(w write.WriteIndent, node BoundNode) {
	switch node.Kind() {
	case BoundNodeKindBlockStatement:
		WriteBlockStatement(w, node.(*BoundBlockStatements))
	case BoundNodeKindVariableDeclaration:
		WriteVariableDeclaration(w, node.(*BoundVariableDeclaration))
	case BoundNodeKindLabelStatement:
		WriteLabelStatement(w, node.(*LabelStatement))
	case BoundNodeKindGotoStatement:
		WriteGotoStatement(w, node.(*GotoStatement))
	case BoundNodeKindConditionalGotoStatement:
		WriteConditionalGotoStatement(w, node.(*ConditionalGotoStatement))
	case BoundNodeKindReturnStatement:
		WriteReturnStatement(w, node.(*BoundReturnStatements))
	case BoundNodeKindExpressionStatement:
		WriteExpressionStatement(w, node.(*BoundExpressStatements))
	case BoundNodeKindErrorExpress:
		WriteErrorExpression(w, node.(*BoundErrorExpression))
	case BoundNodeKindLiteralExpress:
		WriteLiteralExpression(w, node.(*BoundLiteralExpression))
	case BoundNodeKindVariableExpress:
		WriteVariableExpression(w, node.(*BoundVariableExpression))
	case BoundNodeKindAssignmentExpress:
		WriteAssignmentExpression(w, node.(*BoundAssignmentExpression))
	case BoundNodeKindUnaryExpress:
		WriteUnaryExpression(w, node.(*BoundUnaryExpression))
	case BoundNodeKindBinaryExpress:
		WriteBinaryExpression(w, node.(*BoundBinaryExpression))
	case BoundNodeKindCallExpress:
		WriteCallExpression(w, node.(*BoundCallExpression))

	default:
		panic(fmt.Sprintf("Unexcepted BoundNodeKind %d", node.Kind()))
	}
}

func WriteNestedExpression(w write.WriteIndent, parentPrecedence int, expression BoundExpression) {
	switch e := expression.(type) {
	case *BoundUnaryExpression:
		WriteNestedExpressionSon(w, parentPrecedence, syntax.GetUnaryOperatorPrecedence(e.Op.syntaxKind), e)
	case *BoundBinaryExpression:
		WriteNestedExpressionSon(w, parentPrecedence, syntax.GetBinaryOperatorPrecedence(e.Op.SyntaxKind), e)
	default:
		WriteTo(w, e)
	}
}
func WriteNestedExpressionSon(w write.WriteIndent, parentPrecedence int, currentPrecedence int, expression BoundExpression) {
	needsParenthesis := parentPrecedence >= currentPrecedence

	if needsParenthesis {
		write.WritePunctuation(w, syntax.SyntaxKindOpenParenthesisToken)
	}
	WriteTo(w, expression)
	if needsParenthesis {
		write.WritePunctuation(w, syntax.SyntaxKindCloseParenthesisToken)
	}
}

func WriteBlockStatement(w write.WriteIndent, node *BoundBlockStatements) {
	write.WritePunctuation(w, syntax.SyntaxKindOpenBraceToken)
	write.WriteLine(w)
	w.IndentChange(1)
	for _, v := range node.Statement {
		WriteTo(w, v)
	}
	w.IndentChange(-1)
	write.WritePunctuation(w, syntax.SyntaxKindCloseBraceToken)
	write.WriteLine(w)
}

func WriteVariableDeclaration(w write.WriteIndent, node *BoundVariableDeclaration) {
	if node.Variable.IsReadOnly() {
		write.WriteKeyword(w, syntax.SyntaxKindLetKeywords)
	} else {
		write.WriteKeyword(w, syntax.SyntaxKindVarKeywords)
	}
	write.WriteSpace(w)
	write.WriteIdentifier(w, node.Variable.GetName())
	write.WriteSpace(w)
	write.WritePunctuation(w, syntax.SyntaxKindEqualToken)
	write.WriteSpace(w)
	WriteTo(w, node.Initializer)
	write.WriteLine(w)
}

func WriteLabelStatement(w write.WriteIndent, node *LabelStatement) {
	unindent := w.Indent() > 0
	if unindent {
		w.IndentChange(-1)
	}
	write.WriteString(w, node.Label.Name)
	write.WritePunctuation(w, syntax.SyntaxKindColon)
	write.WriteLine(w)

	if unindent {
		w.IndentChange(1)
	}
}

func WriteGotoStatement(w write.WriteIndent, node *GotoStatement) {
	write.WriteKeywordStr(w, "goto ")
	write.WriteIdentifier(w, node.Label.Name)
	write.WriteLine(w)
}

func WriteConditionalGotoStatement(w write.WriteIndent, node *ConditionalGotoStatement) {
	write.WriteKeywordStr(w, "goto ")
	write.WriteIdentifier(w, node.Label.Name)
	if node.JumpIfFalse {
		write.WriteKeywordStr(w, " if not ")
	} else {
		write.WriteKeywordStr(w, " if ")
	}
	WriteTo(w, node.Condition)
	write.WriteLine(w)
}

func WriteReturnStatement(w write.WriteIndent, node *BoundReturnStatements) {
	write.WriteKeyword(w, syntax.SyntaxKindReturnKeywords)
	if node.Express != nil {
		write.WriteSpace(w)
		WriteTo(w, node.Express)
	}
	write.WriteLine(w)
}
func WriteExpressionStatement(w write.WriteIndent, node *BoundExpressStatements) {
	WriteTo(w, node.Express)
	write.WriteLine(w)
}

func WriteErrorExpression(w write.WriteIndent, node *BoundErrorExpression) {
	write.WriteKeywordStr(w, "?")
}

func WriteLiteralExpression(w write.WriteIndent, node *BoundLiteralExpression) {
	switch node.Type() {
	case symbol.TypeInt:
		write.WriteString(w, strconv.FormatInt(node.Value.(int64), 10))
	case symbol.TypeBool:
		write.WriteString(w, strconv.FormatBool(node.Value.(bool)))
	case symbol.TypeString:
		write.WriteString(w, "\""+node.Value.(string)+"\"")
	default:
		panic("Unexcepted Type")
	}
}

func WriteVariableExpression(w write.WriteIndent, node *BoundVariableExpression) {
	write.WriteIdentifier(w, node.Variable.GetName())
}

func WriteAssignmentExpression(w write.WriteIndent, node *BoundAssignmentExpression) {
	write.WriteIdentifier(w, node.Variable.GetName())
	write.WriteSpace(w)
	write.WritePunctuation(w, syntax.SyntaxKindEqualToken)
	write.WriteSpace(w)
	WriteTo(w, node.Express)
}

func WriteUnaryExpression(w write.WriteIndent, node *BoundUnaryExpression) {
	precedence := syntax.GetBinaryOperatorPrecedence(node.Op.syntaxKind)
	write.WritePunctuation(w, node.Op.syntaxKind)
	WriteNestedExpression(w, precedence, node.Operand)
}

func WriteBinaryExpression(w write.WriteIndent, node *BoundBinaryExpression) {
	precedence := syntax.GetBinaryOperatorPrecedence(node.Op.SyntaxKind)

	WriteNestedExpression(w, precedence, node.Left)
	write.WriteSpace(w)
	write.WritePunctuation(w, node.Op.SyntaxKind)
	write.WriteSpace(w)
	WriteNestedExpression(w, precedence, node.Right)

}

func WriteCallExpression(w write.WriteIndent, node *BoundCallExpression) {
	write.WriteIdentifier(w, node.Function.GetName())
	write.WritePunctuation(w, syntax.SyntaxKindOpenParenthesisToken)
	isFirst := true

	for _, arg := range node.Arguments {
		if isFirst {
			isFirst = false
		} else {
			write.WritePunctuation(w, syntax.SyntaxKindCommaToken)
			write.WriteSpace(w)
		}
		WriteTo(w, arg)
	}
	write.WritePunctuation(w, syntax.SyntaxKindCloseParenthesisToken)
}
