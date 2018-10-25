package aes

import (
	"log"
)

const (
	// BytesOfWords represents the number of bytes for each word
	BytesOfWords = 4

	// KeyLength128 represents the key length (word) for AES-128
	KeyLength128 = 4
	// BlockSize128 represents the block size (word) for AES-128
	BlockSize128 = 4
	// NumOfRounds128 represents the number of round for AES-128
	NumOfRounds128 = 10

	// KeyLength192 represents the key length (word) for AES-192
	KeyLength192 = 6
	// BlockSize192 represents the block size (word) for AES-192
	BlockSize192 = 4
	// NumOfRounds192 represents the number of round for AES-192
	NumOfRounds192 = 12

	// KeyLength256 represents the key length (word) for AES-256
	KeyLength256 = 8
	// BlockSize256 represents the block size (word) for AES-256
	BlockSize256 = 4
	// NumOfRounds256 represents the number of round for AES-256
	NumOfRounds256 = 14
)

var (
	sbox = [][]byte{
		[]byte{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
		[]byte{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
		[]byte{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
		[]byte{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
		[]byte{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
		[]byte{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
		[]byte{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
		[]byte{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
		[]byte{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
		[]byte{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
		[]byte{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
		[]byte{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
		[]byte{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
		[]byte{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
		[]byte{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
		[]byte{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
	}

	invSbox = [][]byte{
		[]byte{0x52, 0x09, 0x6a, 0xd5, 0x30, 0x36, 0xa5, 0x38, 0xbf, 0x40, 0xa3, 0x9e, 0x81, 0xf3, 0xd7, 0xfb},
		[]byte{0x7c, 0xe3, 0x39, 0x82, 0x9b, 0x2f, 0xff, 0x87, 0x34, 0x8e, 0x43, 0x44, 0xc4, 0xde, 0xe9, 0xcb},
		[]byte{0x54, 0x7b, 0x94, 0x32, 0xa6, 0xc2, 0x23, 0x3d, 0xee, 0x4c, 0x95, 0x0b, 0x42, 0xfa, 0xc3, 0x4e},
		[]byte{0x08, 0x2e, 0xa1, 0x66, 0x28, 0xd9, 0x24, 0xb2, 0x76, 0x5b, 0xa2, 0x49, 0x6d, 0x8b, 0xd1, 0x25},
		[]byte{0x72, 0xf8, 0xf6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xd4, 0xa4, 0x5c, 0xcc, 0x5d, 0x65, 0xb6, 0x92},
		[]byte{0x6c, 0x70, 0x48, 0x50, 0xfd, 0xed, 0xb9, 0xda, 0x5e, 0x15, 0x46, 0x57, 0xa7, 0x8d, 0x9d, 0x84},
		[]byte{0x90, 0xd8, 0xab, 0x00, 0x8c, 0xbc, 0xd3, 0x0a, 0xf7, 0xe4, 0x58, 0x05, 0xb8, 0xb3, 0x45, 0x06},
		[]byte{0xd0, 0x2c, 0x1e, 0x8f, 0xca, 0x3f, 0x0f, 0x02, 0xc1, 0xaf, 0xbd, 0x03, 0x01, 0x13, 0x8a, 0x6b},
		[]byte{0x3a, 0x91, 0x11, 0x41, 0x4f, 0x67, 0xdc, 0xea, 0x97, 0xf2, 0xcf, 0xce, 0xf0, 0xb4, 0xe6, 0x73},
		[]byte{0x96, 0xac, 0x74, 0x22, 0xe7, 0xad, 0x35, 0x85, 0xe2, 0xf9, 0x37, 0xe8, 0x1c, 0x75, 0xdf, 0x6e},
		[]byte{0x47, 0xf1, 0x1a, 0x71, 0x1d, 0x29, 0xc5, 0x89, 0x6f, 0xb7, 0x62, 0x0e, 0xaa, 0x18, 0xbe, 0x1b},
		[]byte{0xfc, 0x56, 0x3e, 0x4b, 0xc6, 0xd2, 0x79, 0x20, 0x9a, 0xdb, 0xc0, 0xfe, 0x78, 0xcd, 0x5a, 0xf4},
		[]byte{0x1f, 0xdd, 0xa8, 0x33, 0x88, 0x07, 0xc7, 0x31, 0xb1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xec, 0x5f},
		[]byte{0x60, 0x51, 0x7f, 0xa9, 0x19, 0xb5, 0x4a, 0x0d, 0x2d, 0xe5, 0x7a, 0x9f, 0x93, 0xc9, 0x9c, 0xef},
		[]byte{0xa0, 0xe0, 0x3b, 0x4d, 0xae, 0x2a, 0xf5, 0xb0, 0xc8, 0xeb, 0xbb, 0x3c, 0x83, 0x53, 0x99, 0x61},
		[]byte{0x17, 0x2b, 0x04, 0x7e, 0xba, 0x77, 0xd6, 0x26, 0xe1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0c, 0x7d},
	}

	// Nk represents the length of key
	Nk int
	// Nb represents the block size
	Nb int
	// Nr represents the number of rounds
	Nr int
)

// Cipher encrypts plain text
func Cipher(in []byte, key []byte) []byte {
	switch len(key) {
	case 16:
		Nk = KeyLength128
		Nb = BlockSize128
		Nr = NumOfRounds128
	case 24:
		Nk = KeyLength192
		Nb = BlockSize192
		Nr = NumOfRounds192
	case 32:
		Nk = KeyLength256
		Nb = BlockSize256
		Nr = NumOfRounds256
	default:
		log.Fatalln("AES key length must be one of 128, 192, 256 bit")
	}

	expandedKey := make([]byte, BytesOfWords*Nb*(Nr+1))
	keyExpansion(key, expandedKey)

	numOfBlocks := len(in) / (Nb * BytesOfWords)
	if len(in)%(Nb*BytesOfWords) != 0 {
		numOfBlocks++
	}
	out := make([]byte, numOfBlocks*Nb*BytesOfWords)

	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		if stateLength < Nb*BytesOfWords {
			padding := Nb*BytesOfWords - stateLength
			for i := stateLength; i < Nb*BytesOfWords; i++ {
				state[i] = byte(padding)
			}
		}

		round := 0
		addRoundKey(state, key, round)

		for round = 1; round < Nr; round++ {
			subBytes(state)
			shiftRows(state)
			mixColumns(state)
			addRoundKey(state, expandedKey, round)
		}

		// FinalRound
		subBytes(state)
		shiftRows(state)
		addRoundKey(state, expandedKey, round)

		copy(out[from:from+Nb*BytesOfWords], state)
	}
	return out
}

// InvCipher decrypt given cipher text
func InvCipher(in, key []byte) []byte {
	switch len(key) {
	case 16:
		Nk = KeyLength128
		Nb = BlockSize128
		Nr = NumOfRounds128
	case 24:
		Nk = KeyLength192
		Nb = BlockSize192
		Nr = NumOfRounds192
	case 32:
		Nk = KeyLength256
		Nb = BlockSize256
		Nr = NumOfRounds256
	default:
		log.Fatalln("AES key length must be one of 128, 192, 256 bit")
	}

	expandedKey := make([]byte, BytesOfWords*Nb*(Nr+1))
	keyExpansion(key, expandedKey)

	numOfBlocks := len(in) / (Nb * BytesOfWords)
	if len(in)%(Nb*BytesOfWords) != 0 {
		numOfBlocks++
	}
	out := make([]byte, numOfBlocks*Nb*BytesOfWords)

	var last int
	for i := 0; i < numOfBlocks; i++ {
		state := make([]byte, Nb*BytesOfWords)
		from := i * Nb * BytesOfWords
		to := (i + 1) * Nb * BytesOfWords

		stateLength := to - from
		if i == numOfBlocks-1 && to > len(in) {
			stateLength = len(in) - from
		}
		copy(state, in[from:from+stateLength])

		round := Nr
		addRoundKey(state, expandedKey, round)

		for round = Nr - 1; round > 0; round-- {
			invShiftRows(state)
			invSubBytes(state)
			addRoundKey(state, expandedKey, round)
			invMixColumns(state)
		}

		invShiftRows(state)
		invSubBytes(state)
		addRoundKey(state, expandedKey, round)

		if i == numOfBlocks-1 && int(state[Nb*BytesOfWords-1]) < Nb*BytesOfWords {
			padding := int(state[Nb*BytesOfWords-1])
			last = from + Nb*BytesOfWords - padding
			copy(out[from:last], state[:Nb*BytesOfWords-padding])
		} else {
			last = from + Nb*BytesOfWords
			copy(out[from:last], state)
		}
	}
	return out[:last]
}

func keyExpansion(key []byte, expanded []byte) {
	copy(expanded, key)

	rc := byte(1) // round constant

	for i := Nk; i < Nb*(Nr+1); i++ {
		tmp := make([]byte, 4)
		copy(tmp, expanded[i*4-BytesOfWords:i*4]) // copy previous word from expanded key to tmp
		if i%Nk == 0 {
			rotWord(tmp)
			subWord(tmp)
			tmp[0] ^= rc
			rc = mul(rc, 2)
		} else if Nk > 6 && i%Nk == 4 {
			subWord(tmp)
		}

		for j := 0; j < BytesOfWords; j++ {
			expanded[i*4+j] = expanded[(i-Nk)*4+j] ^ tmp[j]
		}
	}
}

func mul(num1, num2 byte) byte {
	switch num2 {
	case 0:
		return 0
	case 1:
		return num1
	case 2:
		return mul2(num1)
	}

	remains := int(num2)
	result := num1
	i := 2
	for ; i <= remains; i *= 2 {
		result = mul2(result)
	}
	remains -= i / 2
	if remains > 0 {
		result ^= mul(num1, byte(remains))
	}
	return result
}

func mul2(num byte) byte {
	tmp := int(num) << 1 // multiplied by 2
	if tmp&0x100 != 0 {  //if 0x100 is set
		tmp ^= 0x11b // mod by 0x11b (x^8 + x^4 + x^3 + x + 1)
	}
	return byte(tmp)
}

func rotWord(word []byte) {
	tmp := word[0]
	for i := 0; i < BytesOfWords-1; i++ {
		word[i] = word[i+1]
	}
	word[BytesOfWords-1] = tmp
}

func subWord(word []byte) {
	for i := 0; i < BytesOfWords; i++ {
		x := word[i] >> 4
		y := word[i] & 0xf
		word[i] = sbox[x][y]
	}
}

func subBytes(state []byte) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		x := state[i] >> 4
		y := state[i] & 0xf
		state[i] = sbox[x][y]
	}
}

func shiftRows(state []byte) {
	tmp := make([]byte, Nb*BytesOfWords)
	copy(tmp, state)

	for i := 1; i < BytesOfWords; i++ { // i is a row index
		colOffset := i
		for j := 0; j < Nb; j++ { // j is a column index
			column := (j + colOffset) % Nb
			state[j*4+i] = tmp[column*4+i]
		}
	}
}

func mixColumns(state []byte) {
	bytes := make([]byte, Nb*BytesOfWords)
	copy(bytes, state)

	for i := 0; i < Nb; i++ {
		mulBy2 := make([]byte, BytesOfWords)
		for j := 0; j < BytesOfWords; j++ {
			mulBy2[j] = mul(bytes[i*4+j], 2)
		}

		state[i*4] = mulBy2[0] ^ (mulBy2[1] ^ bytes[i*4+1]) ^ bytes[i*4+2] ^ bytes[i*4+3]
		state[i*4+1] = mulBy2[1] ^ (mulBy2[2] ^ bytes[i*4+2]) ^ bytes[i*4+0] ^ bytes[i*4+3]
		state[i*4+2] = mulBy2[2] ^ (mulBy2[3] ^ bytes[i*4+3]) ^ bytes[i*4+0] ^ bytes[i*4+1]
		state[i*4+3] = mulBy2[3] ^ (mulBy2[0] ^ bytes[i*4+0]) ^ bytes[i*4+1] ^ bytes[i*4+2]
	}
}

func addRoundKey(state, key []byte, round int) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		state[i] ^= key[round*BlockSize128*BytesOfWords+i]
	}
}

func invShiftRows(state []byte) {
	tmp := make([]byte, Nb*BytesOfWords)
	copy(tmp, state)

	for i := 1; i < BytesOfWords; i++ { // i is a row index
		colOffset := i
		for j := 0; j < Nb; j++ { // j is a column index
			column := (j + colOffset) % Nb
			state[column*4+i] = tmp[j*4+i]
		}
	}
}

func invSubBytes(state []byte) {
	for i := 0; i < Nb*BytesOfWords; i++ {
		x := state[i] >> 4
		y := state[i] & 0xf
		state[i] = invSbox[x][y]
	}
}

func invMixColumns(state []byte) {
	bytes := make([]byte, Nb*BytesOfWords)
	copy(bytes, state)

	for i := 0; i < Nb; i++ {
		state[i*4] = mul(bytes[i*4], 0x0e) ^ mul(bytes[i*4+1], 0x0b) ^ mul(bytes[i*4+2], 0x0d) ^ mul(bytes[i*4+3], 0x09)
		state[i*4+1] = mul(bytes[i*4+1], 0x0e) ^ mul(bytes[i*4+2], 0x0b) ^ mul(bytes[i*4+3], 0x0d) ^ mul(bytes[i*4], 0x09)
		state[i*4+2] = mul(bytes[i*4+2], 0x0e) ^ mul(bytes[i*4+3], 0x0b) ^ mul(bytes[i*4], 0x0d) ^ mul(bytes[i*4+1], 0x09)
		state[i*4+3] = mul(bytes[i*4+3], 0x0e) ^ mul(bytes[i*4], 0x0b) ^ mul(bytes[i*4+1], 0x0d) ^ mul(bytes[i*4+2], 0x09)
	}
}
