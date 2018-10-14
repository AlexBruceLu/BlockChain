package main

import "fmt"

func main() {
	//1.每出21万个块励衰减一次
	//2.当奖励必须大于0，小于0的时候发行完毕
	total := 0.0 	//总的发行量
	blocks := 21.0  //每次衰减时，产生区块的数量，单位：万个
	reward := 50.0	//初始的挖矿奖励
	for reward > 0 {
		sum := reward * blocks
		reward *= 0.5
		total += sum
	}
	fmt.Printf("比特币总发行量 %v%s\n",total,"万枚")
}
