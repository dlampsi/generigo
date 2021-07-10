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

func Test_JsonMapToString(t *testing.T) {
	f := func(m map[string]string, noerr bool, expected string) {
		s, err := JsonMapToString(m)
		if noerr {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
		require.Equal(t, s, expected)
	}

	f(nil, true, "")
	f(map[string]string{"one": "two"}, true, `{"one":"two"}`)
}

func Test_StringToJsonMap(t *testing.T) {
	f := func(s string, noerr bool, expected map[string]string) {
		m, err := StringToJsonMap(s)
		if noerr {
			require.NoError(t, err)
		} else {
			require.Error(t, err)
		}
		require.Equal(t, m, expected)
	}

	f("", true, map[string]string{})
	f("string", false, nil)
	f(`{"one":"two"}`, true, map[string]string{"one": "two"})
}
