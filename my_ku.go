package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"
)

func main() {
	var key = "8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"
	keySlice, _ := hex.DecodeString(key)
	var data = "1122334455667700ffeeddccbbaa9988"
	dataSlice, _ := hex.DecodeString(data)
	ku := NewKuznets(bytes.NewReader(dataSlice), keySlice)
	cyphered := make([]byte, 16)
	ku.Read(cyphered)
	fmt.Println("Ku(" + data + ") = " + hex.EncodeToString(cyphered) + "\n")

	ku = NewKuznets(randomReader{}, keySlice)

	start := time.Now()

	for i := 0; i < 1024*1024/16; i++ {
		ku.Read(cyphered)
	}
	elapsed := time.Since(start)

	fmt.Println("Time to encrypt 1Mb ", elapsed)
}
