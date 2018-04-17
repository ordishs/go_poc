package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func main() {
	j, _ := readJob("j2")
	s, _ := readSubmit("s2")
	// fmt.Println(j.toString(), s.toString())

	cb := buildCoinbase(j.Coinbase1, j.Coinbase2, s.ExtraNonce1, s.ExtraNonce2)
	fmt.Printf("cb: %x\n\n", cb)

	cbHash := sha256d(cb)
	fmt.Printf("ch: %x\n\n", cbHash)

	merkleRoot := buildMerkleRootFromCoinbase(cbHash, j.MerkleBranches)
	fmt.Printf("mr: %x\n\n", merkleRoot)

	time, _ := hex.DecodeString(s.NTime)
	reverseBytes(time)
	fmt.Printf("ti: %x\n\n", time)

	bits, _ := hex.DecodeString(j.Bits)
	reverseBytes(bits)
	fmt.Printf("bi: %x\n\n", bits)

	nonce, _ := hex.DecodeString(s.Nonce)
	reverseBytes(nonce)
	fmt.Printf("no: %x\n\n", nonce)

	header := buildBlockHeader(j.Version, j.PreviousBlockHash, merkleRoot, time, bits, nonce)
	fmt.Printf("he: %x\nlen=%v\n\n", header, len(header))

	hash := sha256d(header)
	reverseBytes(hash)
	fmt.Printf("ha: %x\n\n", hash)
	// a := hex.EncodeToString(hash[:])
	// fmt.Println(a)
}

// sha256d calculates hash(hash(b)) and returns the resulting bytes.
func sha256d(b []byte) []byte {
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
		hash := sha256d(concat)
		acc = hash[:]
	}
	reverseBytes(acc)
	return acc
}

func buildBlockHeader(version uint32, previousBlockHash string, merkleRoot []byte, time []byte, bits []byte, nonce []byte) []byte {
	v := make([]byte, 4)
	binary.LittleEndian.PutUint32(v, version)
	p, _ := hex.DecodeString(previousBlockHash)

	reverseBytes(p)

	a := []byte{}
	a = append(a, v...)
	a = append(a, p...)
	a = append(a, merkleRoot...)
	a = append(a, time...)
	a = append(a, bits...)
	a = append(a, nonce...)
	return a
}
