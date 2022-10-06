// Sajjad Bukhari
// 19I-1686
// CySec-M
// Assignment # 01: Creating a Simple BlockChain

package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Block struct {
	transaction  string
	nonce        int
	hash         string
	previousHash string
}

type BlockList struct {
	blocks []*Block
}

func (b *BlockList) NewBlock(transaction string, nonce int, previousHash string) *Block {
	// A method to add new block. To keep things simple, you could provide a string of your choice as a transaction (e.g., "bob to alice"). Also, use any integer value as a nonce. The CreateHash()
	// method will provide you the block Hash value.

	bl := new(Block)
	bl.transaction = transaction
	bl.nonce = nonce
	bl.previousHash = previousHash

	Str := strconv.Itoa(bl.nonce) + bl.previousHash + bl.transaction
	bl.hash = CalculateHash(Str)

	b.blocks = append(b.blocks, bl)

	return bl
}

func ListBlocks(ls *BlockList) {
	// A method to print all the blocks in a nice format showing block data such as transactions, nonce, previous hash, current block hash.

	counter := 1
	for _, i := range ls.blocks {

		fmt.Printf("%s Transaction %d %s \n", strings.Repeat("=", 35), counter, strings.Repeat("=", 35))
		fmt.Println("Transaction:   ", i.transaction)
		fmt.Println("Nonce:         ", i.nonce)
		fmt.Println("Previous Hash: ", i.previousHash)
		fmt.Println()
		counter++
	}
}

func ChangeBlock(b *BlockList, count int, transaction string) *Block {
	// Function to change block transaction of the given block reference

	count = count - 1
	b.blocks[count].transaction = transaction
	Stri := strconv.Itoa(b.blocks[count].nonce) + b.blocks[count].previousHash + b.blocks[count].transaction
	b.blocks[count].hash = CalculateHash(Stri)
	modBlock := b.blocks[count]

	return modBlock
}

func VerifyChain(b *BlockList) int {
	// Function to verify blockchain in case any chages are made. This can be done in two different ways

	count := 0
	flag := 0

	for count = 0; count < (len(b.blocks) - 1); count++ {
		if b.blocks[count].hash != b.blocks[count+1].previousHash {
			flag = 1
			fmt.Println("Changes detected in Block", count+1)
		}
	}

	return flag
}

func CalculateHash(stringToHash string) string {
	// Function for calculating hash of a block

	hash := sha256.Sum256([]byte(stringToHash))
	Hash := fmt.Sprintf("%x", hash)

	return Hash
}

func main() {
	ls := new(BlockList)

	ls.NewBlock("Sajjad to Muzammil", 1111, "")

	previousHash := strconv.Itoa(ls.blocks[0].nonce) + ls.blocks[0].previousHash + ls.blocks[0].transaction
	calcPH := CalculateHash(previousHash)
	ls.NewBlock("Sajjad to Asfandyaar", 2222, calcPH)

	previousHash = strconv.Itoa(ls.blocks[1].nonce) + ls.blocks[1].previousHash + ls.blocks[1].transaction
	calcPH = CalculateHash(previousHash)
	ls.NewBlock("Sajjad to Haroon", 3333, calcPH)

	previousHash = strconv.Itoa(ls.blocks[2].nonce) + ls.blocks[2].previousHash + ls.blocks[2].transaction
	calcPH = CalculateHash(previousHash)
	ls.NewBlock("Sajjad to Khattak", 4444, calcPH)

	previousHash = strconv.Itoa(ls.blocks[3].nonce) + ls.blocks[3].previousHash + ls.blocks[3].transaction
	calcPH = CalculateHash(previousHash)
	ls.NewBlock("Sajjad to Brigadier", 5555, calcPH)

	ListBlocks(ls)

	block := 3
	NmodBlock := ChangeBlock(ls, block, "Sajjad to Wahab")

	fmt.Println()
	fmt.Printf("%s Changing Block %d %s \n", strings.Repeat("*", 35), block, strings.Repeat("*", 35))
	fmt.Println()

	fmt.Println()
	fmt.Printf("%s Modified Block %s \n", strings.Repeat("-", 35), strings.Repeat("-", 35))
	fmt.Println("Transaction:   ", NmodBlock.transaction)
	fmt.Println("Nonce:         ", NmodBlock.nonce)
	fmt.Println("Previous Hash: ", NmodBlock.previousHash)
	fmt.Println()

	fmt.Println()
	fmt.Printf("%s Verifying Blockchain %s \n", strings.Repeat("=", 35), strings.Repeat("=", 35))
	fmt.Println()

	detected := VerifyChain(ls)
	if detected == 0 {
		fmt.Println("BlockChain Verified!")
		fmt.Println("No Changes Detected!!")
	}
}
