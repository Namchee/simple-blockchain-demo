package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type block struct {
	index         int
	timestamp     int
	hash          string
	data          map[string]interface{}
	precedingHash string
	nonce         int
}

func NewBlock(index int, timestamp int, data map[string]interface{}, precedingHash string) *block {
	newBlock := &block{
		index,
		timestamp,
		"",
		data,
		precedingHash,
		0,
	}

	newBlock.hash = newBlock.calculateHash()

	return newBlock
}

func (b *block) calculateHash() string {
	stringify, _ := json.Marshal(b.data)

	data := []byte(strconv.Itoa(b.index))
	data = append(data, []byte(strconv.Itoa(b.timestamp))...)
	data = append(data, stringify...)
	data = append(data, []byte(b.precedingHash)...)
	data = append(data, []byte(strconv.Itoa(b.nonce))...)

	h := sha256.New()
	h.Write(data)

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func (b *block) mineBlock(difficulty int) {
	for b.hash[:difficulty] != strings.Join(make([]string, difficulty+1), "0") {
		b.nonce++
		b.hash = b.calculateHash()
	}
}
