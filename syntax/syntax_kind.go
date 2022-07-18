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
	SyntaxKindEqualToken              //=
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
	SyntaxKindNameExpress
	SyntaxKindUnaryExpress
	SyntaxKindBinaryExpress
	SyntaxKindParenthesizedExpress
	SyntaxKindAssignmentExpress
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
	SyntaxKindEqualToken:              "SyntaxKindEqualToken",
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
	SyntaxKindNameExpress:          "SyntaxKindNameExpress",
	SyntaxKindUnaryExpress:         "SyntaxKindUnaryExpress",
	SyntaxKindBinaryExpress:        "SyntaxKindBinaryExpress",
	SyntaxKindParenthesizedExpress: "SyntaxKindParenthesizedExpress",
	SyntaxKindAssignmentExpress:    "SyntaxKindAssignmentExpress",
}

func (k SyntaxKind) Text() string {
	return syntaxKindTextMap[k]
}

func ListTextSyntaxKind() []SyntaxKind {
	res := make([]SyntaxKind, 0, len(syntaxKindTextMap))
	for k := range syntaxKindTextMap {
		res = append(res, k)
	}
	return res
}

var syntaxKindTextMap = map[SyntaxKind]string{
	SyntaxKindPlusToken:               "+",
	SyntaxKindMinusToken:              "-",
	SyntaxKindStarToken:               "*",
	SyntaxKindSlashToken:              "/",
	SyntaxKindBangToken:               "!",
	SyntaxKindEqualToken:              "=",
	SyntaxKindAmpersandAmpersandToken: "&&",
	SyntaxKindPipePileToken:           "||",
	SyntaxKindEqualEqualToken:         "==",
	SyntaxKindBangEqualToken:          "!=",
	SyntaxKindOpenParenthesisToken:    "(",
	SyntaxKindCloseParenthesisToken:   ")",

	SyntaxKindTrueKeywords:  "true",
	SyntaxKindFalseKeywords: "false",
}
