/*
	For testing MerklePatriciaTrie.go
	Author: Kei Fukutani
	Date  : February 24, 2019
*/
package main

import (
	"cs686/cs686-project-2/p1"
	"cs686/cs686-project-2/p2"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	mpt := p1.MerklePatriciaTrie{}
	mpt.Initial()
	mpt.Insert("hello", "world")
	mpt.Insert("charles", "ge")
	b1 := p2.NewBlock(1, 1234567890, "genesis", mpt)
	b2 := p2.NewBlock(2, 1234567890, b1.Header.Hash, mpt)

	fmt.Println("b1")
	fmt.Printf("Hash: %s\n", b1.Header.Hash)
	fmt.Printf("Timestamp: %d\n", b1.Header.Timestamp)
	fmt.Printf("Height: %d\n", b1.Header.Height)
	fmt.Printf("ParentHash: %s\n", b1.Header.ParentHash)
	fmt.Printf("Size: %d\n", b1.Header.Size)

	fmt.Println("b2")
	fmt.Printf("Hash: %s\n", b2.Header.Hash)
	fmt.Printf("Timestamp: %d\n", b2.Header.Timestamp)
	fmt.Printf("Height: %d\n", b2.Header.Height)
	fmt.Printf("ParentHash: %s\n", b2.Header.ParentHash)
	fmt.Printf("Size: %d\n", b2.Header.Size)

	// Test DecodeFromJson(): a JSON string to a Block.
	jsonString := "{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}"
	block := p2.DecodeFromJson(jsonString)
	// fmt.Printf("block: %+v\n", block)
	fmt.Println("block")
	fmt.Printf("Hash: %s\n", block.Header.Hash)
	fmt.Printf("Timestamp: %d\n", block.Header.Timestamp)
	fmt.Printf("Height: %d\n", block.Header.Height)
	fmt.Printf("ParentHash: %s\n", block.Header.ParentHash)
	fmt.Printf("Size: %d\n", block.Header.Size)
	str, _ := block.Value.Get("hello")
	fmt.Println("str: ", str)
	str, _ = block.Value.Get("charles")
	fmt.Println("str: ", str)

	// Test DecodeFromJson(): a Block to a JSON string.
	jsonString = p2.EncodeToJson(block)
	fmt.Println("jsonString: ", jsonString)

	// Test bc.EncodeToJson: Blocks to JSON strings.
	bc := p2.NewBlockChain()
	bc.Insert(b1)
	bc.Insert(b2)
	jsonStrings, err := bc.EncodeToJson()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println("jsonStrings: ", jsonStrings)
	ret := bc.Length
	fmt.Println(ret) // 2
	if ret != 2 {
		t.Errorf("Expected %d, but was %d", 2, ret)
	}

	// Test cases provided by the instructors.
	fmt.Println("Instructor's test cases.")
	jsonBlockChain := "[{\"hash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"timeStamp\": 1234567890, \"height\": 1, \"parentHash\": \"genesis\", \"size\": 1174, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}, {\"hash\": \"24cf2c336f02ccd526a03683b522bfca8c3c19aed8a1bed1bbc23c33cd8d1159\", \"timeStamp\": 1234567890, \"height\": 2, \"parentHash\": \"3ff3b4efe9177f705550231079c2459ba54a22d340a517e84ec5261a0d74ca48\", \"size\": 1231, \"mpt\": {\"hello\": \"world\", \"charles\": \"ge\"}}]"

	bc = p2.NewBlockChain()
	err = bc.DecodeFromJson(jsonBlockChain)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	jsonNew, err := bc.EncodeToJson()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	var realValue []BlockJson
	var expectedValue []BlockJson
	err = json.Unmarshal([]byte(jsonNew), &realValue)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	err = json.Unmarshal([]byte(jsonBlockChain), &expectedValue)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	if !reflect.DeepEqual(realValue, expectedValue) {
		fmt.Println("=========Real=========")
		fmt.Println(realValue)
		fmt.Println("=========Expcected=========")
		fmt.Println(expectedValue)
		t.Fail()
	}

	fmt.Println("done.")

}

// BlockJson For testing.
type BlockJson struct {
	Height     int32             `json:"height"`
	Timestamp  int64             `json:"timeStamp"`
	Hash       string            `json:"hash"`
	ParentHash string            `json:"parentHash"`
	Size       int32             `json:"size"`
	MPT        map[string]string `json:"mpt"`
}
