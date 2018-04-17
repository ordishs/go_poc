package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func main() {
	// j, _ := readJob("j1")
	// s, _ := readSubmit("s1")
	// // fmt.Println(j.toString(), s.toString())

	// cb := buildCoinbase(j.Coinbase1, j.Coinbase2, s.ExtraNonce1, s.ExtraNonce2)
	// merkleRoot := buildMerkleRootFromCoinbase(cb, j.MerkleBranches)
	// header := buildBlockHeader(j.Version, j.PreviousBlockHash, merkleRoot, s.NTime, j.Bits, s.Nonce)

	cb := []byte("b82c849c6abcd1ad17f4457333afc45723557348d2dda6974363253223b0f378")
	merkleBranches := []string{
		"4ea2296eff2cab120ecaa8ea268f663de7b0129f9132ba02de3bde7d3341431f",
		"39864fb58564307f69945596a9a8e188c256f09a3d67e17bba888e0ae2a99cba",
		"df46f060d8e9b986eda616f02a22a8cae6f8ae3f19cf2b1ef2c78d03f7bdbf61",
		"eb1a02c53bc472a480ef12dc609ca71a0137e5d4e80c1e87b138d8d96d4a37f6",
	}
	merkleRoot := buildMerkleRootFromCoinbase(cb, merkleBranches)
	fmt.Println(hex.EncodeToString(merkleRoot))

	time := make([]byte, 4)
	binary.LittleEndian.PutUint32(time, 1482001679)

	bits := make([]byte, 4)
	binary.LittleEndian.PutUint32(bits, 402885509)

	nonce := make([]byte, 4)
	binary.LittleEndian.PutUint32(nonce, 3814348197)

	header := buildBlockHeader(536870912, "000000000000000002424db0163641940c9fd999ec897b412ce64e36d6ab7650", merkleRoot, string(time), string(bits), string(nonce))

	s := hex.EncodeToString(header)
	fmt.Println(s)

	hash := sha256.Sum256(header)
	a := hex.EncodeToString(hash[:])
	fmt.Println(a)
	// h := DoubleHashB(cb)
	// fmt.Printf("%x", h)
}

// HashB calculates hash(b) and returns the resulting bytes.
func HashB(b []byte) []byte {
	hash := sha256.Sum256(b)
	return hash[:]
}

// DoubleHashB calculates hash(hash(b)) and returns the resulting bytes.
func DoubleHashB(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}

// 00000000000004fb02a2e369c302023101a36fa1c0c5a798b8b968276fac93c9-9a46fde8a01fee54-8eb10100-075140a9-5ad4f304.json

func buildCoinbase(coinbase1 string, coinbase2 string, extraNonce1 string, extraNonce2 string) []byte {
	c1, _ := hex.DecodeString(coinbase1)
	c2, _ := hex.DecodeString(coinbase2)
	e1, _ := hex.DecodeString(extraNonce1)
	e2, _ := hex.DecodeString(extraNonce2)

	a := []byte{}
	a = append(a, c1...)
	a = append(a, e1...)
	a = append(a, e2...)
	a = append(a, c2...)
	return a
}

func reverseBytes(a []byte) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func buildMerkleRootFromCoinbase(coinbaseHash []byte, merkleBranches []string) []byte {
	acc := coinbaseHash
	reverseBytes(acc)
	for i := 0; i < len(merkleBranches); i++ {
		branch, _ := hex.DecodeString(merkleBranches[i])
		concat := append(acc, branch...)
		hash := DoubleHashB(concat)
		acc = hash[:]
	}
	return acc
}

func buildBlockHeader(version uint32, previousBlockHash string, merkleRoot []byte, nTime string, nBits string, nonce string) []byte {
	v := make([]byte, 4)
	binary.LittleEndian.PutUint32(v, version)
	p, _ := hex.DecodeString(previousBlockHash)
	t, _ := hex.DecodeString(nTime)
	b, _ := hex.DecodeString(nBits)
	n, _ := hex.DecodeString(nonce)

	reverseBytes(p)
	reverseBytes(merkleRoot)
	reverseBytes(t)
	reverseBytes(b)
	reverseBytes(n)

	a := []byte{}
	a = append(a, v...)
	a = append(a, p...)
	a = append(a, merkleRoot...)
	a = append(a, t...)
	a = append(a, b...)
	a = append(a, n...)
	return a
}

/*
const blockHeader = new BlockHeader({
    version: Hex.fromInt(version),
    previousBlockHash: Hex.fromRPCByteOrder(previousBlockHash),
    merkleRoot: Hex.fromInternalByteOrder(merkleRoot),
    time: new Hex(time),
    bits: bitsHex,
    nonce: Hex.fromRPCByteOrder(nonce)
  })
*/
