package main

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestBuildCoinbase(t *testing.T) {
	coinbase1 := "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff1f034707080e2f636f696e6765656b2e636f6d2f"
	coinbase2 := "ffffffff0146e0824a000000001976a914513c8cb8f8f031c88b63a58aabce2f7560d2658188ac00000000"
	extraNonce1 := "9a46fde8a01fee54"
	extraNonce2 := "8eb10100"
	expectedCoinbase := []byte{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 31, 3, 71, 7, 8, 14, 47, 99, 111, 105, 110, 103, 101, 101, 107, 46, 99, 111, 109, 47, 154, 70, 253, 232, 160, 31, 238, 84, 142, 177, 1, 0, 255, 255, 255, 255, 1, 70, 224, 130, 74, 0, 0, 0, 0, 25, 118, 169, 20, 81, 60, 140, 184, 248, 240, 49, 200, 139, 99, 165, 138, 171, 206, 47, 117, 96, 210, 101, 129, 136, 172, 0, 0, 0, 0}
	coinbase := buildCoinbase(coinbase1, coinbase2, extraNonce1, extraNonce2)

	if len(coinbase) != 116 {
		t.Errorf("Expected coinbase to be 116 long, got %v", len(coinbase))
	}
	if !reflect.DeepEqual(coinbase, expectedCoinbase) {
		t.Errorf("Expected coinbase to be %v, got %v", hex.EncodeToString(expectedCoinbase), hex.EncodeToString(coinbase))
	}
}

func TestReverseBytes(t *testing.T) {
	initialBytes := []byte{1, 0, 0, 0, 1, 0, 0}
	expectedBytes := []byte{0, 0, 1, 0, 0, 0, 1}

	reverseBytes(initialBytes)
	if !reflect.DeepEqual(initialBytes, expectedBytes) {
		t.Errorf("Expected reversed bytes to be %v, got %v", expectedBytes, initialBytes)
	}
}

func TestBuildMerkleRootFromCoinbase(t *testing.T) {
	coinbaseHash := []byte{1, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 255, 255, 31, 3, 71, 7, 8, 14, 47, 99, 111, 105, 110, 103, 101, 101, 107, 46, 99, 111, 109, 47, 154, 70, 253, 232, 160, 31, 238, 84, 142, 177, 1, 0, 255, 255, 255, 255, 1, 70, 224, 130, 74, 0, 0, 0, 0, 25, 118, 169, 20, 81, 60, 140, 184, 248, 240, 49, 200, 139, 99, 165, 138, 171, 206, 47, 117, 96, 210, 101, 129, 136, 172, 0, 0, 0, 0}
	merkleBranches := []string{
		"7dc75d361adc8013fd6064eef56116c000268fd4c0bad998f614405c3fe368cc",
		"d8558567ff06e44620a048d85281d99dab174b5efdce4134aa9d338a91aa86d7",
		"5bb5ffe301c895689636c0e167ed792cd8d83e9c62587ff76601ef5ecc8b8a82",
	}

	expectedMerkleRoot := []byte{37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244}
	merkleRoot := buildMerkleRootFromCoinbase(coinbaseHash, merkleBranches)

	if len(merkleRoot) != 32 {
		t.Errorf("Expected merkle root to be 32 long, got %v", len(merkleRoot))
	}
	if !reflect.DeepEqual(merkleRoot, expectedMerkleRoot) {
		t.Errorf("Expected merkle root to be %v, got %v", expectedMerkleRoot, merkleRoot)
	}
}

func TestBuildBlockHeader(t *testing.T) {
	version := uint32(536870912)
	previousBlockHash := "000000000000000000b429c70d0c2520216b2ef63b2e183fd92476dd0fe46183"
	merkleRoot := []byte{37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244}
	time := []byte{37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244}
	bits := []byte{37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244}
	nonce := []byte{37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244}

	expectedBlockHeader := []byte{0, 0, 0, 32, 131, 97, 228, 15, 221, 118, 36, 217, 63, 24, 46, 59, 246, 46, 107, 33, 32, 37, 12, 13, 199, 41, 180, 0, 0, 0, 0, 0, 0, 0, 0, 0, 37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244, 37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244, 37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244, 37, 196, 176, 1, 49, 66, 103, 0, 243, 75, 216, 164, 5, 16, 89, 136, 83, 53, 163, 224, 183, 177, 0, 48, 252, 243, 35, 184, 225, 124, 37, 244}

	blockHeader := buildBlockHeader(version, previousBlockHash, merkleRoot, time, bits, nonce)

	if len(blockHeader) != 164 {
		t.Errorf("Expected block header to be 32 long, got %v", len(blockHeader))
	}

	if !reflect.DeepEqual(blockHeader, expectedBlockHeader) {
		t.Errorf("Expected block header to be %v, got %v", expectedBlockHeader, blockHeader)
	}

}
