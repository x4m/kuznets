package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"time"
	"unsafe"
)

func main() {
	var key = "8899aabbccddeeff0011223344556677fedcba98765432100123456789abcdef"
	keySlice, _ := hex.DecodeString(key)
	var data = "1122334455667700ffeeddccbbaa9988"
	dataSlice, _ := hex.DecodeString(data)
	ku := NewKuznetsReader(bytes.NewReader(dataSlice), keySlice)
	cyphered := make([]byte, 16)
	ku.Read(cyphered)
	fmt.Println("Ku(" + data + ") = " + hex.EncodeToString(cyphered) + "\n")
	//return

	kuz := NewKuznets(keySlice)

	start := time.Now()

	for i := 0; i < 100*1024*1024/16; i++ {
		var data [16]byte
		*(*int64)(unsafe.Pointer(&data[0])) = int64(i);//rand.Int63n(math.MaxInt64)
		*(*int64)(unsafe.Pointer(&data[8])) = int64(i);//rand.Int63n(math.MaxInt64)
		kuz.EncryptBlock(data)
	}
	elapsed := time.Since(start)

	fmt.Println("Time to encrypt 100Mb ", elapsed)
}
