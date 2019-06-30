/*
	CS686 Project 2 Blockchain.
	Author: Kei Fukutani
	Date  : February 24, 2019
*/
package p2

import (
	"encoding/json"
	"fmt"
)

// Hold blockchain.
type BlockChain struct {
	// K: Block's height, V: List of Block.
	Chain  map[int32][]Block
	Length int32
}

// Create a new blockchain.
func NewBlockChain() BlockChain {
	bc := BlockChain{}
	bc.Chain = make(map[int32][]Block)
	bc.Length = 0
	return bc
}

// This function takes a height as the argument,
// returns the list of blocks stored in that height or None
// if the height doesn't exist.
func (bc *BlockChain) Get(height int32) []Block {
	list, ok := bc.Chain[height]
	if !ok {
		return nil
	}
	return list
}

// This function takes a block as the argument, use its height
// to find the corresponding list in blockchain's Chain map.
// If the list has already contained that block's hash, ignore it
// because we don't store duplicate blocks; if not, insert the block into the list.
func (bc *BlockChain) Insert(block Block) {
	heightToBlocks := bc.Chain
	height := block.Header.Height
	previousList := heightToBlocks[height-1]
	parentHash := block.Header.ParentHash

	if height < 0 || height > bc.Length+1 {
		return
	}
	if bc.Length == 0 {
		newList := []Block{}
		newList = append(newList, block)
		heightToBlocks[height] = newList
		bc.Length++
		return
	}
	if height == bc.Length+1 {
		// Check if the parent exists.
		if existsInList(parentHash, previousList) {
			newList := []Block{}
			newList = append(newList, block)
			heightToBlocks[height] = newList
			bc.Length++
			return
		}
	}

	list := heightToBlocks[height]
	hash := block.Header.Hash
	// Check if it is in the list.
	if existsInList(hash, list) {
		// Ignore the block.
		return
	}
	// Check if the parent exists.
	if existsInList(parentHash, previousList) {
		list = append(list, block)
		heightToBlocks[height] = list
		return
	}

	return
}

// Take hash of a block and a list of blocks, and check to see if
// the block exists in the list.
func existsInList(hash string, list []Block) bool {
	for _, blockInList := range list {
		if blockInList.Header.Hash == hash {
			return true
		}
	}
	return false
}

// Return JSON strings from BlockChain.
func (bc *BlockChain) EncodeToJson() (string, error) {
	blockJsonList := []BlockJson{}

	heightToBlocks := bc.Chain
	for _, list := range heightToBlocks {
		for _, b := range list {
			bj := buildBjFromBlock(b)
			blockJsonList = append(blockJsonList, bj)
		}
	}

	jsonStrings, err := json.Marshal(blockJsonList)
	if err != nil {
		fmt.Println("error:", err)
		return "", err
	}

	return string(jsonStrings), nil
}

// Take JSON strings and build BlockChain.
func (bc *BlockChain) DecodeFromJson(jsonStrings string) error {
	bytes := []byte(jsonStrings)
	var bjs []BlockJson
	err := json.Unmarshal(bytes, &bjs)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	for _, bj := range bjs {
		block := buildBlockFromBj(bj)
		bc.Insert(block)
	}

	return nil
}
