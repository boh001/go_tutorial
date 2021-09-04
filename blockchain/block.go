package blockchain

import (
	"coin_tutorial/db"
	"coin_tutorial/utils"
	"crypto/sha256"
	"errors"
	"fmt"
)

type Block struct {
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prev_hash,omitempty"`
	Height   int    `json:"height"`
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data: data,
		Hash: "",
		PrevHash: prevHash,
		Height: height,
	}

	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("block not found")
func Findblock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}