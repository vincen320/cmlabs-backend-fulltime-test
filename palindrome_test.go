package cmlabsbackendfulltimetest

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPalindrome(s string) bool {
	length := len(s)
	first_char := 0
	last_char := length - 1

	s = strings.ToLower(s)

	for first_char < last_char {
		for !(s[first_char] >= 97 && s[first_char] <= 122) {
			first_char++
		}
		for !(s[last_char] >= 97 && s[last_char] <= 122) {
			last_char--
		}
		if s[first_char] != s[last_char] {
			return false
		}
		first_char++
		last_char--
	}
	return true
}

func TestPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("SAIPPUAKIVIKAUPPIAS"))
	assert.True(t, isPalindrome("Aibohphobia"))
	assert.True(t, isPalindrome("Anna"))
	assert.True(t, isPalindrome("Civic"))
	assert.True(t, isPalindrome("My gym"))
	assert.True(t, isPalindrome("No lemon, no melon"))

	assert.False(t, isPalindrome("No lemon, mo melon"))
	assert.False(t, isPalindrome("My gim"))
	assert.False(t, isPalindrome("Civcc"))
	assert.False(t, isPalindrome("Anma"))
	assert.False(t, isPalindrome("Aibobphobia"))
	assert.False(t, isPalindrome("SAIPPUAKIVJKAUPPIAS"))
}
