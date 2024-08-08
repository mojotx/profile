package fnord

import (
	"fmt"
	"testing"
)

func TestGetRandomRune(t *testing.T) {
	r := GetRandomRune()
	if r < 'A' || r > 'Z' {
		t.Errorf("Expected rune between 'A' and 'Z', got %q", r)
	}
}
func TestGetRandomString(t *testing.T) {
	length := 10
	s := GetRandomString(length)
	if len(s) != length {
		t.Errorf("Expected string of length %d, got %d", length, len(s))
	}
}
func TestStringContainsName(t *testing.T) {

	s := fmt.Sprintf("Hello, %s!", magicWord)
	if !StringContainsName(s) {
		t.Errorf("Expected string to contain magic word, but it didn't")
	}

	if !StringRegexName(s) {
		t.Errorf("Expected string to contain magic word, but it didn't")
	}

	s = "Goodbye"
	if StringContainsName(s) {
		t.Errorf("Expected string to not contain magic word, but it did")
	}

	if StringRegexName(s) {
		t.Errorf("Expected string to not contain magic word, but it did")
	}
}

func BenchmarkStringRegex(b *testing.B) {
	s := fmt.Sprintf("Hello, %s!", magicWord)
	for i := 0; i < b.N; i++ {
		StringRegexName(s)
	}
}

func BenchmarkStringContains(b *testing.B) {
	s := fmt.Sprintf("Hello, %s!", magicWord)
	for i := 0; i < b.N; i++ {
		StringContainsName(s)
	}
}
