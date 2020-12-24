package zapctx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	l := New(nil)
	require.NotNil(t, l)
}
