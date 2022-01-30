package data

import (
	"crypto/sha256"
	"encoding/json"
	"strconv"
	"math/big"
	"fmt"
)

type Block struct {
	Nr 				int
	Hash			string
	PrevHash  string
	Timestamp int64
	Data      map[string]interface{}
	Nonce     int
}

func (b Block) calculateHash() big.Int {
	var bigInt big.Int
	var blockHash [32]byte

	data, _ := json.Marshal(b.Data)
	blockData := b.PrevHash +
		string(data) +
		fmt.Sprint(b.Timestamp) +
		strconv.Itoa(b.Nonce) +
		strconv.Itoa(b.Nr)

	blockHash = sha256.Sum256([]byte(blockData))

	bigInt.SetBytes(blockHash[:])
	return bigInt
}

func (b Block) validate(pow PoW) bool {
	return pow.proof(b.calculateHash())
}