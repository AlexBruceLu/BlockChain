package main

import (
	"fmt"
	"os"
	"strconv"
)

type CLI struct {
	BC *BlockChain
}

const Usage = `
	getBalance	--address	ADDRESS	"获取指定地址的余额"
	printChain						"正向打印区块"
	printChainR						"反向打印区块"
	send FROM TO AMOUNT MINER DATA	"有FROM转AMOUNT给TO，MINER挖矿，同时写入DATA"
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
	//case "addBlock":
	//	fmt.Println("添加新区块")
	//	if len(args) == 4 && args[2] == "--data" {
	//		data := args[3]
	//		c.AddBlock(data)
	//	} else {
	//		fmt.Println("添加区块使用参数不当")
	//		fmt.Println(Usage)
	//	}
	case "send":
		fmt.Println("转账开始...")
		if len(args) != 7 {
			fmt.Println("参数个数错误，请检查")
			fmt.Println(Usage)
		}
		from := args[2]
		to := args[3]
		amount, _ := strconv.ParseFloat(args[4], 64) //字符串转float64
		miner := args[5]
		data := args[6]
		c.Send(from, to, amount, miner, data)
	case "getBalance":
		if len(args) == 4 && args[2] == "--address" {
			addr := args[3]
			c.GetBalance(addr)
		} else {
			fmt.Println("参数格式有误，请检查")
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
