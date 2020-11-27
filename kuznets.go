package main

import "C"
import (
	"io"
	"math/rand"
	"unsafe"
)

var Sbox = [0x100]byte{
	252, 238, 221, 17, 207, 110, 49, 22, 251, 196, 250, 218, 35, 197, 4, 77, 233, 119, 240, 219, 147, 46, 153, 186, 23, 54, 241, 187, 20, 205, 95, 193, 249, 24, 101, 90, 226, 92, 239, 33, 129, 28, 60, 66, 139, 1, 142, 79, 5, 132, 2, 174, 227, 106, 143, 160, 6, 11, 237, 152, 127, 212, 211, 31, 235, 52, 44, 81, 234, 200, 72, 171, 242, 42, 104, 162, 253, 58, 206, 204, 181, 112, 14, 86, 8, 12, 118, 18, 191, 114, 19, 71, 156, 183, 93, 135, 21, 161, 150, 41, 16, 123, 154, 199, 243, 145, 120, 111, 157, 158, 178, 177, 50, 117, 25, 61, 255, 53, 138, 126, 109, 84, 198, 128, 195, 189, 13, 87, 223, 245, 36, 169, 62, 168, 67, 201, 215, 121, 214, 246, 124, 34, 185, 3, 224, 15, 236, 222, 122, 148, 176, 188, 220, 232, 40, 80, 78, 51, 10, 74, 167, 151, 96, 115, 30, 0, 98, 68, 26, 184, 56, 130, 100, 159, 38, 65, 173, 69, 70, 146, 39, 94, 85, 47, 140, 163, 165, 125, 105, 213, 149, 59, 7, 88, 179, 64, 134, 172, 29, 247, 48, 55, 107, 228, 136, 217, 231, 137, 225, 27, 131, 73, 76, 63, 248, 254, 141, 83, 170, 144, 202, 216, 133, 97, 32, 113, 103, 164, 45, 43, 9, 91, 203, 155, 37, 208, 190, 229, 108, 82, 89, 166, 116, 210, 230, 244, 180, 192, 209, 102, 175, 194, 57, 75, 99, 182}

func xx(res *[16]byte, b [16]byte) {
	*(*uint64)(unsafe.Pointer(&res[0])) = *(*uint64)(unsafe.Pointer(&res[0])) ^ *(*uint64)(unsafe.Pointer(&b[0]))
	*(*uint64)(unsafe.Pointer(&res[8])) = *(*uint64)(unsafe.Pointer(&res[8])) ^ *(*uint64)(unsafe.Pointer(&b[8]))
}

func s(a *[16]byte) {
	a[0] = Sbox[a[0]]
	a[1] = Sbox[a[1]]
	a[2] = Sbox[a[2]]
	a[3] = Sbox[a[3]]
	a[4] = Sbox[a[4]]
	a[5] = Sbox[a[5]]
	a[6] = Sbox[a[6]]
	a[7] = Sbox[a[7]]
	a[8] = Sbox[a[8]]
	a[9] = Sbox[a[9]]
	a[10] = Sbox[a[10]]
	a[11] = Sbox[a[11]]
	a[12] = Sbox[a[12]]
	a[13] = Sbox[a[13]]
	a[14] = Sbox[a[14]]
	a[15] = Sbox[a[15]]
}

var ls_array [16][256][16]byte

func ls(res *[16]byte) {
	a := *res
	*res = ls_array[0][a[0]]
	for i := 1; i < 16; i++ {
		b := ls_array[i][a[i]]
		*(*uint64)(unsafe.Pointer(&res[0])) = *(*uint64)(unsafe.Pointer(&res[0])) ^ *(*uint64)(unsafe.Pointer(&b[0]))
		*(*uint64)(unsafe.Pointer(&res[8])) = *(*uint64)(unsafe.Pointer(&res[8])) ^ *(*uint64)(unsafe.Pointer(&b[8]))
	}
}

var gfm_array [256][256]byte

func init() {
	for i := 0; i < 256; i++ {
		for o := 0; o < 256; o++ {
			gfm_array[i][o] = gfm_manual(byte(i), byte(o))
		}
	}

	for iteration := 0; iteration < 16; iteration++ {
		for i := 0; i < 256; i++ {
			var bytes [16]byte
			bytes[iteration] = Sbox[i]
			l(&bytes)
			ls_array[iteration][i] = bytes
		}
	}
}

func gfm(x byte, y byte) byte {
	return gfm_array[x][y]
}

func gfm_manual(x byte, y byte) byte {
	z := byte(0)
	for y != 0 {
		if y&1 != 0 {
			z ^= x
		}
		mult := byte(0xC3)
		if x&0x80 == 0 {
			mult = 0
		}
		x = (x << 1) ^ (mult)
		y >>= 1
	}

	return z
}

var gfc = [16]byte{148, 32, 133, 16, 194, 192, 1, 251, 1, 192, 194, 16, 133, 32, 148, 1}

func lv128v8(a *[16]byte) byte {
	return gfm(148, a[0]) ^
		gfm(32, a[1]) ^
		gfm(133, a[2]) ^
		gfm(16, a[3]) ^
		gfm(194, a[4]) ^
		gfm(192, a[5]) ^
		a[6] ^
		gfm(251, a[7]) ^
		a[8] ^
		gfm(192, a[9]) ^
		gfm(194, a[10]) ^
		gfm(16, a[11]) ^
		gfm(133, a[12]) ^
		gfm(32, a[13]) ^
		gfm(148, a[14]) ^
		a[15]
}

func r(a *[16]byte) {
	var a15 = lv128v8(a)

	a[15] = a[14]
	a[14] = a[13]
	a[13] = a[12]
	a[12] = a[11]
	a[11] = a[10]
	a[10] = a[9]
	a[9] = a[8]
	a[8] = a[7]
	a[7] = a[6]
	a[6] = a[5]
	a[5] = a[4]
	a[4] = a[3]
	a[3] = a[2]
	a[2] = a[1]
	a[1] = a[0]
	a[0] = a15
}

func l(a *[16]byte) {
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
	r(a)
}

func keys(master [32]byte) (keys [10][16]byte) {
	copy(keys[0][:], master[:16])
	copy(keys[1][:], master[16:])

	var x [16]byte
	var y [16]byte
	var z [16]byte
	x = keys[0]
	y = keys[1]

	for i := byte(1); i <= 32; i++ {
		var c [16]byte
		c[15] = i

		l(&c)
		z = x
		xx(&z, c)
		s(&z)
		l(&z)
		xx(&z, y)
		y = x
		x = z

		if (i % 8) == 0 {
			keys[i>>2] = x
			keys[i>>2+1] = y
		}
	}
	return
}

type Kuznets struct {
	internal io.Reader
	keys     [10][16]byte
}

func NewKuznets(masterKey []byte) Kuznets {
	var fixed [32]byte
	copy(fixed[:], masterKey)
	return Kuznets{nil, keys(fixed)}
}

func NewKuznetsReader(reader io.Reader, masterKey []byte) io.Reader {
	var fixed [32]byte
	copy(fixed[:], masterKey)
	return Kuznets{reader, keys(fixed)}
}

func (k Kuznets) Read(p []byte) (n int, err error) {
	if len(p) != 16 {
		panic("TODO")
	}
	var b [16]byte
	read, _ := k.internal.Read(p)
	if read != 16 {
		panic("TODO")
	}
	copy(b[:], p)

	b = k.EncryptBlock(b)

	copy(p, b[:])
	return 16, nil
}

func (k Kuznets) EncryptBlock(b [16]byte) [16]byte {
	for i := 0; i < 9; i++ {
		*(*uint64)(unsafe.Pointer(&b[0])) = *(*uint64)(unsafe.Pointer(&b[0])) ^ *(*uint64)(unsafe.Pointer(&k.keys[i][0]))
		*(*uint64)(unsafe.Pointer(&b[8])) = *(*uint64)(unsafe.Pointer(&b[8])) ^ *(*uint64)(unsafe.Pointer(&k.keys[i][8]))

		b1 := ls_array[0][b[0]]
		b2 := ls_array[1][b[1]]
		b3 := ls_array[2][b[2]]
		b4 := ls_array[3][b[3]]
		b5 := ls_array[4][b[4]]
		b6 := ls_array[5][b[5]]
		b7 := ls_array[6][b[6]]
		b8 := ls_array[7][b[7]]
		b9 := ls_array[8][b[8]]
		b10 := ls_array[9][b[9]]
		b11 := ls_array[10][b[10]]
		b12 := ls_array[11][b[11]]
		b13 := ls_array[12][b[12]]
		b14 := ls_array[13][b[13]]
		b15 := ls_array[14][b[14]]
		b16 := ls_array[15][b[15]]
		*(*uint64)(unsafe.Pointer(&b[0])) = *(*uint64)(unsafe.Pointer(&b1[0])) ^
			*(*uint64)(unsafe.Pointer(&b2[0])) ^
			*(*uint64)(unsafe.Pointer(&b3[0])) ^
			*(*uint64)(unsafe.Pointer(&b4[0])) ^
			*(*uint64)(unsafe.Pointer(&b5[0])) ^
			*(*uint64)(unsafe.Pointer(&b6[0])) ^
			*(*uint64)(unsafe.Pointer(&b7[0])) ^
			*(*uint64)(unsafe.Pointer(&b8[0])) ^
			*(*uint64)(unsafe.Pointer(&b9[0])) ^
			*(*uint64)(unsafe.Pointer(&b10[0])) ^
			*(*uint64)(unsafe.Pointer(&b11[0])) ^
			*(*uint64)(unsafe.Pointer(&b12[0])) ^
			*(*uint64)(unsafe.Pointer(&b13[0])) ^
			*(*uint64)(unsafe.Pointer(&b14[0])) ^
			*(*uint64)(unsafe.Pointer(&b15[0])) ^
			*(*uint64)(unsafe.Pointer(&b16[0]))

		*(*uint64)(unsafe.Pointer(&b[8])) = *(*uint64)(unsafe.Pointer(&b1[8])) ^
			*(*uint64)(unsafe.Pointer(&b2[8])) ^
			*(*uint64)(unsafe.Pointer(&b3[8])) ^
			*(*uint64)(unsafe.Pointer(&b4[8])) ^
			*(*uint64)(unsafe.Pointer(&b5[8])) ^
			*(*uint64)(unsafe.Pointer(&b6[8])) ^
			*(*uint64)(unsafe.Pointer(&b7[8])) ^
			*(*uint64)(unsafe.Pointer(&b8[8])) ^
			*(*uint64)(unsafe.Pointer(&b9[8])) ^
			*(*uint64)(unsafe.Pointer(&b10[8])) ^
			*(*uint64)(unsafe.Pointer(&b11[8])) ^
			*(*uint64)(unsafe.Pointer(&b12[8])) ^
			*(*uint64)(unsafe.Pointer(&b13[8])) ^
			*(*uint64)(unsafe.Pointer(&b14[8])) ^
			*(*uint64)(unsafe.Pointer(&b15[8])) ^
			*(*uint64)(unsafe.Pointer(&b16[8]))

		// At this point I really wish GO had macro functions
	}
	*(*uint64)(unsafe.Pointer(&b[0])) = *(*uint64)(unsafe.Pointer(&b[0])) ^ *(*uint64)(unsafe.Pointer(&k.keys[9][0]))
	*(*uint64)(unsafe.Pointer(&b[8])) = *(*uint64)(unsafe.Pointer(&b[8])) ^ *(*uint64)(unsafe.Pointer(&k.keys[9][8]))
	return b
}

func (k Kuznets) EncryptBlockRef(b *[16]byte) {
	for i := 0; i < 9; i++ {
		*(*uint64)(unsafe.Pointer(&b[0])) = *(*uint64)(unsafe.Pointer(&b[0])) ^ *(*uint64)(unsafe.Pointer(&k.keys[i][0]))
		*(*uint64)(unsafe.Pointer(&b[8])) = *(*uint64)(unsafe.Pointer(&b[8])) ^ *(*uint64)(unsafe.Pointer(&k.keys[i][8]))
		a1 := b
		*b = ls_array[0][a1[0]]
		for o := 1; o < 16; o++ {
			*(*uint64)(unsafe.Pointer(&b[0])) = *(*uint64)(unsafe.Pointer(&b[0])) ^ *(*uint64)(unsafe.Pointer(&ls_array[o][a1[o]][0]))
			*(*uint64)(unsafe.Pointer(&b[8])) = *(*uint64)(unsafe.Pointer(&b[8])) ^ *(*uint64)(unsafe.Pointer(&ls_array[o][a1[o]][8]))
		}
	}
	*(*uint64)(unsafe.Pointer(&b[0])) = *(*uint64)(unsafe.Pointer(&b[0])) ^ *(*uint64)(unsafe.Pointer(&k.keys[9][0]))
	*(*uint64)(unsafe.Pointer(&b[8])) = *(*uint64)(unsafe.Pointer(&b[8])) ^ *(*uint64)(unsafe.Pointer(&k.keys[9][8]))
}

type randomReader struct {
}

func (_ randomReader) Read(p []byte) (n int, err error) {
	return rand.Read(p)
}
