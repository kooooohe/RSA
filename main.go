package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	var p, q int
	fmt.Println("enter two defferent primary numbers")
	fmt.Println("like 17 19")
	fmt.Scan(&p, &q)

	var planText int64
	fmt.Println("enter plan text")
	fmt.Printf("[planText < %d] \n", p*q)
	fmt.Scan(&planText)
	fmt.Printf("plan text: %d\n", planText)

	N := int64(p * q)
	L := culcLeastCommonMultiple(p-1, q-1)

	// gcd(E,L) = 1
	E := makePublicKey(L)
	fmt.Printf("public key: %d\n", E)

	// E*D mod L = 1
	D := int64(makePrivateKey(E, int64(L)))
	fmt.Printf("private key: %d\n", D)

	// planText^E mod N
	cryptgram := new(big.Int).Exp(big.NewInt(planText), big.NewInt(E), big.NewInt(N))
	fmt.Printf("crypt gram: %d\n", cryptgram.Int64())

	// cryptgram^D mod N
	res := new(big.Int).Exp(cryptgram, big.NewInt(D), big.NewInt(N))
	fmt.Printf("plan text: %d\n", res.Int64())
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
