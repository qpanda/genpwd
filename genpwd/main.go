package main

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
)

const (
	LowerAlphabetic   = "abcdefghijklmnopqrstuvwxyz"
	UpperAlphabetic   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numeric           = "0123456789"
	AllSpecial        = " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	RestrictedSpecial = "#$%&*+,-./:;=?@^_|~"
	LimitedSpecial    = "#$%&*+-/:=?@_|~"

	Alphabetic   = LowerAlphabetic + UpperAlphabetic
	Alphanumeric = Alphabetic + Numeric
)

func SpecialCharSet(specialCharSet string) (string, error) {
	if specialCharSet == "all" {
		return AllSpecial, nil
	} else if specialCharSet == "restricted" {
		return RestrictedSpecial, nil
	} else if specialCharSet == "limited" {
		return LimitedSpecial, nil
	} else {
		return "", errors.New("Invalid special character set")
	}
}

func RandomUint(max uint) uint {
	maxUint := ^uint(0)
	bytes := make([]byte, 8)
	rangeEnd := maxUint - (maxUint % max)
	for do := true; do; do = uint(binary.BigEndian.Uint64(bytes)) > rangeEnd {
		if _, err := rand.Read(bytes); err != nil {
			panic(err)
		}
	}

	return uint(binary.BigEndian.Uint64(bytes)) % max
}

func RandomChar(charSet string) byte {
	return charSet[RandomUint(uint(len(charSet)))]
}

func RandomPasswordChar(bytes []byte, charSet string) (uint, byte) {
	index := RandomUint(uint(len(bytes)))
	startIndex := index
	for bytes[index] != 0 {
		index = ((index + 1) % uint(len(bytes)))
		if index == startIndex {
			panic(errors.New("All characters are set"))
		}
	}

	return index, RandomChar(charSet)
}

func RandomPassword(length uint, minLowerChars uint, minUpperChars uint, minNumericChars uint, minSpecialChars uint, specialCharSet string) (string, error) {
	if length < 5 || length > 1024 {
		return "", errors.New("Invalid password length: Password length must be between 5 and 1024 characters")
	}

	if (minLowerChars + minUpperChars + minNumericChars + minSpecialChars) > length {
		return "", errors.New("Invalid password complexity rules: Number of lower, upper, numeric, and special characters must not exceed password length")
	}

	specialChars, err := SpecialCharSet(specialCharSet)
	if err != nil {
		return "", err
	}

	passwordCharSet := Alphanumeric + specialChars

	bytes := make([]byte, length)

	for i := uint(0); i < minLowerChars; i++ {
		index, char := RandomPasswordChar(bytes, LowerAlphabetic)
		bytes[index] = char
	}

	for i := uint(0); i < minUpperChars; i++ {
		index, char := RandomPasswordChar(bytes, UpperAlphabetic)
		bytes[index] = char
	}

	for i := uint(0); i < minNumericChars; i++ {
		index, char := RandomPasswordChar(bytes, Numeric)
		bytes[index] = char
	}

	for i := uint(0); i < minSpecialChars; i++ {
		index, char := RandomPasswordChar(bytes, specialChars)
		bytes[index] = char
	}

	for i := uint(0); i < length; i++ {
		if bytes[i] == 0 {
			bytes[i] = RandomChar(passwordCharSet)
		}
	}

	return string(bytes), nil
}

func main() {
	length := flag.Uint("l", 16, "password length")
	minLowerChars := flag.Uint("lower", 2, "minimum number of lower case characters")
	minUpperChars := flag.Uint("upper", 2, "minimum number of upper case characters")
	minNumericChars := flag.Uint("numeric", 2, "minimum number of numeric characters")
	minSpecialChars := flag.Uint("special", 2, "minimum number of special characters")
	specialCharSet := flag.String("specialCharSet", "limited", "special character set \"all\", \"restricted\", or \"limited\"")

	flag.Parse()

	password, err := RandomPassword(*length, *minLowerChars, *minUpperChars, *minNumericChars, *minSpecialChars, *specialCharSet)
	if err != nil {
		panic(err)
	}

	fmt.Println(password)
}
