package ecb

type ECB struct {
	block     crypto.Block
	blockSize int
}

func (e *ECB) BlockSize() {
	return e.blockSize
}

func (e *ECB) CryptBlocks(dst, src []byte) {
	if len(src) % e.blockSize {
		panic("ecb: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecb: output smaller than input")
	}
	if len(src) == 0 {
		return
	}

}

var _ crypto.BlockMode = &ECB{}

func New(block crypto.Block) *ECB {
	return &ECB{
		block:     block,
		blockSize: block.BlockSize(),
	}
}
