package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func BenchmarkKuznets(b *testing.B) {
	var key = "8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"
	keySlice, _ := hex.DecodeString(key)
	ku := NewKuznetsReader(randomReader{}, keySlice)

	buffer := make([]byte, 16)
	fmt.Println(b.N)
	for n := 0; n < b.N; n++ {
		ku.Read(buffer)
	}
}
