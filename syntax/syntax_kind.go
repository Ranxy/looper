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
	SyntaxKindColon                   // :
	SyntaxKindCommaToken              //,
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
	SyntaxKindBreakKeywords
	SyntaxKindContinueKeywords
	SyntaxkindForKeywords
	SyntaxkindFunctionKeywords
	SyntaxKindReturnKeywords

	//Nodes
	SyntaxKindCompilationUnit
	SyntaxKindGlobalStatement
	SyntaxKindFunctionDeclaration
	SyntaxKindTypeClause
	SyntaxKindParameter

	//Statement
	SyntaxKindBlockStatement
	SyntaxKindVariableDeclaration
	SyntaxKindIfStatement
	SyntaxKindWhileStatement
	SyntaxkindForStatement
	SyntaxKindBreakStatement
	SyntaxKindContinueStatement
	SyntaxKindReturnStatement
	SyntaxKindExpressStatement

	SyntaxKindLiteralExpress
	SyntaxKindNameExpress
	SyntaxKindUnaryExpress
	SyntaxKindBinaryExpress
	SyntaxKindParenthesizedExpress
	SyntaxKindAssignmentExpress
	SyntaxKindCallExpress
	SyntaxKindUnitExpress
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
	SyntaxKindColon:                   "Colon",
	SyntaxKindCommaToken:              "CommaToken",
	SyntaxKindIdentifierToken:         "IdentifierToken",

	SyntaxKindTildeToken:     "TildeToken",
	SyntaxKindHatToken:       "HatToken",
	SyntaxKindAmpersandToken: "AmpersandToken",
	SyntaxKindPipeToken:      "PipeToken",

	SyntaxKindTrueKeywords:     "TrueKeywords",
	SyntaxKindFalseKeywords:    "FalseKeywords",
	SyntaxKindLetKeywords:      "LetKeywords",
	SyntaxKindVarKeywords:      "VarKeywords",
	SyntaxKindIfKeywords:       "IfKeywords",
	SyntaxKindElseKeywords:     "ElseKeywords",
	SyntaxKindWhileKeywords:    "WhileKeywords",
	SyntaxKindBreakKeywords:    "BreakKeywords",
	SyntaxKindContinueKeywords: "ContinueKeywords",
	SyntaxkindForKeywords:      "ForKeywords",
	SyntaxkindFunctionKeywords: "FunctionKeywords",
	SyntaxKindReturnKeywords:   "ReturnKeywords",

	SyntaxKindCompilationUnit:     "CompilationUnit",
	SyntaxKindFunctionDeclaration: "FunctionDeclaration",
	SyntaxKindGlobalStatement:     "GlobalStatement",
	SyntaxKindTypeClause:          "TypeClause",
	SyntaxKindParameter:           "Parameter",

	SyntaxKindBlockStatement:      "BlockStatement",
	SyntaxKindVariableDeclaration: "VariableDeclaration",
	SyntaxKindIfStatement:         "IfStatement",
	SyntaxKindWhileStatement:      "WhileStatement",
	SyntaxkindForStatement:        "ForStatement",
	SyntaxKindBreakStatement:      "BreakStatement",
	SyntaxKindContinueStatement:   "ContinueStatement",
	SyntaxKindExpressStatement:    "ExpressStatement",
	SyntaxKindReturnStatement:     "ReturnStatement",

	SyntaxKindLiteralExpress:       "LiteralExpress",
	SyntaxKindNameExpress:          "NameExpress",
	SyntaxKindUnaryExpress:         "UnaryExpress",
	SyntaxKindBinaryExpress:        "BinaryExpress",
	SyntaxKindParenthesizedExpress: "ParenthesizedExpress",
	SyntaxKindAssignmentExpress:    "AssignmentExpress",
	SyntaxKindCallExpress:          "CallExpress",
	SyntaxKindUnitExpress:          "UnitExpress",
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
	SyntaxKindColon:                   ":",
	SyntaxKindCommaToken:              ",",

	SyntaxKindTildeToken:     "~",
	SyntaxKindHatToken:       "^",
	SyntaxKindAmpersandToken: "&",
	SyntaxKindPipeToken:      "|",

	SyntaxKindTrueKeywords:     "true",
	SyntaxKindFalseKeywords:    "false",
	SyntaxKindLetKeywords:      "let",
	SyntaxKindVarKeywords:      "var",
	SyntaxKindIfKeywords:       "if",
	SyntaxKindElseKeywords:     "else",
	SyntaxKindWhileKeywords:    "while",
	SyntaxkindForKeywords:      "for",
	SyntaxKindBreakKeywords:    "break",
	SyntaxKindContinueKeywords: "continue",
	SyntaxkindFunctionKeywords: "fn",
	SyntaxKindReturnKeywords:   "return",
}
