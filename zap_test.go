package zapctx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	l := New(nil)
	require.NotNil(t, l)
}
