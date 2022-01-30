package db

import (
	"blockchain/data/pow"
)

var POW_MODEL = "pows"

func InsertPoW(p *pow.PoW) (*int32, error) {
	return Create(p, POW_MODEL)
}

func GetPoWList() ([]pow.PoW, error) {
	return GetList[pow.PoW](POW_MODEL)
}

func ResetPoWCollection() error {
	return ResetCollection(POW_MODEL)
}
