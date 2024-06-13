package main

import (
	"fmt"
	"math/rand"
)

var justString string

func createHugeString(size int) string {
	result := make([]rune, size)
	randomizer := New()
	minimum := 97
	maximum := 122

	for idx := range result {
		result[idx] = int32(randomizer.Intn(maximum-minimum+1) + minimum)
	}

	return string(result)
}

func someFuncV1() {
	v := createHugeString(1 << 10)
	temp := []rune(v)
	justString = string(temp[:1<<10])
	fmt.Println(justString)
}

func someFuncV2() {
	v := createHugeString(100)
	temp := []rune(v)
	justString = string(temp[:100])
	fmt.Println(justString)
}

func main() {
	someFuncV1()
	someFuncV2()
}

func New() *rand.Rand {
	salt := rand.NewSource(42)
	return rand.New(salt)
}
