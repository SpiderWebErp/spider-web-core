package main

import (
	"fmt"
	"github.com/SpiderWebErp/spider-web-core/internal/core"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	blockchain := core.NewBlockchain()

	blockchain.AddBlock([]byte(`{"data":{"balance": 25}`))
	blockchain.AddBlock([]byte(`{"data":{"balance": 26}`))
	blockchain.AddBlock([]byte(`{"data":{"balance": 27}`))
	blockchain.AddBlock([]byte(`{"data":{"balance": 28}`))
	blockchain.AddBlock([]byte(`{"data":{"balance": 29}`))

	if !blockchain.IsValid() {
		panic("")
	}
	str, _ := jsoniter.Marshal(blockchain.GetBlocks())
	fmt.Println(string(str))
}
