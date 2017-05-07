package caesar

import (
	"math"
	"strings"
)

const (
	alpha = "abcdefghijklmnopqrstuvwxyz"
	space = byte(' ')
)

// Transpose takes plaintext, and returns ciphertext by shifting the alphabet
// by the given offset.
func Transpose(plaintext []byte, offset int) []byte {
	cipher := Cipher(offset)
	legend := Key(cipher)

	ciphertext := make([]byte, len(plaintext), len(plaintext))

	for i, b := range plaintext {
		c := legend[b]

		if c == 0 {
			c = space
		}

		ciphertext[i] = c
	}

	return ciphertext
}

// Cipher return a new cipher, used for transposing the alphabet into the
// ciphertext.
func Cipher(offset int) []byte {
	return Shift([]byte(alpha), offset)
}

// Key returns a map where the key maps to the transposed characters.
//
// For example, a maps to b, b, maps to c.
func Key(cipher []byte) map[byte]byte {
	key := map[byte]byte{}

	for i, c := range alpha {
		key[byte(c)] = byte(cipher[i])
	}

	return key
}

// Shift shifts an array of bytes by the given offset.
//
// The offset can be positive or negative, and can exceed the length of the
// byte array.
func Shift(b []byte, offset int) []byte {
	// deal with any negative offset
	if offset < 0 {
		offset = len(b) + offset
	}

	downcase := []byte(strings.ToLower(string(b)))

	// modulus of offset/number of characters, to avoid bounds issues
	idx := int(math.Mod(float64(offset), float64(len(downcase))))

	// switch the bytes around
	return append(downcase[idx:len(downcase)], downcase[0:idx]...)
}
