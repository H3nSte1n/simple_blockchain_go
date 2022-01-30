package chain

import (
	"time"
	"fmt"
	"blockchain/data/block"
	"blockchain/db"
)

func (chain *Blockchain) AddBlock(data map[string]interface{}) {
	currentPoW := chain.PoWs[len(chain.PoWs)-1]
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := block.Block{
		Nr: prevBlock.Nr + 1,
		PrevHash: prevBlock.Hash,
		Timestamp: time.Now().Unix(),
		Data: data,
		Nonce: 0,
	}

	chain.mineBlock(&newBlock)
	if !newBlock.Validate(currentPoW) {
		return
	}
	chain.Blocks = append(chain.Blocks, newBlock)

	if len(chain.Blocks) % 2016 == 0 {
		difficulty := *currentPoW.Difficulty
		chain.AddPoW(difficulty)
	}


	db.InsertBlock(&newBlock)
	fmt.Println("chain:", chain)
}