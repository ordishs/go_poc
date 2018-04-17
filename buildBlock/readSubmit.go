package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// MaestroSubmit : used to map Maestro submit to JSON
type MaestroSubmit struct {
	ExtraNonce1 string `json:"extraNonce1"`
	ExtraNonce2 string `json:"extraNonce2"`
	NTime       string `json:"nTime"`
	Nonce       string `json:"nonce"`
	SubmitTS    uint64 `json:"submitTS"`
	TS          uint64 `json:"ts"`
	JobID       string `json:"jobId"`
	Coin        string `json:"coin"`
	Height      uint32 `json:"height"`
	Hash        string `json:"hash"`
}

func (s MaestroSubmit) toString() string {
	bytes, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func readSubmit(name string) (MaestroSubmit, error) {
	raw, err := ioutil.ReadFile(name + ".json")
	if err != nil {
		return MaestroSubmit{}, err
	}

	var s MaestroSubmit
	json.Unmarshal(raw, &s)

	return s, nil
}
