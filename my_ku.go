package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var s = "8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"
	slice, _ := hex.DecodeString(s)
	var fixed [32]byte
	copy(fixed[:], slice)
	fixed1 := Keys(fixed)[9]
	copy(slice,fixed1[:])
	subst := hex.EncodeToString(slice)
	fmt.Print("L(" + s + ") = " + subst + "\n")
}
