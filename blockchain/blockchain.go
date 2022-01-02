package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string // block의 특징
	PrevHash string
}

type blockchain struct {
	blocks []*Block
}

var b *blockchain
var once sync.Once

func (block *Block) calculateHash() {
	hash := sha256.Sum256([]byte(block.Data + block.PrevHash))
	block.Hash = fmt.Sprintf("%x", hash)
}

func getLastBlockHash() string {
	totalBlock := len(GetBlockChain().blocks)
	if totalBlock == 0 {
		return ""
	}
	return GetBlockChain().blocks[totalBlock-1].Hash
}

func createBlock(Data string) *Block {
	newBlock := Block{Data, "", getLastBlockHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(Data string) {
	b.blocks = append(b.blocks, createBlock(Data))
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

func GetBlockChain() *blockchain {
	if b == nil {
		// once.Do(...) : 프로그램을 병렬적으로 실행했을때 "오직 한 번만 실행하도록 한다."
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}
