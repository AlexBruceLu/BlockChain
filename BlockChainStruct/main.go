package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.AddBlocks("小明向小红转了5枚比特币")
	bc.AddBlocks("小明向小红转了5枚比特币")
	for i, v := range bc.Blocks {
		fmt.Printf("----------当前区块高度%d----------\n", i)
		fmt.Printf("当前区块前区块哈希值%x\n", v.PrevHash)
		fmt.Printf("当前区块所传递信息%s\n", v.Data)
		fmt.Printf("当前区块哈希值%x\n", v.Hash)
	}

}
