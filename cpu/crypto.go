package cpu

import (
	"crypto"
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
	"log"
	mrand "math/rand"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/sha3"
)

func Hash(hashType crypto.Hash, rounds int, dataSize int64, verbose bool) {
	buffer := make([]byte, dataSize)
	h := newHash(hashType)
	mrand.Seed(0)
	for i := 0; i < rounds; i++ {
		mrand.Read(buffer)
		_, err := h.Write(buffer)
		if err != nil {
			log.Panic(err)
		}
		result := h.Sum(nil)
		if verbose {
			fmt.Println(result)
		}
		h.Reset()
	}
}

func newHash(hType crypto.Hash) hash.Hash {
	key := make([]byte, 64)
	rand.Read(key)
	switch hType {
	case crypto.SHA224:
		return sha256.New224()
	case crypto.SHA256:
		return sha256.New()
	case crypto.SHA384:
		return sha512.New384()
	case crypto.SHA512:
		return sha512.New()
	case crypto.SHA3_224:
		return sha3.New224()
	case crypto.SHA3_256:
		return sha3.New256()
	case crypto.SHA3_384:
		return sha3.New384()
	case crypto.SHA3_512:
		return sha3.New512()
	case crypto.BLAKE2s_256:
		h, _ := blake2s.New256(nil)
		return h
	case crypto.BLAKE2b_256:
		h, _ := blake2b.New256(nil)
		return h
	case crypto.BLAKE2b_384:
		h, _ := blake2b.New384(nil)
		return h
	case crypto.BLAKE2b_512:
		h, _ := blake2b.New512(nil)
		return h
	}
	return nil
}

func RSA(rounds, keyLength int, dataSize int64, verbose bool) {
	priv, _ := rsa.GenerateKey(rand.Reader, keyLength)
	buffer := make([]byte, dataSize)
	h := newHash(crypto.BLAKE2b_256)
	mrand.Seed(0)
	for i := 0; i < rounds; i++ {
		mrand.Read(buffer)
		data, _ := rsa.EncryptOAEP(h, rand.Reader, &priv.PublicKey, buffer, nil)
		if verbose {
			fmt.Println(data)
		}
	}
}

func ChaCha20(rounds int, dataSize int64, verbose bool) {
	key := make([]byte, chacha20.KeySize)
	rand.Read(key)
	nonce := make([]byte, chacha20.NonceSize)
	rand.Read(nonce)
	mrand.Seed(0)
	c, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	target := make([]byte, dataSize)
	buffer := make([]byte, dataSize)

	for i := 0; i < rounds; i++ {
		mrand.Read(buffer)
		c.XORKeyStream(target, buffer)
	}
}

func AES(rounds, bits int, dataSize int64, verbose bool) {
	key := make([]byte, bits/8)
	rand.Read(key)
	c, _ := aes.NewCipher(key)
	mrand.Seed(0)
	buffer := make([]byte, dataSize)
	dst := make([]byte, dataSize)
	for i := 0; i < rounds; i++ {
		mrand.Read(buffer)
		c.Encrypt(dst, buffer)
	}
}

func Test() {
	key := [32]byte{}
	fmt.Println(cap(key[:11]))
	fmt.Println(len(key[:11]))
	fmt.Println(key)
	rand.Read(key[:])
	fmt.Println(key)

	// m:= poly1305.Sum()
}
