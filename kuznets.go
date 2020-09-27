package main

var sbox = [0x100]byte{
	252, 238, 221, 17, 207, 110, 49, 22, 251, 196, 250, 218, 35, 197, 4, 77, 233, 119, 240, 219, 147, 46, 153, 186, 23, 54, 241, 187, 20, 205, 95, 193, 249, 24, 101, 90, 226, 92, 239, 33, 129, 28, 60, 66, 139, 1, 142, 79, 5, 132, 2, 174, 227, 106, 143, 160, 6, 11, 237, 152, 127, 212, 211, 31, 235, 52, 44, 81, 234, 200, 72, 171, 242, 42, 104, 162, 253, 58, 206, 204, 181, 112, 14, 86, 8, 12, 118, 18, 191, 114, 19, 71, 156, 183, 93, 135, 21, 161, 150, 41, 16, 123, 154, 199, 243, 145, 120, 111, 157, 158, 178, 177, 50, 117, 25, 61, 255, 53, 138, 126, 109, 84, 198, 128, 195, 189, 13, 87, 223, 245, 36, 169, 62, 168, 67, 201, 215, 121, 214, 246, 124, 34, 185, 3, 224, 15, 236, 222, 122, 148, 176, 188, 220, 232, 40, 80, 78, 51, 10, 74, 167, 151, 96, 115, 30, 0, 98, 68, 26, 184, 56, 130, 100, 159, 38, 65, 173, 69, 70, 146, 39, 94, 85, 47, 140, 163, 165, 125, 105, 213, 149, 59, 7, 88, 179, 64, 134, 172, 29, 247, 48, 55, 107, 228, 136, 217, 231, 137, 225, 27, 131, 73, 76, 63, 248, 254, 141, 83, 170, 144, 202, 216, 133, 97, 32, 113, 103, 164, 45, 43, 9, 91, 203, 155, 37, 208, 190, 229, 108, 82, 89, 166, 116, 210, 230, 244, 180, 192, 209, 102, 175, 194, 57, 75, 99, 182}

func x(a [16]byte, b [16]byte) [16]byte {
	var result [16]byte
	result[0] = a[0] ^ b[0]
	result[1] = a[1] ^ b[1]
	result[2] = a[2] ^ b[2]
	result[3] = a[3] ^ b[3]
	result[4] = a[4] ^ b[4]
	result[5] = a[5] ^ b[5]
	result[6] = a[6] ^ b[6]
	result[7] = a[7] ^ b[7]
	result[8] = a[8] ^ b[8]
	result[9] = a[9] ^ b[9]
	result[10] = a[10] ^ b[10]
	result[11] = a[11] ^ b[11]
	result[12] = a[12] ^ b[12]
	result[13] = a[13] ^ b[13]
	result[14] = a[14] ^ b[14]
	result[15] = a[15] ^ b[15]
	return result
}

func s(a [16]byte) [16]byte {
	var bytes [16]byte
	bytes[0] = sbox[a[0]]
	bytes[1] = sbox[a[1]]
	bytes[2] = sbox[a[2]]
	bytes[3] = sbox[a[3]]
	bytes[4] = sbox[a[4]]
	bytes[5] = sbox[a[5]]
	bytes[6] = sbox[a[6]]
	bytes[7] = sbox[a[7]]
	bytes[8] = sbox[a[8]]
	bytes[9] = sbox[a[9]]
	bytes[10] = sbox[a[10]]
	bytes[11] = sbox[a[11]]
	bytes[12] = sbox[a[12]]
	bytes[13] = sbox[a[13]]
	bytes[14] = sbox[a[14]]
	bytes[15] = sbox[a[15]]
	return bytes
}
func gfm(x byte, y byte) byte {
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

func lv128v8(a [16]byte) byte {
	//var b byte = 0
	//for i := 0; i < 16; i++ {
	//	b ^= gfm(a[i], gfc[i])
	//}
	//return b
	return gfm(148, a[0]) ^
		gfm(32, a[1]) ^
		gfm(133, a[2]) ^
		gfm(16, a[3]) ^
		gfm(194, a[4]) ^
		gfm(192, a[5]) ^
		gfm(1, a[6]) ^
		gfm(251, a[7]) ^
		gfm(1, a[8]) ^
		gfm(192, a[9]) ^
		gfm(194, a[10]) ^
		gfm(16, a[11]) ^
		gfm(133, a[12]) ^
		gfm(32, a[13]) ^
		gfm(148, a[14]) ^
		gfm(1, a[15])
}

func r(a [16]byte) [16]byte {
	var bytes [16]byte
	var a15 = lv128v8(a)
	bytes[0]=a15
	bytes[1]=a[0]
	bytes[2]=a[1]
	bytes[3]=a[2]
	bytes[4]=a[3]
	bytes[5]=a[4]
	bytes[6]=a[5]
	bytes[7]=a[6]
	bytes[8]=a[7]
	bytes[9]=a[8]
	bytes[10]=a[9]
	bytes[11]=a[10]
	bytes[12]=a[11]
	bytes[13]=a[12]
	bytes[14]=a[13]
	bytes[15]=a[14]

	return bytes// append([]byte{a15}, a[:15]...)
}

func L(a [16]byte) [16]byte {
	// LISPers gonna LISP
	return r(r(r(r(r(r(r(r(r(r(r(r(r(r(r(r(a))))))))))))))))
}

func unwrap_keys(master []byte) [][]byte {
	//keys := make([][]byte,)
	return nil
}
