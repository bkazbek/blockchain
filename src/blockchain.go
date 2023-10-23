package src

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index        uint64
	Timestamp    time.Time
	Proof        uint64
	PreviousHash string
}

type Blockchain struct {
	Chain []Block
}

func NewBlockchain() *Blockchain {
	chain := make([]Block, 0)
	blockchain := &Blockchain{
		Chain: chain,
	}
	blockchain.CreateBlock(1, "0")
	return blockchain
}

func (b *Blockchain) CreateBlock(proof uint64, previousHash string) Block {
	block := Block{
		Index:        uint64(len(b.Chain)),
		Timestamp:    time.Now(),
		Proof:        proof,
		PreviousHash: previousHash,
	}
	b.Chain = append(b.Chain, block)
	return block
}

func (b *Blockchain) GetPreviousBlock() Block {
	return b.Chain[len(b.Chain)-1]
}

func (b *Blockchain) ProofOfWork(previousProof uint64) uint64 {
	newProof := uint64(1)
	checkProof := false
	for !checkProof {
		hashedOperation := fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.Itoa(int(math.Pow(float64(previousProof), 2)-math.Pow(float64(newProof), 2))))))
		if strings.HasPrefix(hashedOperation, "0000") {
			checkProof = true
		} else {
			newProof += 1
		}
	}
	return newProof
}

func (b *Blockchain) Hash(block Block) string {
	jsonData, _ := json.Marshal(block)
	return fmt.Sprintf("%x", sha256.Sum256(jsonData))
}

func (b *Blockchain) IsValidChain() bool {
	previousBlock := b.Chain[0]
	blockIndex := 1
	for blockIndex < len(b.Chain) {
		block := b.Chain[blockIndex]
		if block.PreviousHash != b.Hash(previousBlock) {
			return false
		}
		previousProof := previousBlock.Proof
		proof := block.Proof
		hashedOperation := fmt.Sprintf("%x", sha256.Sum256([]byte(strconv.Itoa(int(math.Pow(float64(previousProof), 2)-math.Pow(float64(proof), 2))))))
		if !strings.HasPrefix(hashedOperation, "0000") {
			return false
		}
		previousBlock = block
		blockIndex++
	}
	return true
}
