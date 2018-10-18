package main

import (
	"fmt"
	"os"
)

type CLI struct {
	BC *BlockChain
}

const Usage = `
	addBlock --data DATA	"添加区块"
	printChain				"正向打印区块"
	printChainR				"反向打印区块"
`

//接受参数的函数
func (c *CLI) Run() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Usage)
		return
	}
	cmd := args[1]
	switch cmd {
	case "addBlock":
		fmt.Println("添加新区块")
		if len(args)==4 && args[2]=="--data" {
			data := args[3]
			c.AddBlock(data)
		}else {
			fmt.Println("添加区块使用参数不当")
			fmt.Println(Usage)
		}

	case "printChain":
		fmt.Println("正向打印区块")
		c.PrintChain()
	case "printChainR":
		fmt.Println("反向打印区块")
		c.PrintChainR()
	default:
		fmt.Println("无效的输出，请检查")
		fmt.Println(Usage)
	}
}
