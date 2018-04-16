package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// MaestroJob : used to map Maestro jobs to JSON
type MaestroJob struct {
	JobID             string   `json:"jobId"`
	Height            uint32   `json:"height"`
	Coin              string   `json:"coin"`
	WalletAddress     string   `json:"walletAddress"`
	TS                uint64   `json:"ts"`
	PreviousBlockHash string   `json:"previousBlockHash"`
	Coinbase1         string   `json:"coinbase1"`
	Coinbase2         string   `json:"coinbase2"`
	MerkleBranches    []string `json:"merkleBranches"`
	Version           uint32   `json:"version"`
	Bits              string   `json:"bits"`
	Time              int64    `json:"time"`
	CleanJob          bool     `json:"cleanJob"`
	Target            string   `json:"target"`
	Difficulty        float64  `json:"difficulty"`
	BlockValue        float64  `json:"blockValue"`
}

func (j MaestroJob) toString() string {
	bytes, err := json.Marshal(j)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

func main() {

	raw, err := ioutil.ReadFile("./gbt.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var j MaestroJob
	json.Unmarshal(raw, &j)

	fmt.Printf("%+v\n\n\n", j)

	fmt.Println(j.toString())
}
