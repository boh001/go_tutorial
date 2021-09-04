package blockchain

import (
	"coin_tutorial/db"
	"coin_tutorial/utils"
	"fmt"
	"sync"
)

type blockchain struct {
	NewHash string `json:"new_hash"`
	Height int `json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewHash, b.Height + 1)
	b.NewHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				b.restore(checkpoint)
			}
		})
	}
	fmt.Printf(b.NewHash)
	return b
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewHash
	for {
		block, _ := Findblock(hashCursor)
		blocks = append(blocks, block)
		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			break
		}
	}
	return blocks
}