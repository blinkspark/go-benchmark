package main

import (
	"crypto"
	"go-benchmark/cpu"
	btime "go-benchmark/time"
	"log"
)

func main() {
	t1 := btime.Time(func() {
		cpu.Hash(crypto.SHA256, 1<<10, 1<<20, false)
		// cpu.RSA(1000, 4096, 1<<20, false)
		// cpu.ChaCha20(1000, 1<<20, false)
		// cpu.AES(1000, 256, 1<<20, false)
		// cpu.Test()
	})

	t2 := btime.Time(func() {
		cpu.Hash(crypto.BLAKE2b_256, 1<<10, 1<<20, false)
	})

	log.Println("Sha-256:", t1)
	log.Println("Blake2b:", t2)
}
