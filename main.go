package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	private, public, err := generateKey(1024) //key size
	if err != nil {
		panic(err)
	}
	n := big.NewInt(0)
	n.Mul(public, private)

	f := big.NewInt(0)
	f.Mul(big.NewInt(0).Sub(private, big.NewInt(1)), big.NewInt(0).Sub(public, big.NewInt(1)))

	e := generateE(f)

	d := big.NewInt(0)
	d.ModInverse(e, f)

	fmt.Printf("\nprivate key: %v\n\n", private)
	fmt.Printf("public key: %v\n\n", public)
	fmt.Printf("n: %v\n\n", n)
	fmt.Printf("f: %v\n\n", f)
	fmt.Printf("e: %v\n\n", e)
	fmt.Printf("d: %v\n\n", d)

	secretMessage := "Привет, Настя! Как там с лабой?))"
	fmt.Printf("Text: %s\n\n", secretMessage)

	encryptedMessage := encrypt(secretMessage, e, n)

	fmt.Printf("Cipher Text: %v\n\n", encryptedMessage)

	decryptedMessage := decrypt(encryptedMessage, d, n)

	fmt.Printf("Decipher Text: %s\n", decryptedMessage)
}

func generateE(f *big.Int) *big.Int {
	e := big.NewInt(2)
	for e.Cmp(f) == -1 {
		gcb := big.NewInt(0)
		gcb.GCD(nil, nil, e, f)

		if gcb.Cmp(big.NewInt(1)) == 0 {
			return e
		}
		e.Add(e, big.NewInt(1))
	}
	return nil
}

func generateKey(bytesize int) (*big.Int, *big.Int, error) {
	var p, q *big.Int
	p, err := rand.Prime(rand.Reader, bytesize)
	if err != nil {
		return nil, nil, err
	}
	for {
		q, err = rand.Prime(rand.Reader, bytesize)
		if err != nil {
			return nil, nil, err
		}
		if p != q {
			return p, q, err
		}
	}
}

func encrypt(msg string, e *big.Int, n *big.Int) []*big.Int {
	runes := []rune(msg)
	var cipher []*big.Int
	for _, v := range runes {
		cipher = append(cipher, big.NewInt(0).Exp(big.NewInt(int64(v)), e, n))
	}
	return cipher
}

func decrypt(encr_msg []*big.Int, d *big.Int, n *big.Int) string {
	var msg_bytes []rune
	for _, v := range encr_msg {
		msg_bytes = append(msg_bytes, int32(big.NewInt(0).Exp(v, d, n).Int64()))
	}
	return string(msg_bytes)
}
