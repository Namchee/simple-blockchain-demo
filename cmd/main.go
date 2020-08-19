package main

import (
	"github.com/Namchee/simple-blockchain-demo/pkg/coin"
)

func main() {
	coin := coin.Init(4)

	coin.AddNewData(map[string]interface{}{"test": "test"})
}
