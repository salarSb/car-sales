package common

import (
	"github.com/salarSb/car-sales/config"
	"math/rand"
	"strings"
	"unicode"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func CheckPassword(password string) bool {
	cfg := config.GetConfig()
	if len(password) < cfg.Password.MinLength {
		return false
	}
	if len(password) > cfg.Password.MaxLength {
		return false
	}
	if cfg.Password.IncludeChars && !HasLetter(password) {
		return false
	}
	if cfg.Password.IncludeDigits && !HasDigit(password) {
		return false
	}
	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}
	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}
	return true
}

func GeneratePassword() string {
	var password strings.Builder
	cfg := config.GetConfig()
	passwordLength := cfg.Password.MinLength + 2
	minSpecialChar := 2
	minNum := 3
	if !cfg.Password.IncludeDigits {
		minNum = 0
	}
	minUppercase := 3
	if !cfg.Password.IncludeUppercase {
		minUppercase = 0
	}
	minLowerCase := 3
	if !cfg.Password.IncludeLowercase {
		minLowerCase = 0
	}

	//set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}
	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}
	//Set uppercase
	for i := 0; i < minUppercase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	//Set lowercase
	for i := 0; i < minLowerCase; i++ {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}
	remainingLength := passwordLength - minSpecialChar - minNum - minUppercase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
