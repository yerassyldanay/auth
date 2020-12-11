package randomer

import (
"math/rand"
"strings"
"time"
)

// variables needed
const lowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"
const upperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const allLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const allDigits = "0123456789"

// helps generate random different numbers
// without it numbers will not be random
func init() {
	rand.Seed(time.Now().UnixNano())
}

// generate random integer
func RandomInt(minv, maxv int64) int64 {
	return minv + rand.Int63n(maxv - minv + 1)
}

// generate lower case letters
func RandomOnlyLowerCaseString(n int) string {
	var sb strings.Builder
	k := len(lowerCaseLetters)

	for i := 0; i < n; i++ {
		c := lowerCaseLetters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

/*
	RandomOnlyUpperCaseLetters
		generate upper case letters
*/
func RandomOnlyUpperCaseLetters(n int) string {
	var sb strings.Builder
	k := len(lowerCaseLetters)

	for i := 0; i < n; i++ {
		c := upperCaseLetters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(allLetters)

	for i := 0; i < n; i++ {
		c := allLetters[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
