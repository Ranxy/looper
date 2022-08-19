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
	SyntaxKindOpenBraceToken          //{
	SyntaxKindCloseBraceToken         //}
	SyntaxKindIdentifierToken

	SyntaxKindTrueKeywords  // True
	SyntaxKindFalseKeywords //False
	SyntaxKindLetKeywords
	SyntaxKindVarKeywords
	SyntaxKindIfKeywords
	SyntaxKindElseKeywords

	//Nodes
	SyntaxKindCompilationUnit

	//Statement
	SyntaxKindBlockStatement
	SyntaxKindVariableDeclaration
	SyntaxKindIfStatement
	SyntaxKindExpressStatement

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
	SyntaxKindBadToken:                "BadToken",
	SyntaxKindEofToken:                "EofToken",
	SyntaxKindWhiteSpaceToken:         "WhiteSpaceToken",
	SyntaxKindNumberToken:             "NumberToken",
	SyntaxKindPlusToken:               "PluxToken",
	SyntaxKindMinusToken:              "MinusToken",
	SyntaxKindStarToken:               "StarToken",
	SyntaxKindSlashToken:              "SlashToken",
	SyntaxKindBangToken:               "BangToken",
	SyntaxKindEqualToken:              "EqualToken",
	SyntaxKindAmpersandAmpersandToken: "AmpersandAmpersandToken",
	SyntaxKindPipePileToken:           "PipePileToken",
	SyntaxKindEqualEqualToken:         "EqualEqualToken",
	SyntaxKindBangEqualToken:          "BangEqualToken",
	SyntaxKindOpenParenthesisToken:    "OpenParenthesisToken",
	SyntaxKindCloseParenthesisToken:   "CloseParenthesisToken",
	SyntaxKindOpenBraceToken:          "OpenBraceToken",
	SyntaxKindCloseBraceToken:         "CloseBraceToken",
	SyntaxKindIdentifierToken:         "IdentifierToken",

	SyntaxKindTrueKeywords:  "TrueKeywords",
	SyntaxKindFalseKeywords: "FalseKeywords",
	SyntaxKindLetKeywords:   "LetKeywords",
	SyntaxKindVarKeywords:   "VarKeywords",
	SyntaxKindIfKeywords:    "IfKeywords",
	SyntaxKindElseKeywords:  "ElseKeywords",

	SyntaxKindCompilationUnit: "CompilationUnit",

	SyntaxKindBlockStatement:      "BlockStatement",
	SyntaxKindVariableDeclaration: "VariableDeclaration",
	SyntaxKindExpressStatement:    "ExpressStatement",

	SyntaxKindLiteralExpress:       "LiteralExpress",
	SyntaxKindNameExpress:          "NameExpress",
	SyntaxKindUnaryExpress:         "UnaryExpress",
	SyntaxKindBinaryExpress:        "BinaryExpress",
	SyntaxKindParenthesizedExpress: "ParenthesizedExpress",
	SyntaxKindAssignmentExpress:    "AssignmentExpress",
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
	SyntaxKindOpenBraceToken:          "{",
	SyntaxKindCloseBraceToken:         "}",

	SyntaxKindTrueKeywords:  "true",
	SyntaxKindFalseKeywords: "false",
	SyntaxKindLetKeywords:   "let",
	SyntaxKindVarKeywords:   "var",
	SyntaxKindIfKeywords:    "if",
	SyntaxKindElseKeywords:  "else",
}
