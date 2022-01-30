package db

import (
	"blockchain/data/block"
)

var BLOCK_MODEL = "blocks"

func GetBlockList() ([]block.Block, error) {
	return GetList[block.Block](BLOCK_MODEL)
}

func InsertBlock(b *block.Block) (*int32, error) {
	return Create(b, BLOCK_MODEL)
}

func ResetBlockCollection() error {
	return ResetCollection(BLOCK_MODEL)
}
