package core

import (
	"crypto/sha256"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Data         []byte
	Hash         string
	PreviousHash string
	Timestamp    time.Time
	Nonce        int
}

func (block *Block) GetBlockData() []byte {
	return block.Data
}

func (block *Block) CalculateHash() string {
	data, err := jsoniter.Marshal(block.Data)
	if err != nil {
		return ""
	}
	blockData := block.PreviousHash + string(data) + block.Timestamp.String() + strconv.Itoa(block.Nonce)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (block *Block) mine(difficulty int) {
	for !strings.HasPrefix(block.Hash, strings.Repeat("0", difficulty)) {
		block.Nonce++
		block.Hash = block.CalculateHash()
	}
}
