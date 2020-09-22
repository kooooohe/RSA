package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	// primary numbers
	p, q := 17, 19

	N := int64(p * q)
	println(N)
	L := culcLeastCommonMultiple(p-1, q-1)
	println(L)

	// gcd(E,L) = 1
	E := makePublicKey(L)
	fmt.Printf("public key: %d\n", E)

	// E*D mod L = 1
	D := int64(makePrivateKey(E, int64(L)))
	fmt.Printf("private key: %d\n", D)

	// p < N
	planText := int64(123)
	fmt.Printf("plan text : %d\n", planText)

	// planText^E mod N
	cryptgram := new(big.Int).Exp(big.NewInt(planText), big.NewInt(E), big.NewInt(N))
	fmt.Printf("crypt gram key: %d\n", cryptgram.Int64())

	// cryptgram^D mod N
	res := new(big.Int).Exp(cryptgram, big.NewInt(D), big.NewInt(N))
	fmt.Printf("plan text : %d\n", res.Int64())
}

func makePublicKey(l int) int64 {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(l)
	for greatestCommonDivisor(r, l) != 1 {
		r = rand.Intn(l)
	}
	return int64(r)
}

func makePrivateKey(e int64, l int64) int64 {
	i := int64(2)
	for i*e%l != 1 {
		i++
	}
	return i
}

func culcLeastCommonMultiple(a, b int) int {
	c := a * b
	if a < b {
		tmp := a
		a = b
		b = tmp
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return c / b
}

func greatestCommonDivisor(a, b int) int {
	if a < b {
		tmp := a
		a = b
		b = tmp
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return b

}
