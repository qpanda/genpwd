package main

import (
	"strings"
	"testing"
)

func TestSpecialCharSet_All(t *testing.T) {
	if allSpecialCharSet, _ := SpecialCharSet("all"); allSpecialCharSet != AllSpecial {
		t.Errorf("Special char set invalid: expected %v, actual %v", allSpecialCharSet, AllSpecial)
	}
}
func TestSpecialCharSet_Restricted(t *testing.T) {
	if restrictedSpecialCharSet, _ := SpecialCharSet("restricted"); restrictedSpecialCharSet != RestrictedSpecial {
		t.Errorf("Special char set invalid: expected %v, actual %v", restrictedSpecialCharSet, RestrictedSpecial)
	}
}

func TestSpecialCharSet_Limited(t *testing.T) {
	if limitedSpecialCharSet, _ := SpecialCharSet("limited"); limitedSpecialCharSet != LimitedSpecial {
		t.Errorf("Special char set invalid: expected %v, actual %v", limitedSpecialCharSet, LimitedSpecial)
	}
}

func TestSpecialCharSet_Invalid(t *testing.T) {
	if _, err := SpecialCharSet("invalid"); err == nil {
		t.Errorf("Expected error for invalid special char set")
	}
}

func TestRandomUint(t *testing.T) {
	max := uint(123)
	if random := RandomUint(max); random > max {
		t.Errorf("Expected random number between 0 and %d but actual number is %d", max, random)
	}
}

func TestRandomChar_Numeric(t *testing.T) {
	if random := RandomChar(Numeric); !strings.Contains(Numeric, string(random)) {
		t.Errorf("Expected one of the characters %s but actual character is %c", Numeric, random)
	}
}

func TestRandomChar_Alphanumeric(t *testing.T) {
	if random := RandomChar(Alphanumeric); !strings.Contains(Alphanumeric, string(random)) {
		t.Errorf("Expected one of the characters %s but actual character is %c", Alphanumeric, random)
	}
}

func TestRandomPasswordChar_NoCharAvailable(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	RandomPasswordChar([]byte("ABC"), UpperAlphabetic)
}

func TestRandomPasswordChar_LastCharAvailable(t *testing.T) {
	bytes := make([]byte, 3)
	bytes[0] = 'A'
	bytes[1] = 'B'

	if index, _ := RandomPasswordChar(bytes, UpperAlphabetic); index != 2 {
		t.Errorf("Invalid index: expected %d, actual %d", 2, index)
	}
}

func TestRandomPasswordChar_FirstCharAvailable(t *testing.T) {
	bytes := make([]byte, 3)
	bytes[1] = 'B'
	bytes[2] = 'C'

	if index, _ := RandomPasswordChar(bytes, UpperAlphabetic); index != 0 {
		t.Errorf("Invalid index: expected %d, actual %d", 0, index)
	}
}

func TestRandomPasswordChar_MiddleCharAvailable(t *testing.T) {
	bytes := make([]byte, 3)
	bytes[0] = 'A'
	bytes[2] = 'C'

	if index, _ := RandomPasswordChar(bytes, UpperAlphabetic); index != 1 {
		t.Errorf("Invalid index: expected %d, actual %d", 1, index)
	}
}

func TestRandomPasswordChar_Sparse(t *testing.T) {
	bytes := make([]byte, 10)
	bytes[0] = 'A'
	bytes[2] = 'B'
	bytes[4] = 'C'

	if index, _ := RandomPasswordChar(bytes, UpperAlphabetic); index == 0 || index == 2 || index == 4 {
		t.Errorf("Invalid index: expected index not to be 0, 2, and 4, actual %d", index)
	}
}

func TestRandomPassword_InvalidLength(t *testing.T) {
	if _, err := RandomPassword(1, 0, 0, 0, 0, "all"); err == nil {
		t.Errorf("Expected error for invalid length")
	}
}

func TestRandomPassword_InvalidPasswordComplexityRules(t *testing.T) {
	if _, err := RandomPassword(6, 2, 2, 2, 2, "limited"); err == nil {
		t.Errorf("Expected error for invalid password complexity rules")
	}
}

func TestRandomPassword_WithoutPasswordComplexityRules(t *testing.T) {
	if password, _ := RandomPassword(8, 0, 0, 0, 0, "restricted"); len(password) != 8 {
		t.Errorf("Expected password with length 8")
	}
}

func TestRandomPassword_WithPasswordComplexityRules(t *testing.T) {
	minLowerChars := uint(5)
	minUpperChars := uint(5)
	minNumericChars := uint(5)
	minSpecialChars := uint(5)
	password, _ := RandomPassword(100, minLowerChars, minUpperChars, minNumericChars, minSpecialChars, "limited")

	if lowerChars := CountAll(password, LowerAlphabetic); lowerChars < minLowerChars {
		t.Errorf("Incorrect number of lower characters, expected at least %d but found only %d", minLowerChars, lowerChars)
	}

	if upperChars := CountAll(password, UpperAlphabetic); upperChars < minUpperChars {
		t.Errorf("Incorrect number of upper characters, expected at least %d but found only %d", minUpperChars, upperChars)
	}

	if numericChars := CountAll(password, Numeric); numericChars < minNumericChars {
		t.Errorf("Incorrect number of numeric characters, expected at least %d but found only %d", minUpperChars, minNumericChars)
	}

	if specialChars := CountAll(password, LimitedSpecial); specialChars < minSpecialChars {
		t.Errorf("Incorrect number of special characters, expected at least %d but found only %d", minUpperChars, minNumericChars)
	}
}

func TestRandomPassword_WithStrictPasswordComplexityRules(t *testing.T) {
	minLowerChars := uint(1)
	minUpperChars := uint(2)
	minNumericChars := uint(3)
	minSpecialChars := uint(4)
	password, _ := RandomPassword(10, minLowerChars, minUpperChars, minNumericChars, minSpecialChars, "limited")

	if lowerChars := CountAll(password, LowerAlphabetic); lowerChars < minLowerChars {
		t.Errorf("Incorrect number of lower characters, expected at least %d but found only %d", minLowerChars, lowerChars)
	}

	if upperChars := CountAll(password, UpperAlphabetic); upperChars < minUpperChars {
		t.Errorf("Incorrect number of upper characters, expected at least %d but found only %d", minUpperChars, upperChars)
	}

	if numericChars := CountAll(password, Numeric); numericChars < minNumericChars {
		t.Errorf("Incorrect number of numeric characters, expected at least %d but found only %d", minUpperChars, minNumericChars)
	}

	if specialChars := CountAll(password, LimitedSpecial); specialChars < minSpecialChars {
		t.Errorf("Incorrect number of special characters, expected at least %d but found only %d", minUpperChars, minNumericChars)
	}
}

func CountAll(s string, substr string) uint {
	count := 0
	for i := 0; i < len(substr); i++ {
		count += strings.Count(s, string(substr[i]))
	}

	return uint(count)
}
