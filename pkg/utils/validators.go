package utils

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// ValidateString validate string is not empty, utf8, and normalized utf8.
func ValidateString(s string) error {
	if len(s) == 0 {
		return fmt.Errorf("string cannot be empty")
	}

	if !utf8.ValidString(s) {
		return fmt.Errorf("string must be valid utf8")
	}

	if !norm.NFKC.IsNormalString(s) {
		return fmt.Errorf("string must be normalized")
	}

	return nil
}
