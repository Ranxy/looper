package syntax

func GetUnaryOperatorPrecedence(kind SyntaxKind) int {
	switch kind {
	case SyntaxKindPlusToken, SyntaxKindMinusToken, SyntaxKindBangToken, SyntaxKindTildeToken:
		return 6
	default:
		return 0
	}
}
func GetBinaryOperatorPrecedence(kind SyntaxKind) int {
	switch kind {
	case SyntaxKindStarToken, SyntaxKindSlashToken:
		return 5
	case SyntaxKindPlusToken, SyntaxKindMinusToken:
		return 4
	case SyntaxKindEqualEqualToken, SyntaxKindBangEqualToken,
		SyntaxKindLessToken, SyntaxKindLessEqualToken,
		SyntaxKindGreatToken, SyntaxKindGreatEqualToken:
		return 3
	case SyntaxKindAmpersandAmpersandToken, SyntaxKindAmpersandToken:
		return 2
	case SyntaxKindPipePileToken, SyntaxKindPipeToken, SyntaxKindHatToken:
		return 1
	default:
		return 0
	}
}

func GetKeyWordsKind(text string) SyntaxKind {
	switch text {
	case "true":
		return SyntaxKindTrueKeywords
	case "false":
		return SyntaxKindFalseKeywords
	case "let":
		return SyntaxKindLetKeywords
	case "var":
		return SyntaxKindVarKeywords
	case "if":
		return SyntaxKindIfKeywords
	case "else":
		return SyntaxKindElseKeywords
	case "while":
		return SyntaxKindWhileKeywords
	case "for":
		return SyntaxkindForKeywords
	default:
		return SyntaxKindIdentifierToken
	}
}
