package fuzzing_test

import (
	"testing"
	"unicode/utf8"

	"github.com/strider2038/go-examples/testing/fuzzing"
)

func FuzzBasicOverwriteString(f *testing.F) {
	f.Fuzz(func(t *testing.T, str string, value rune, n int) {
		fuzzing.OverwriteString(str, value, n)
	})
}

func FuzzOverwriteStringSuffix(f *testing.F) {
	f.Add("Hello, world!", 'A', 15)

	f.Fuzz(func(t *testing.T, str string, value rune, n int) {
		result := fuzzing.OverwriteString(str, value, n)
		if n > 0 && n < utf8.RuneCountInString(str) {
			// If we modified characters [0:n], then characters [n:] should stay the same
			resultSuffix := string([]rune(result)[n:])
			strSuffix := string([]rune(str)[:])
			if resultSuffix != strSuffix {
				t.Fatalf("OverwriteString modified too many characters! Expected %s, got %s.", strSuffix, resultSuffix)
			}
		}
	})
}
