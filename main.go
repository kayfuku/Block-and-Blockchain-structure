/*
	CS686 Project 2 main function.
	Author: Kei Fukutani
	Date  : February 24, 2019
*/
package main

import (
	"cs686/cs686-project-2/p1"
	"cs686/cs686-project-2/p2"
	"fmt"
)

func main() {
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

}
