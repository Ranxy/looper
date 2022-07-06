package syntax

type SyntaxKind uint8

const (
	SyntaxKindBadToken SyntaxKind = iota
	SyntaxKindEofToken
	SyntaxKindWhiteSpaceToken
	SyntaxKindNumberToken
	SyntaxKindPlusToken
	SyntaxKindMinusToken
	SyntaxKindStarToken
	SyntaxKindSlashToken
	SyntaxKindOpenParenthesisToken
	SyntaxKindCloseParenthesisToken

	SyntaxKindNumberExpress
	SyntaxKindBinaryExpress
	SyntaxKindParenthesizedExpress
)

func (k SyntaxKind) String() string {
	if str, has := syntaxKindKeyMap[k]; has {
		return str
	} else {
		return "UnexpectedToken"
	}
}

var syntaxKindKeyMap = map[SyntaxKind]string{
	SyntaxKindBadToken:              "SyntaxKindBadToken",
	SyntaxKindEofToken:              "SyntaxKindEofToken",
	SyntaxKindWhiteSpaceToken:       "SyntaxKindWhiteSpaceToken",
	SyntaxKindNumberToken:           "SyntaxKindNumberToken",
	SyntaxKindPlusToken:             "SyntaxKindPluxToken",
	SyntaxKindMinusToken:            "SyntaxKindMinusToken",
	SyntaxKindStarToken:             "SyntaxKindStarToken",
	SyntaxKindSlashToken:            "SyntaxKindSlashToken",
	SyntaxKindOpenParenthesisToken:  "SyntaxKindOpenParenthesisToken",
	SyntaxKindCloseParenthesisToken: "SyntaxKindCloseParenthesisToken",

	SyntaxKindNumberExpress:        "SyntaxKindNumberExpress",
	SyntaxKindBinaryExpress:        "SyntaxKindBinaryExpress",
	SyntaxKindParenthesizedExpress: "SyntaxKindParenthesizedExpress",
}
