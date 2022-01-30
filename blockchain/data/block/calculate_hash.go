package block

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

func (b Block) CalculateHash() big.Int {
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
