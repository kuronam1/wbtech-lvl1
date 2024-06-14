package main

import (
	"fmt"
	"slices"
)

var str = "царьпушка"

func V1() {
	defer fmt.Println()
	for _, val := range str {
		defer fmt.Print(string(val))
	}
}

func V2() string {
	runes := []rune(str)
	result := make([]rune, len(runes))
	for i := len(runes) - 1; i != -1; i-- {
		result = append(result, runes[i])
	}

	return string(result)
}

func V3() string {
	runes := []rune(str)
	slices.Reverse(runes)
	return string(runes)
}

func main() {
	V1()
	fmt.Println(V2())
	fmt.Println(V3())
}
