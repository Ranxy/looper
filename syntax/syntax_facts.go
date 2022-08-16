package syntax

func GetUnaryOperatorPrecedence(kind SyntaxKind) int {
	switch kind {
	case SyntaxKindPlusToken, SyntaxKindMinusToken, SyntaxKindBangToken:
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
	case SyntaxKindEqualEqualToken, SyntaxKindBangEqualToken:
		return 3
	case SyntaxKindAmpersandAmpersandToken:
		return 2
	case SyntaxKindPipePileToken:
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
	default:
		return SyntaxKindIdentifierToken
	}
}
