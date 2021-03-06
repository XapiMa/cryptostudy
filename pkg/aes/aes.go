package aes

import (
	"fmt"
	"log"
)

// Cipher encrypts plain text
func Cipher(in []byte, key []byte, mode int, iv []byte) []byte {
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

	var out []byte
	switch mode {
	case ModeECB:
		out = ECBCipher(in, expandedKey, numOfBlocks)
	case ModeCBC:
		out = CBCCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCFB:
		out = CFBCipher(in, expandedKey, iv, numOfBlocks)
	case ModeOFB:
		out = OFBCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCTR:
		out = CTRCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCBCCTS:
		out = CBCCTSCipher(in, expandedKey, iv, numOfBlocks)
	default:
		log.Fatalln("Invalid encryption mode")
	}
	return out
}

// InvCipher decrypt given cipher text
func InvCipher(in, key []byte, mode int, iv []byte) []byte {
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

	var out []byte
	switch mode {
	case ModeECB:
		out = ECBInvCipher(in, expandedKey, numOfBlocks)
	case ModeCBC:
		out = CBCInvCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCFB:
		out = CFBInvCipher(in, expandedKey, iv, numOfBlocks)
	case ModeOFB:
		out = OFBInvCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCTR:
		out = CTRInvCipher(in, expandedKey, iv, numOfBlocks)
	case ModeCBCCTS:
		out = CBCCTSInvCipher(in, expandedKey, iv, numOfBlocks)
	default:
		log.Fatalln("Invalid encryption mode")
	}
	return out
}

func blockCipher(state, key []byte) {
	round := 0
	if round == PrintNRound {
		fmt.Printf("[Round %d]\n", round)
	}
	AddRoundKey(state, key[:Nb*BytesOfWords])
	printRoundBytes(state, round, "AddRoundKey")

	for round = 1; round <= Nr; round++ {
		if round == PrintNRound {
			fmt.Printf("[Round %d]\n", round)
		}
		SubBytes(state)
		printRoundBytes(state, round, "SubBytes")

		ShiftRows(state)
		printRoundBytes(state, round, "ShiftRows")

		if round < Nr {
			MixColumns(state)
			printRoundBytes(state, round, "MixColumns")

		}
		AddRoundKey(state, key[round*Nb*BytesOfWords:(round+1)*Nb*BytesOfWords])
		printRoundBytes(state, round, "AddRoundKey")

	}
}

func invBlockCipher(state, key []byte) {
	round := Nr
	if round == PrintNRound {
		fmt.Printf("[Round %d]\n", round)
	}
	AddRoundKey(state, key[round*Nb*BytesOfWords:(round+1)*Nb*BytesOfWords])
	printRoundBytes(state, round, "AddRoundKey")

	for round = Nr - 1; round >= 0; round-- {
		if round == PrintNRound {
			fmt.Printf("[Round %d]\n", round)
		}
		InvShiftRows(state)
		printRoundBytes(state, round, "InvShiftRows")

		InvSubBytes(state)
		printRoundBytes(state, round, "InvSubBytes")

		AddRoundKey(state, key[round*Nb*BytesOfWords:(round+1)*Nb*BytesOfWords])
		printRoundBytes(state, round, "AddRoundKey")

		if round > 0 {
			InvMixColumns(state)
			printRoundBytes(state, round, "InvMixColumns")

		}
	}
}
