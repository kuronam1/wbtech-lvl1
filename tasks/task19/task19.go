package main

import (
	"fmt"
)

var str = "царьпушка"

func V1() {
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

func main() {
	fmt.Println(V2())
	V1()
}
