package utils

import (
	"testing"
)

func TestEmptyString(t *testing.T) {
	if !IsEmptyString("") {
		t.Errorf("Expected empty string to be empty")
	}
}

func TestEmptyStringWithSpaces(t *testing.T) {
	if !IsEmptyString("        ") {
		t.Errorf("Expected empty string with spaces to be empty")
	}
}

func TestNonEmptyString(t *testing.T) {
	if IsEmptyString("test") {
		t.Errorf("Expected non-empty string to be non-empty")
	}
}

func TestWithMinLength(t *testing.T) {
	if !HaveMinLength("test", 4) {
		t.Errorf("Expected string with minimum length to be with minimum length")
	}
}

func TestWithMinLengthWithSpaces(t *testing.T) {
	if HaveMinLength("      tes      ", 4) {
		t.Errorf("Expected string with minimum length with spaces to be with minimum length")
	}
}

func TestMinLengthWithEmptyString(t *testing.T) {
	if HaveMinLength("               ", 4) {
		t.Errorf("Expected empty string to be with minimum length")
	}
}

func TestMaxLength(t *testing.T) {
	if !HaveMaxLength("test", 4) {
		t.Errorf("Expected string with maximum length to be with maximum length")
	}
}

func TestMaxLengthWithMoreThanCharacters(t *testing.T) {
	if HaveMaxLength("testtesttest", 4) {
		t.Errorf("Expected string with maximum length with more than 4 characters to be with maximum length")
	}
}

func TestMaxLengthWithLessThanMaxCharacters(t *testing.T) {
	if !HaveMaxLength("tes", 4) {
		t.Errorf("Expected string with maximum length with less than 4 characters to be with maximum length")
	}
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(3)
	if err != nil {
		t.Errorf("Expected token to be generated without errors")
	}

	if len(token) != 6 {
		t.Errorf("Expected token to be %d characters long", 6)
	}
}
