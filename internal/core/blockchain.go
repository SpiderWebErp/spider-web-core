package core

import "time"

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func NewBlockchain() *Blockchain {
	genesisBlock := Block{
		Hash:      "0",
		Timestamp: time.Now(),
	}
	return &Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		4,
	}
}

func (b *Blockchain) AddBlock(data []byte) {
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		Data:         data,
		PreviousHash: lastBlock.Hash,
		Timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b *Blockchain) IsValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.Hash != currentBlock.CalculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func (b *Blockchain) GetBlocks() []Block {
	return b.chain
}
