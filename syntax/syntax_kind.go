package syntax

type SyntaxKind uint8

const (
	SyntaxKindBadToken SyntaxKind = iota
	SyntaxKindEofToken
	SyntaxKindWhiteSpaceToken
	SyntaxKindNumberToken             // 12
	SyntaxKindStringToken             //"abc"
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
	SyntaxKindLessToken               // <
	SyntaxKindLessEqualToken          // <=
	SyntaxKindGreatToken              // >
	SyntaxKindGreatEqualToken         //>=
	SyntaxKindOpenParenthesisToken    //(
	SyntaxKindCloseParenthesisToken   //)
	SyntaxKindOpenBraceToken          //{
	SyntaxKindCloseBraceToken         //}
	SyntaxKindSemicolon               // ;
	SyntaxKindIdentifierToken

	//bitwise operators
	SyntaxKindTildeToken     // ~
	SyntaxKindHatToken       // ^
	SyntaxKindAmpersandToken // &
	SyntaxKindPipeToken      // |

	SyntaxKindTrueKeywords  // True
	SyntaxKindFalseKeywords //False
	SyntaxKindLetKeywords
	SyntaxKindVarKeywords
	SyntaxKindIfKeywords
	SyntaxKindElseKeywords
	SyntaxKindWhileKeywords
	SyntaxkindForKeywords

	//Nodes
	SyntaxKindCompilationUnit

	//Statement
	SyntaxKindBlockStatement
	SyntaxKindVariableDeclaration
	SyntaxKindIfStatement
	SyntaxKindWhileStatement
	SyntaxkindForStatement
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
	SyntaxKindLessToken:               "SyntaxKindLessToken",
	SyntaxKindLessEqualToken:          "SyntaxKindLessEqualToken",
	SyntaxKindGreatToken:              "SyntaxKindGreatToken",
	SyntaxKindGreatEqualToken:         "SyntaxKindGreatEqualToken",
	SyntaxKindOpenParenthesisToken:    "OpenParenthesisToken",
	SyntaxKindCloseParenthesisToken:   "CloseParenthesisToken",
	SyntaxKindOpenBraceToken:          "OpenBraceToken",
	SyntaxKindCloseBraceToken:         "CloseBraceToken",
	SyntaxKindSemicolon:               "Semicolon",
	SyntaxKindIdentifierToken:         "IdentifierToken",

	SyntaxKindTildeToken:     "TildeToken",
	SyntaxKindHatToken:       "HatToken",
	SyntaxKindAmpersandToken: "AmpersandToken",
	SyntaxKindPipeToken:      "PipeToken",

	SyntaxKindTrueKeywords:  "TrueKeywords",
	SyntaxKindFalseKeywords: "FalseKeywords",
	SyntaxKindLetKeywords:   "LetKeywords",
	SyntaxKindVarKeywords:   "VarKeywords",
	SyntaxKindIfKeywords:    "IfKeywords",
	SyntaxKindElseKeywords:  "ElseKeywords",
	SyntaxKindWhileKeywords: "WhileKeywords",
	SyntaxkindForKeywords:   "ForKeywords",

	SyntaxKindCompilationUnit: "CompilationUnit",

	SyntaxKindBlockStatement:      "BlockStatement",
	SyntaxKindVariableDeclaration: "VariableDeclaration",
	SyntaxKindIfStatement:         "IfStatement",
	SyntaxKindWhileStatement:      "WhileStatement",
	SyntaxkindForStatement:        "ForStatement",
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
	SyntaxKindLessToken:               "<",
	SyntaxKindLessEqualToken:          "<=",
	SyntaxKindGreatToken:              ">",
	SyntaxKindGreatEqualToken:         ">=",
	SyntaxKindOpenParenthesisToken:    "(",
	SyntaxKindCloseParenthesisToken:   ")",
	SyntaxKindOpenBraceToken:          "{",
	SyntaxKindCloseBraceToken:         "}",
	SyntaxKindSemicolon:               ";",

	SyntaxKindTildeToken:     "~",
	SyntaxKindHatToken:       "^",
	SyntaxKindAmpersandToken: "&",
	SyntaxKindPipeToken:      "|",

	SyntaxKindTrueKeywords:  "true",
	SyntaxKindFalseKeywords: "false",
	SyntaxKindLetKeywords:   "let",
	SyntaxKindVarKeywords:   "var",
	SyntaxKindIfKeywords:    "if",
	SyntaxKindElseKeywords:  "else",
	SyntaxKindWhileKeywords: "while",
	SyntaxkindForKeywords:   "for",
}
