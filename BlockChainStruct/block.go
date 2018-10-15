package main

import (
	"bytes"
	"encoding/binary"
	"time"
)

//定义区块属性
type Block struct {
	Version    uint64 //版本号
	PrevHash   []byte //前区块哈希值
	MerKelRoot []byte //梅歇尔根
	TimeStamp  uint64 //时间戳
	Difficulty uint64 //难度值
	Nonce      uint64 //随机数，挖矿要找的数
	Hash       []byte //当前哈希值
	Data       []byte //所要传递数据
}

//创建区块
func NewBlock(prevHash []byte, data string) *Block {
	block := Block{
		Version:    00,
		PrevHash:   prevHash,
		MerKelRoot: []byte{},
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0, //写的随机无效值
		Nonce:      0, //写的随机无效值
		Hash:       []byte{},
		Data:       []byte(data),
	}
	//block.SetHash()
	//利用工作量证明计算哈希值

	pow:=NewProofOfWork(&block)

	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce
	return &block
}

//实现uint64转字符切片的辅助函数
func Uint64ToByte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

//生成哈希、
//func (block *Block) SetHash() {
//	//var blockInfo []byte
//	//blockInfo = append(blockInfo, block.Hash...)
//	//blockInfo = append(blockInfo, block.PrevHash...)
//	//blockInfo = append(blockInfo, block.Data...)
//	//blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
//	//blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
//	//blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
//	//blockInfo = append(blockInfo, Uint64ToByte(block.TimeStamp)...)
//	//blockInfo = append(blockInfo, block.MerKelRoot...)
//	//--------------------代码优化------------------------
//	tmp := [][]byte{
//		block.Hash,
//		Uint64ToByte(block.Version),
//		Uint64ToByte(block.Nonce),
//		Uint64ToByte(block.TimeStamp),
//		Uint64ToByte(block.Difficulty),
//		block.MerKelRoot,
//		block.Data,
//		block.PrevHash,
//	}
//	blockInfo := bytes.Join(tmp, []byte{})
//	hash := sha256.Sum256(blockInfo)
//	block.Hash = hash[:] //定长数组转切片
//}
