package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"strings"
)

const (
	defaultLength            = 16
	defaultIncludeUpper      = true
	defaultIncludeLower      = true
	defaultIncludeDigits     = true
	defaultIncludeSpecial    = true
	defaultNumberOfPasswords = 1
)

var (
	lowerLetters      = "abcdefghijklmnopqrstuvwxyz"
	upperLetters      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits            = "0123456789"
	specialCharacters = "!@#$%^&*()-_+=~:;?<>,./|\\{}[]`"
	length            int
	includeUpper      bool
	includeLower      bool
	includeDigits     bool
	includeSpecial    bool
	numberOfPasswords int
)

func init() {
	flag.IntVar(&length, "length", defaultLength, "Length of the generated password")
	flag.BoolVar(&includeUpper, "upper", defaultIncludeUpper, "Include uppercase letters")
	flag.BoolVar(&includeLower, "lower", defaultIncludeLower, "Include lowercase letters")
	flag.BoolVar(&includeDigits, "digits", defaultIncludeDigits, "Include digits")
	flag.BoolVar(&includeSpecial, "special", defaultIncludeSpecial, "Include special characters")
	flag.IntVar(&numberOfPasswords, "number", defaultNumberOfPasswords, "Number of passwords to generate")
	flag.Parse()
}

func containsRequiredCharacters(password string, list string) bool {
	for _, char := range password {
		if strings.ContainsRune(list, char) {
			return true
		}
	}
	return false
}

func containsUppercaseCharacter(password string) bool {
	if !includeUpper {
		return true
	}

	return containsRequiredCharacters(password, upperLetters)
}
func containsLowercaseCharacter(password string) bool {
	if !includeLower {
		return true
	}

	return containsRequiredCharacters(password, lowerLetters)
}

func containsDigits(password string) bool {
	if !includeDigits {
		return true
	}

	return containsRequiredCharacters(password, digits)
}

func containsSpecialCharacter(password string) bool {
	if !includeSpecial {
		return true
	}

	return containsRequiredCharacters(password, specialCharacters)
}
func generatePassword() string {
	var passwordChars string
	if includeUpper {
		passwordChars += upperLetters
	}
	if includeLower {
		passwordChars += lowerLetters
	}
	if includeDigits {
		passwordChars += digits
	}
	if includeSpecial {
		passwordChars += specialCharacters
	}

	password := make([]byte, length)
	for i := 0; i < length; i++ {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(passwordChars))))
		if err != nil {
			panic(err)
		}
		password[i] = passwordChars[idx.Int64()]
	}
	return string(password)
}

func main() {
	for numberOfPasswords > 0 {
		for {
			password := generatePassword()
			if containsUppercaseCharacter(password) && containsLowercaseCharacter(password) && containsDigits(password) && containsSpecialCharacter(password) {
				fmt.Println("Generated Password:", password)
				numberOfPasswords--
				break
			}
		}
	}
}
