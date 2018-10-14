package main

import (
	"crypto/sha256"
	"fmt"
)

//定义结构
type Block struct {
	PreHash []byte
	Hash    []byte
	Data    []byte
}

//创建区块
func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		PreHash: preBlockHash,
		Hash:    []byte{},
		Data:    []byte(data),
	}
	return &block
}

//生成哈希
func (block *Block) SetBlock() {
	blockInfo := append(block.PreHash, block.Data...)//block.Data...把切片拆分成单个元素，分别添加
	hash := sha256.Sum256(blockInfo)
	block.Hash = hash[:]
}


func main() {
	block := NewBlock("老师转一枚比特币给班长", []byte{})
	fmt.Printf("前区块哈希值 %x\n", block.PreHash)
	fmt.Printf("当前区块哈希值 %x\n", block.Hash)
	fmt.Printf("前区块哈希值 %s\n", block.Data)
}
