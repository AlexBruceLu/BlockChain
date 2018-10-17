package main

import (
	"github.com/bolt"
	"log"
)

//实现迭代器功能
type BlockChainIterator struct {
	Db *bolt.DB
	//游标,不断索引哈希值
	CurrentHashPoint []byte
}

//创建迭代器
func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{
		bc.Db,
		bc.Tail,
	}
}

//实现迭代器Next方法
//1. 返回当前区块
//2. 指针前移
func (it *BlockChainIterator) Next() *Block {
	var block Block
	it.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BlockBucket))
		if bucket == nil {
			log.Panic("Next bucket不能为空请检查")
		}
		tmpBlock := bucket.Get(it.CurrentHashPoint)
		//fmt.Printf("tmpblock: %s\n", tmpBlock)
		block = Deserialize(tmpBlock)
		it.CurrentHashPoint = block.PrevHash
		return nil
	})
	return &block
}
