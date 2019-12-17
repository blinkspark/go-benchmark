package main

import (
	"crypto"
	"github.com/blinkspark/go-benchmark/memory"
	"log"

	"github.com/blinkspark/go-benchmark/cpu"
	btime "github.com/blinkspark/go-benchmark/time"
)

type entry struct {
	name   string
	runner btime.Runner
}

var tests = []entry{
	entry{"sha256", sha256},
	entry{"blake2b", blake2b},
	entry{"rsa", rsa},
	entry{"chacha20", chacha20},
	entry{"aes256", aes256},
	entry{"memcopy", memcopy},
}

func main() {
	for _, v := range tests {
		t := btime.Time(v.runner)
		log.Println(v.name, ":", t)
	}
}

func sha256() {
	cpu.Hash(crypto.SHA256, 1000, 1<<20, false)
}

func blake2b() {
	cpu.Hash(crypto.BLAKE2b_256, 1000, 1<<20, false)
}

func rsa() {
	cpu.RSA(1000, 4096, 1<<20, false)
}

func chacha20() {
	cpu.ChaCha20(1000, 1<<20, false)
}

func aes256() {
	cpu.AES(1000, 256, 1<<20, false)
}

func memcopy() {
	memory.MemCopy(10000, 1<<20)
}
