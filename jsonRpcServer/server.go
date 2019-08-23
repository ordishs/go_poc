package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Handler comment
type Handler struct{}

// Request comment
type Request struct {
	// JobID             string   `json:"jobId"`
	// Height            uint32   `json:"height"`
	// Coin              string   `json:"coin"`
	// WalletAddress     string   `json:"walletAddress"`
	// TS                uint64   `json:"ts"`
	PreviousBlockHash string   `json:"previousBlockHash"`
	Coinbase1         string   `json:"coinbase1"`
	Coinbase2         string   `json:"coinbase2"`
	MerkleBranches    []string `json:"merkleBranches"`
	Version           uint32   `json:"version"`
	Bits              string   `json:"bits"`
	Time              string   `json:"time"`
	// CleanJob          bool     `json:"cleanJob"`
	// Target            string   `json:"target"`
	// Difficulty        float64  `json:"difficulty"`
	// BlockValue        float64  `json:"blockValue"`
	ExtraNonce1 string `json:"extraNonce1"`
	ExtraNonce2 string `json:"extraNonce2"`
	Nonce       string `json:"nonce"`
}

// Response comment
type Response struct {
	Accepted bool
	Hash     string
}

// Hello is a comment
func (h *Handler) Hello(req Request, res *Response) (err error) {
	// fmt.Printf("%+v\n", req)

	// req is a map with keys of string and values of interface{}
	// a := req
	// fmt.Printf("%+v", a)
	// b := a.(map[string]interface{}) // Cast to map[string]interface{}
	accepted, hash := handleBlock(req)
	// fmt.Printf("b is %v\n", b)

	res.Accepted = accepted
	res.Hash = hash
	return nil
}

func startServer() {
	handler := new(Handler)

	server := rpc.NewServer()
	server.Register(handler)

	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

func main() {
	startServer()
}
