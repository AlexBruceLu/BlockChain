package main

//定义一个区块链数组
type BlockChain struct {
	Blocks []*Block
}

//定义一个区块链
func NewBlockChain() *BlockChain {
	//创建创世块，添加到区块链中
	genesisBlock := GenesisBlock()
	return &BlockChain{
		Blocks: []*Block{genesisBlock},
	}
}

//定义一个创世块
func GenesisBlock() *Block {
	return NewBlock([]byte{}, "这是第一个创世块")
}

//添加区块
func (bc *BlockChain) AddBlocks(data string) {
	//获取最后一个区块，得到前区块的哈希值
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	prevHash := lastBlock.Hash
	block := NewBlock(prevHash, data)
	//添加新区块到区块链上
	bc.Blocks = append(bc.Blocks, block)
}
