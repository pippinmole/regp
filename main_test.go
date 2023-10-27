package main

import "testing"

func TestInvalidPatternInMatchThrowsException(t *testing.T) {
	invalidPattern := "a(b"

	invalidErr := match(invalidPattern, []string{
		"abc",
	})

	if invalidErr == nil {
		t.Errorf("Expected error for invalid pattern")
	}
}

func TestNoPatternInMatchThrowsNoException(t *testing.T) {
	noPattern := ""

	err := match(noPattern, []string{
		"abc",
	})

	if err != nil {
		t.Errorf("Expected no error for no pattern")
	}
}

func TestNoStringsInMatchThrowsNoException(t *testing.T) {
	noStrings := "abc"

	err := match(noStrings, []string{})

	if err != nil {
		t.Errorf("Expected no error for no strings")
	}
}
