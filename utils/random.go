package utils

import (
	b64 "encoding/base64"
	"math/rand"
)

// Returns an int >= min, < max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of Aa-Zz 0-9 chars with len = l
func RandomString(length int) string {
	var char = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = char[rand.Intn(len(char)-1)]
	}
	return (string(bytes))
}

// Generate a random interger string 0-9 chars with len = l
func RandomIntString(length int) string {
	var char = "0123456789"
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = char[rand.Intn(len(char)-1)]
	}
	return (string(bytes))
}

func RandomBase64String(size int, prefix string) string {
	// Here's the `string` we'll encode/decode.
	data := RandomString(size)

	// Go supports both standard and URL-compatible
	// base64. Here's how to encode using the standard
	// encoder. The encoder requires a `[]byte` so we
	// cast our `string` to that type.
	uEnc := b64.RawURLEncoding.EncodeToString([]byte(data))
	// fmt.Println(sEnc)

	// Decoding may return an error, which you can check
	// if you don't already know the input to be
	// well-formed.
	// uDec, _ := b64.URLEncoding.DecodeString(sEnc)
	// fmt.Println(string(sDec))
	// fmt.Println()

	return prefix + uEnc
}
