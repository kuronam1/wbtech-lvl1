package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func invert(num int64, i int) int64 {
	i -= 1
	num ^= 1 << i // сдвигаемся на i - 1 количество битов и инвертируем i-ый бит
	return num
}

func main() {
	randomizer := New()
	num := randomizer.Int63()

	fmt.Printf("num in dec: %d, in bits: %064b\n", num, uint64(num))
	fmt.Print("enter a bite's number to invert:")
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Fatal("error while scanning number", err.Error())
	}

	newNum := invert(num, i)
	fmt.Printf("newNum in dec: %d, in bits: %064b", newNum, uint64(newNum))
}

func New() *rand.Rand {
	salt := rand.NewSource(time.Now().UnixNano())
	return rand.New(salt)
}
