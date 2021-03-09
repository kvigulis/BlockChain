package blockchain

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int
}

// CreateBlock - returns a pointer to a new block object
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock - add a now block into the blockchain
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis - creates the first block
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain - Intialises BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
