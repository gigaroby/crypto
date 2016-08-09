package ecb

import "crypto/cipher"

type ECB struct {
	block     cipher.Block
	blockSize int

	encrypt bool
}

func (e *ECB) BlockSize() int {
	return e.blockSize
}

func (e *ECB) CryptBlocks(dst, src []byte) {
	if len(src)%e.blockSize != 0 {
		panic("ecb: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("ecb: output smaller than input")
	}
	if len(src) == 0 {
		return
	}

	for i := 0; i < len(dst)/e.blockSize; i++ {
		if e.encrypt {
			e.block.Encrypt(src[i*e.blockSize:], dst[i*e.blockSize:])
		} else {
			e.block.Decrypt(src[i*e.blockSize:], dst[i*e.blockSize:])
		}
	}
}

func NewEncrypter(block cipher.Block) *ECB {
	return &ECB{
		block:     block,
		blockSize: block.BlockSize(),

		encrypt: true,
	}
}

func NewDecrypter(block cipher.Block) *ECB {
	return &ECB{
		block:     block,
		blockSize: block.BlockSize(),

		encrypt: false,
	}
}

var _ cipher.BlockMode = &ECB{}
