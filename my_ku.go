package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main() {
	var key = "8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"
	keySlice, _ := hex.DecodeString(key)
	var data = "1122334455667700ffeeddccbbaa9988"
	dataSlice, _ := hex.DecodeString(data)
	ku := NewKuznets(bytes.NewReader(dataSlice), keySlice)
	cyphered := make([]byte, 16)
	ku.Read(cyphered)
	fmt.Print("Ku(" + data + ") = " + hex.EncodeToString(cyphered) + "\n")
}
