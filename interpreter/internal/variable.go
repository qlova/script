package internal

type Address interface {
	Address() int
}

type Variable struct {
	BlockPointer *Block
}
