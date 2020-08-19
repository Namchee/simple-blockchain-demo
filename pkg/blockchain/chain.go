package blockchain

import "time"

type BlockChain struct {
	difficulty int
	blocks     []block
}

func NewBlockChain(difficulty int) *BlockChain {
	var blocks []block
	blocks = append(blocks, *createGenesisBlock())

	blockchain := &BlockChain{difficulty, blocks}

	return blockchain
}

func createGenesisBlock() *block {
	data := map[string]interface{}{
		"hello": "world",
	}

	return NewBlock(
		0,
		int(time.Now().Unix()),
		data,
		"",
	)
}

func (bc *BlockChain) GetBlockLength() int {
	return len(bc.blocks)
}

func (bc *BlockChain) CheckBlockValidity() bool {
	for i := 1; i < len(bc.blocks); i++ {
		curBlock := bc.blocks[i]
		prevBlock := bc.blocks[i-1]

		if curBlock.hash != curBlock.calculateHash() ||
			curBlock.precedingHash != prevBlock.hash {
			return false
		}
	}

	return true
}

func (bc *BlockChain) AddNewBlock(b *block) {
	b.precedingHash = bc.blocks[len(bc.blocks)-1].hash

	b.mineBlock(bc.difficulty)
	bc.blocks = append(bc.blocks, *b)
}
