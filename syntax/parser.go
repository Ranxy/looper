package syntax

import "fmt"

type Parser struct {
	tokenList []SyntaxToken
	len       int
	errors    []string
	pos       int
}

func NewParser(text string) *Parser {
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
		tokenList: tokenList,
		len:       len(tokenList),
		errors:    lex.errors,
		pos:       0,
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
	p.errors = append(p.errors, fmt.Sprintf("Unexpected token %v, expected %v", p.Current().Kind(), kind))
	return SyntaxToken{kind, p.Current().Position, "", nil}
}

func (p *Parser) Parse() *SyntaxTree {
	expr := p.ParserExpress()
	eof := p.MatchToken(SyntaxKindEofToken)

	return NewSyntaxTree(expr, eof, p.errors)
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
	case SyntaxKindIdentifierToken:
		{
			identifierToken := p.NextToken()
			return NewNameExpress(identifierToken)
		}
	default:
		{
			number := p.MatchToken(SyntaxKindNumberToken)
			return NewLiteralExpress(number)
		}
	}

}
