package generigo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CopyMap(t *testing.T) {
	f := func(sourceMap map[string]interface{}) {
		t.Helper()
		newMap := CopyMap(sourceMap)
		require.Equal(t, sourceMap, newMap)
	}

	f(map[string]interface{}{"one": 1, "two": "three"})
}

func Test_CopyMapString(t *testing.T) {
	f := func(sourceMap map[string]string) {
		t.Helper()
		newMap := CopyMapString(sourceMap)
		require.Equal(t, sourceMap, newMap)
	}

	f(map[string]string{"one": "1", "two": "three"})
	f(map[string]string{})
}
