package db

import (
	"blockchain/data/block"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

var BLOCK_MODEL = "blocks"

func InsertBlock(block *block.Block) (primitive.ObjectID, error) {
	return Create[*block.Block](block, BLOCK_MODEL)
}

func GetBlockList() ([]block.Block, error) {
	return GetList[block.Block](BLOCK_MODEL)
}