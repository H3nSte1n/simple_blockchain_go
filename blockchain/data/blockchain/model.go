package chain

import (
	"blockchain/data/block"
	"blockchain/data/pow"
)

type Blockchain struct {
	Blocks  []block.Block
	Genesis block.Block
	PoWs    []pow.PoW
}
