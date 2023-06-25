package syntax

import (
	"fmt"

	"github.com/Ranxy/looper/diagnostic"
	"github.com/Ranxy/looper/texts"
)

type Parser struct {
	text        *texts.TextSource
	tokenList   []SyntaxToken
	len         int
	diagnostics *diagnostic.DiagnosticBag
	pos         int
}

func NewParser(text *texts.TextSource) *Parser {
	tokenList := make([]SyntaxToken, 0)
	lex := NewLexer(text)

	var token SyntaxToken

	for token.Kind() != SyntaxKindEofToken {
		token = lex.NextToken()
		if token.Kind() != SyntaxKindWhiteSpaceToken && token.Kind() != SyntaxKindBadToken {
			tokenList = append(tokenList, token)
		}
	}

	return &Parser{
		text:        text,
		tokenList:   tokenList,
		len:         len(tokenList),
		diagnostics: diagnostic.MergeDiagnostics(lex.diagnostics),
		pos:         0,
	}
}

func (p *Parser) Peek(offset int) SyntaxToken {
	idx := p.pos + offset
	if idx >= p.len {
		return p.tokenList[p.len-1]
	}
	return p.tokenList[idx]
}

func (p *Parser) Current() SyntaxToken {
	return p.Peek(0)
}

func (p *Parser) NextToken() SyntaxToken {
	current := p.Current()
	p.pos += 1
	return current
}

func (p *Parser) MatchToken(kind SyntaxKind) SyntaxToken {
	if p.Current().Kind() == kind {
		return p.NextToken()
	}
	p.diagnostics.Report(p.Current().Span(), fmt.Sprintf("Unexpected token %v, expected %v", p.Current().Kind(), kind))
	return SyntaxToken{kind, p.Current().Position, "", nil, true}
}

func (p *Parser) ParseComplitionUnit() *CompliationUnit {
	var statement = p.ParseMembers()
	eofToken := p.MatchToken(SyntaxKindEofToken)
	return NewCompliationUnit(statement, eofToken)
}

func (p *Parser) ParseStatement() Statement {
	switch p.Current().Kind() {
	case SyntaxKindOpenBraceToken:
		return p.ParseBlockStatement()
	case SyntaxKindLetKeywords:
		fallthrough
	case SyntaxKindVarKeywords:
		return p.ParseVariableDeclaration()
	case SyntaxKindIfKeywords:
		return p.ParseIfStatement()
	case SyntaxKindWhileKeywords:
		return p.ParseWhileStatement()
	case SyntaxkindForKeywords:
		return p.ParseForStatement()
	case SyntaxKindBreakKeywords:
		return p.ParseBreakStatement()
	case SyntaxKindContinueKeywords:
		return p.ParseContinueStatement()
	case SyntaxkindFunctionKeywords:
		p.diagnostics.Report(p.Current().Span(), "Function must be defined at top level")
		return p.ParseFunctionDeclaration()
	default:
		return p.ParseExpressStatement()
	}
}

func (p *Parser) ParseBlockStatement() *BlockStatement {
	var statements = make([]Statement, 0)
	openBraceToken := p.MatchToken(SyntaxKindOpenBraceToken)
	for p.Current().Kind() != SyntaxKindEofToken && p.Current().Kind() != SyntaxKindCloseBraceToken {
		start := p.Current()
		statement := p.ParseStatement()
		statements = append(statements, statement)
		if start == p.Current() {
			p.NextToken()
		}
	}
	closeBraceToken := p.MatchToken(SyntaxKindCloseBraceToken)

	return NewBlockStatement(openBraceToken, statements, closeBraceToken)
}

func (p *Parser) ParseVariableDeclaration() Statement {
	var exprected SyntaxKind
	if p.Current().Kind() == SyntaxKindLetKeywords {
		exprected = SyntaxKindLetKeywords
	} else {
		exprected = SyntaxKindVarKeywords
	}

	keyword := p.MatchToken(exprected)
	identifier := p.MatchToken(SyntaxKindIdentifierToken)
	optionTypeClause := p.ParseOptionTypeClause()
	equals := p.MatchToken(SyntaxKindEqualToken)
	initializer := p.ParserExpress()

	return NewVariableDeclarationSyntax(keyword, identifier, optionTypeClause, equals, initializer)
}

func (p *Parser) ParseForStatement() Statement {
	keywords := p.MatchToken(SyntaxkindForKeywords)
	initCondition := p.ParseStatement()
	firstSemicolon := p.MatchToken(SyntaxKindSemicolon)
	endCondtion := p.ParserExpress()
	secondSemicolon := p.MatchToken(SyntaxKindSemicolon)
	updateCondtion := p.ParseStatement()
	statement := p.ParseStatement()

	return NewForStatement(keywords, initCondition, firstSemicolon, endCondtion, secondSemicolon, updateCondtion, statement)
}

func (p *Parser) ParseWhileStatement() Statement {
	keywords := p.MatchToken(SyntaxKindWhileKeywords)
	condition := p.ParserExpress()
	statement := p.ParseStatement()

	return NewWhileStatement(keywords, condition, statement)
}

func (p *Parser) ParseBreakStatement() Statement {
	keywords := p.MatchToken(SyntaxKindBreakKeywords)
	return NewBreakStatement(keywords)
}

func (p *Parser) ParseContinueStatement() Statement {
	keywords := p.MatchToken(SyntaxKindContinueKeywords)
	return NewContinueStatement(keywords)
}

func (p *Parser) ParseIfStatement() Statement {
	keywords := p.MatchToken(SyntaxKindIfKeywords)
	condition := p.ParserExpress()
	statement := p.ParseStatement()
	elseClause := p.ParseElseClause()

	return NewIfStatement(keywords, condition, statement, elseClause)
}

func (p *Parser) ParseElseClause() *ElseClauseSyntax {
	if p.Current().Kind() != SyntaxKindElseKeywords {
		return nil
	}
	keywords := p.NextToken()
	statement := p.ParseStatement()
	return NewElseClause(keywords, statement)
}

func (p *Parser) ParseExpressStatement() *ExpressStatement {
	express := p.ParserExpress()
	return NewExpressStatement(express)
}

func (p *Parser) ParserExpress() Express {
	return p.ParserAssignmentExpress()
}

func (p *Parser) ParserAssignmentExpress() Express {
	if p.Peek(0).Kind() == SyntaxKindIdentifierToken && p.Peek(1).Kind() == SyntaxKindEqualToken {
		identifierToken := p.NextToken()
		operator := p.NextToken()
		right := p.ParserAssignmentExpress()

		return NewAssignmentExpress(identifierToken, operator, right)
	}

	return p.ParserBinaryExpress(0)
}

func (p *Parser) ParserBinaryExpress(parentPrecedence int) Express {
	var left Express
	unaryOperatorPrecedence := GetUnaryOperatorPrecedence(p.Current().Kind())
	if unaryOperatorPrecedence != 0 && unaryOperatorPrecedence >= parentPrecedence {
		operator := p.NextToken()
		operand := p.ParserBinaryExpress(unaryOperatorPrecedence)
		left = NewUnaryExpress(operator, operand)
	} else {
		left = p.parsePrimaryExpress()
	}

	for {
		precedence := GetBinaryOperatorPrecedence(p.Current().Kind())

		if precedence == 0 || precedence <= parentPrecedence {
			break
		}
		operator := p.NextToken()
		right := p.ParserBinaryExpress(precedence)
		left = NewBinaryExpress(left, operator, right)
	}
	return left
}

func (p *Parser) parsePrimaryExpress() Express {
	switch p.Current().Kind() {
	case SyntaxKindOpenParenthesisToken:
		{
			open := p.NextToken()
			expr := p.ParserBinaryExpress(0)
			right := p.MatchToken(SyntaxKindCloseParenthesisToken)

			return NewParenthesisExpress(open, expr, right)
		}
	case SyntaxKindTrueKeywords, SyntaxKindFalseKeywords:
		{
			keywords := p.NextToken()
			value := keywords.Kind() == SyntaxKindTrueKeywords
			return NewLiteralValueExpress(keywords, value)
		}
	case SyntaxKindNumberToken:
		{
			number := p.MatchToken(SyntaxKindNumberToken)
			return NewLiteralExpress(number)
		}

	case SyntaxKindStringToken:
		{
			str := p.MatchToken(SyntaxKindStringToken)
			return NewLiteralExpress(str)
		}
	case SyntaxKindIdentifierToken:
		{
			if p.Current().Kind() == SyntaxKindIdentifierToken && p.Peek(1).Kind() == SyntaxKindOpenParenthesisToken {
				return p.ParseCallExpress()
			} else {
				return p.ParserNameExpress()
			}
		}
	default:
		{
			identifierToken := p.MatchToken(SyntaxKindIdentifierToken)
			return NewNameExpress(identifierToken)
		}
	}

}

func (p *Parser) ParserNameExpress() Express {
	identifierToken := p.NextToken()
	return NewNameExpress(identifierToken)
}

func (p *Parser) ParseCallExpress() Express {
	identifierToken := p.NextToken()
	open := p.MatchToken(SyntaxKindOpenParenthesisToken)
	params := p.ParseParams()
	close := p.MatchToken(SyntaxKindCloseParenthesisToken)

	return NewCallExpress(identifierToken, open, *params, close)
}

func (p *Parser) ParseParams() *SeparatedList {
	nodeList := make([]SyntaxNode, 0)
	for p.Current().Kind() != SyntaxKindCloseParenthesisToken && p.Current().Kind() != SyntaxKindEofToken {
		express := p.ParserExpress()

		nodeList = append(nodeList, express)

		if p.Current().Kind() != SyntaxKindCloseParenthesisToken {
			current := p.Current()
			comma := p.MatchToken(SyntaxKindCommaToken)
			if current == p.Current() {
				p.NextToken()
			}
			nodeList = append(nodeList, comma)
		}
	}
	return NewSeptartedList(nodeList)
}

//parameters

func (p *Parser) ParseOptionTypeClause() *TypeClauseSyntax {
	if p.Current().Kind() != SyntaxKindColon {
		return nil
	}
	return p.ParseTypeClause()
}

func (p *Parser) ParseTypeClause() *TypeClauseSyntax {
	colonToken := p.MatchToken(SyntaxKindColon)
	identifier := p.MatchToken(SyntaxKindIdentifierToken)
	return NewTypeClauseSyntax(colonToken, identifier)
}

func (p *Parser) ParseParameter() *ParameterSyntax {
	identifier := p.MatchToken(SyntaxKindIdentifierToken)
	tp := p.ParseTypeClause()
	return NewParameterSyntax(identifier, tp)
}

func (p *Parser) ParseParameterList() *SeparatedList {
	nodeList := make([]SyntaxNode, 0)
	for p.Current().Kind() != SyntaxKindCloseParenthesisToken && p.Current().Kind() != SyntaxKindEofToken {
		current := p.Current()
		parameter := p.ParseParameter()
		nodeList = append(nodeList, parameter)

		if p.Current().Kind() != SyntaxKindCloseParenthesisToken {
			comma := p.MatchToken(SyntaxKindCommaToken)
			nodeList = append(nodeList, comma)
		}
		if p.Current() == current {
			p.NextToken()
		}
	}
	return NewSeptartedList(nodeList)
}

func (p *Parser) ParseFunctionDeclaration() MemberSyntax {
	functionKeyword := p.MatchToken(SyntaxkindFunctionKeywords)
	identifier := p.MatchToken(SyntaxKindIdentifierToken)
	openParenthesis := p.MatchToken(SyntaxKindOpenParenthesisToken)
	parameters := p.ParseParameterList()
	closeParenthesis := p.MatchToken(SyntaxKindCloseParenthesisToken)
	tp := p.ParseOptionTypeClause()
	body := p.ParseBlockStatement()
	return NewFunctionDeclaration(functionKeyword, identifier, openParenthesis, parameters, closeParenthesis, tp, body)
}

func (p *Parser) ParseMembers() []MemberSyntax {
	members := make([]MemberSyntax, 0)
	for p.Current().Kind() != SyntaxKindEofToken {
		startToken := p.Current()
		member := p.ParseMember()
		members = append(members, member)

		// If ParseMember() did not consume any tokens,
		// we need to skip the current token and continue
		// in order to avoid an infinite loop.
		//
		// We don't need to report an error, because we'll
		// already tried to parse an expression statement
		// and reported one.
		if p.Current() == startToken {
			p.NextToken()
		}
	}
	return members
}

func (p *Parser) ParseMember() MemberSyntax {
	if p.Current().Kind() == SyntaxkindFunctionKeywords {
		return p.ParseFunctionDeclaration()
	}
	return p.ParseGlobalStatement()
}

func (p *Parser) ParseGlobalStatement() MemberSyntax {
	statement := p.ParseStatement()
	return NewGlobalStatement(statement)
}
