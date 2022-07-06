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
	expr := p.ParserExpress(0)
	eof := p.MatchToken(SyntaxKindEofToken)

	return NewSyntaxTree(expr, eof, p.errors)
}

func (p *Parser) ParserExpress(parentPrecedence int) Express {
	var left Express
	unaryOperatorPrecedence := GetUnaryOperatorPrecedence(p.Current().Kind())
	if unaryOperatorPrecedence != 0 && unaryOperatorPrecedence >= parentPrecedence {
		operator := p.NextToken()
		operand := p.ParserExpress(unaryOperatorPrecedence)
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
		right := p.ParserExpress(precedence)
		left = NewBinaryExpress(left, operator, right)
	}
	return left
}

func (p *Parser) parsePrimaryExpress() Express {
	if p.Current().Kind() == SyntaxKindOpenParenthesisToken {
		open := p.NextToken()
		expr := p.ParserExpress(0)
		right := p.MatchToken(SyntaxKindCloseParenthesisToken)

		return NewParenthesisExpress(open, expr, right)
	}
	number := p.MatchToken(SyntaxKindNumberToken)
	return NewLiteralExpress(number)
}
