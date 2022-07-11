package syntax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSyntaxKind(t *testing.T) {
	for _, k := range ListTextSyntaxKind() {
		text := k.Text()
		if text == "" {
			break
		}
		tokenList := ParseTokens(text)
		require.Len(t, tokenList, 1)
		token := tokenList[0]
		require.Equal(t, k, token.Kind())
		require.Equal(t, k.Text(), token.Text)
	}
}
