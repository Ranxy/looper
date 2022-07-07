package syntax

type SyntaxKind uint8

const (
	SyntaxKindBadToken SyntaxKind = iota
	SyntaxKindEofToken
	SyntaxKindWhiteSpaceToken
	SyntaxKindNumberToken             // 12
	SyntaxKindPlusToken               //+
	SyntaxKindMinusToken              //-
	SyntaxKindStarToken               //*
	SyntaxKindSlashToken              // /
	SyntaxKindBangToken               //!
	SyntaxKindAmpersandAmpersandToken //&&
	SyntaxKindPipePileToken           // ||
	SyntaxKindEqualEqualToken         // ==
	SyntaxKindBangEqualToken          // !=
	SyntaxKindOpenParenthesisToken    //(
	SyntaxKindCloseParenthesisToken   //)
	SyntaxKindIdentifierToken

	SyntaxKindTrueKeywords  // True
	SyntaxKindFalseKeywords //False

	SyntaxKindLiteralExpress
	SyntaxKindUnaryExpress
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
	SyntaxKindBadToken:                "SyntaxKindBadToken",
	SyntaxKindEofToken:                "SyntaxKindEofToken",
	SyntaxKindWhiteSpaceToken:         "SyntaxKindWhiteSpaceToken",
	SyntaxKindNumberToken:             "SyntaxKindNumberToken",
	SyntaxKindPlusToken:               "SyntaxKindPluxToken",
	SyntaxKindMinusToken:              "SyntaxKindMinusToken",
	SyntaxKindStarToken:               "SyntaxKindStarToken",
	SyntaxKindSlashToken:              "SyntaxKindSlashToken",
	SyntaxKindBangToken:               "SyntaxKindBangToken",
	SyntaxKindAmpersandAmpersandToken: "SyntaxKindAmpersandAmpersandToken",
	SyntaxKindPipePileToken:           "SyntaxKindPipePileToken",
	SyntaxKindEqualEqualToken:         "SyntaxKindEqualEqualToken",
	SyntaxKindBangEqualToken:          "SyntaxKindBangEqualToken",
	SyntaxKindOpenParenthesisToken:    "SyntaxKindOpenParenthesisToken",
	SyntaxKindCloseParenthesisToken:   "SyntaxKindCloseParenthesisToken",
	SyntaxKindIdentifierToken:         "SyntaxKindIdentifierToken",

	SyntaxKindTrueKeywords:  "SyntaxKindTrueKeywords",
	SyntaxKindFalseKeywords: "SyntaxKindFalseKeywords",

	SyntaxKindLiteralExpress:       "SyntaxKindLiteralExpress",
	SyntaxKindUnaryExpress:         "SyntaxKindUnaryExpress",
	SyntaxKindBinaryExpress:        "SyntaxKindBinaryExpress",
	SyntaxKindParenthesizedExpress: "SyntaxKindParenthesizedExpress",
}
