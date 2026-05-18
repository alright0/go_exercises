package generator

import "crypto/md5"

type Entropy struct {
	Hash     [16]byte
	bitIndex int
}

func (e *Entropy) NextBool() bool {
	bitsPerByte := 8

	if e.bitIndex >= len(e.Hash)*bitsPerByte {
		e.Hash = md5.Sum(e.Hash[:])
		e.bitIndex = 0
	}

	byteIndex := e.bitIndex / bitsPerByte
	bitOffset := e.bitIndex % bitsPerByte

	b := (e.Hash[byteIndex] >> bitOffset) & 1

	e.bitIndex++

	return b != 0
}

func (e *Entropy) getHash(seed []byte) [16]byte {
	return md5.Sum(seed)
}

func NewEntropy(seed string) *Entropy {
	return &Entropy{Hash: md5.Sum([]byte(seed))}
}
