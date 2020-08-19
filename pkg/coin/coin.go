package coin

import (
	"time"

	bc "github.com/Namchee/simple-blockchain-demo/pkg/blockchain"
)

type namcheeCoin struct {
	Chain bc.BlockChain
}

func Init(difficulty int) *namcheeCoin {
	blockChain := bc.NewBlockChain(4)

	return &namcheeCoin{*blockChain}
}

func (nc *namcheeCoin) AddNewData(data map[string]interface{}) {
	block := bc.NewBlock(
		nc.Chain.GetBlockLength(),
		int(time.Now().Unix()),
		data,
		"",
	)

	nc.Chain.AddNewBlock(block)
}

func (nc *namcheeCoin) ReplaceChain(chain *bc.BlockChain) {
	if chain.CheckBlockValidity() &&
		nc.Chain.GetBlockLength() < chain.GetBlockLength() {
		nc.Chain = *chain
	}
}
