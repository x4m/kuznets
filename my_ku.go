package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	var s = "64a59400000000000000000000000000"
	slice, _ := hex.DecodeString(s)
	var fixed [16]byte
	copy(fixed[:], slice)
	fixed = L(fixed)
	copy(slice,fixed[:])
	subst := hex.EncodeToString(slice)
	fmt.Print("L(" + s + ") = " + subst + "\n")
}
