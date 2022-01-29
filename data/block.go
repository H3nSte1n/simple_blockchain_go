package data

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Hash			string
	PrevHash  string
	Timestamp time.Time
	Data      map[string]interface{}
	PoW     int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.Data)
	blockData := b.PrevHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.PoW) // b.Timestamp.String() == nounce
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}