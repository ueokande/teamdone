package shared

import "math/rand"

var charset = []byte("abcdefghijklmnopqrstuvwxyz234567")

func RandomKey() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
