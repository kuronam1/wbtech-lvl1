package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		panic(err)
	}

	fmt.Println(CheckUniquenessV1(str))
	fmt.Println(CheckUniquenessV2(str))
	fmt.Println(CheckUniquenessV3(str))
}

func CheckUniquenessV1(s string) bool {
	s = strings.ToLower(s)
	var slice []rune
	for _, symbol := range s {
		if slices.Contains(slice, symbol) {
			return false
		}
		slice = append(slice, symbol)
	}
	return true
}

func CheckUniquenessV2(s string) bool {
	s = strings.ToLower(s)
	mp := make(map[rune]rune)
	for _, symbol := range s {
		mp[symbol] = symbol
	}
	return len(mp) == len(s)
}

func CheckUniquenessV3(s string) bool {
	s = strings.ToLower(s)
	var res string
	for _, symbol := range s {
		if strings.Contains(res, string(symbol)) {
			return false
		}
		res += string(symbol)
	}
	return true
}
