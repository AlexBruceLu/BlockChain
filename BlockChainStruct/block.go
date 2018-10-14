package main

//定义区块属性
type Block struct {
	PrevHash []byte //前区块哈希值
	Hash []byte //当前哈希值
	Data []byte //所要传递数据
}

//创建区块
func NewBlock(prevHash,data []byte) *Block {
	block:=Block{
		PrevHash:prevHash,
		Hash:[]byte{},
		Data:[]byte(data),
	}
	return &block
}
