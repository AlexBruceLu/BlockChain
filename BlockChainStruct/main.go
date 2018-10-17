package main

func main() {
	bc := NewBlockChain()
	bc.AddBlocks("小明向小红转了5枚比特币")
	bc.AddBlocks("小明向小红转了5枚比特币")
	cli:=CLI{bc}
	cli.Run()
	//for i, v := range bc{
	//	fmt.Printf("----------当前区块高度 %d----------\n", i)
	//	fmt.Printf("当前区块前区块哈希值%x\n", v.PrevHash)
	//	fmt.Printf("当前区块所传递信息%s\n", v.Data)
	//	fmt.Printf("当前区块哈希值%x\n", v.Hash)
	//}

	//创建迭代器
	//it := bc.NewIterator()
	//for {
	//	block := it.Next()
	//	fmt.Printf("------------------------------------------\n\n")
	//	fmt.Printf("当前区块前区块哈希值: %x\n", block.PrevHash)
	//	fmt.Printf("当前区块所传递信息: %s\n", block.Data)
	//	fmt.Printf("当前区块哈希值: %x\n", block.Hash)
	//	if len(block.PrevHash) == 0 {
	//		fmt.Println("区块遍历完毕，已退出")
	//		break
	//	}
	//}


}
