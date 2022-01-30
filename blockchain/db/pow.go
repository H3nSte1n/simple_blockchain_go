package db

import (
	"blockchain/data/pow"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var POW_MODEL = "pows"

func InsertPoW(pow *pow.PoW) (primitive.ObjectID, error) {
	return Create[*pow.PoW](pow, POW_MODEL)
}

func GetPoWList() ([]pow.PoW, error) {
	return GetList[pow.PoW](POW_MODEL)
}