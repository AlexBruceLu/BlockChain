package main

import "fmt"

func main() {
	prevHash:=[]byte("")
	data:=[]byte("小明向小红转了5枚比特币")
	bc:=NewBlock(prevHash,data)
	fmt.Printf("当前区块前区块哈希值%v\n",bc.PrevHash)
	fmt.Printf("当前区块所传递信息%s\n",bc.Data)
	fmt.Printf("当前区块哈希值%v\n",bc.Hash)
}
