package main

import "fmt"

func (c *CLI) AddBlock(data string) {
	c.BC.AddBlocks(data)
	fmt.Println("区块添加成功！")
}

func (c *CLI) PrintChain() {
	//正向打印区块链
	c.BC.PrintBlockChain()
	fmt.Println("区块正向打印完毕！")
}

func (c *CLI) PrintChainR() {
	//反向打印区块
	it := c.BC.NewIterator()
	for {
		block := it.Next()
		fmt.Printf("==========================\n\n")
		fmt.Printf("版本号: %d\n", block.Version)
		fmt.Printf("前区块哈希值: %x\n", block.PrevHash)
		fmt.Printf("梅克尔根: %x\n", block.MerKelRoot)
		fmt.Printf("时间戳: %d\n", block.TimeStamp)
		fmt.Printf("难度值(随便写的）: %d\n", block.Difficulty)
		fmt.Printf("随机数 : %d\n", block.Nonce)
		fmt.Printf("当前区块哈希值: %x\n", block.Hash)
		fmt.Printf("区块数据 :%s\n", block.Data)
		if len(block.PrevHash) == 0 {
			fmt.Println("区块反向打印完毕！")
			break
		}
	}
}
