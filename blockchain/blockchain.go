package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"` // block의 특징
	PrevHash string `json:"prevHash,omitempty"`
	Height   int    `json:"height"`
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
	newBlock := Block{Data, "", getLastBlockHash(), len(GetBlockChain().blocks) + 1}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(Data string) {
	b.blocks = append(b.blocks, createBlock(Data))
}

func (b *blockchain) AllBlocks() []*Block {
	return b.blocks
}

var NotFoundError = errors.New("block not found")

func (b *blockchain) GetBlock(height int) (*Block, error) {
	if height > len(b.blocks) {
		return nil, NotFoundError
	}
	return b.blocks[height-1], nil
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
