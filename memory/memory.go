package memory

import (
	"crypto/rand"
)

func MemCopy(rounds, datasize int) {
	buffer := make([]byte, datasize)
	target := make([]byte, datasize)
	rand.Read(buffer)
	for i := 0; i < rounds; i++ {
		copy(target, buffer)
	}
}
