package syntax

func GetUnaryOperatorPrecedence(kind SyntaxKind) int {
	switch kind {
	case SyntaxKindPlusToken, SyntaxKindMinusToken:
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
	default:
		return 0
	}
}
