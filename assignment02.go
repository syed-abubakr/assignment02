package assignment02

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce       int
	BlockData   []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

func GenerateNonce(blockData []Transaction) int {
	nonce := rand.Intn(1000)
	return nonce
}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
			blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	nonce := GenerateNonce(blockData)
	currentHash := CalculateHash(blockData, nonce)
	PrevHash := ""
	if chainHead != nil {
		PrevHash = chainHead.CurrentHash
	}
	block := Block{Nonce: nonce, BlockData: blockData, PrevPointer: chainHead, PrevHash: PrevHash, CurrentHash: currentHash}
	return &block
}

func ListBlocks(chainHead *Block) {
	currentNode := chainHead
	for currentNode != nil {
		data := *currentNode
		DisplayTransactions(data.BlockData)
		currentNode = data.PrevPointer
	}
}

func DisplayTransactions(blockData []Transaction) {
	for i := 1; i < len(blockData); i++ {
		print("ID: ", blockData[i].TransactionID, " ")
		print("Sender: ", blockData[i].Sender, " ")
		print("Receiver: ", blockData[i].Receiver, " ")
		print("Amount: ", blockData[i].Amount, "\n")
	}
}

func TransactionID(sender string, receiver string, amount int) string {
	dataString := sender + receiver + string(amount)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	id := TransactionID(sender, receiver, amount)
	transaction := Transaction{TransactionID: id, Sender: sender, Receiver: receiver, Amount: amount}
	return transaction
}
