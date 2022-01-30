package chain

import (
	"blockchain/data/pow"
	"blockchain/data/block"
)

type Blockchain struct {
	Blocks []block.Block
	Genesis block.Block
	PoWs 		[]pow.PoW
}