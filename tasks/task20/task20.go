package main

import (
	"fmt"
	"strings"
)

func revertWords(str string) {
	strConvert := strings.Fields(str)

	for _, word := range strConvert {
		defer fmt.Print(word, " ")
	}
}

func revertWordsV2(str string) string {
	strConvert := strings.Fields(str)

	sl := make([]string, len(strConvert))
	for _, word := range strConvert {
		defer appendToSlice(&sl, word)
	}

	return strings.Join(sl, " ")
}

func revertWordsV3(str string) string {
	strConvert := strings.Fields(str)

	sl := make([]string, 0)
	for i := len(strConvert) - 1; i != -1; i-- {
		sl = append(sl, strConvert[i])
	}
	return strings.Join(sl, " ")
}

func appendToSlice(slice *[]string, elem string) {
	*slice = append(*slice, elem)
}

func main() {
	str := "snow dog sun"
	revertWords(str)
	fmt.Println()
	v2 := revertWordsV2(str)
	fmt.Println(v2)

	v3 := revertWordsV3(str)
	fmt.Println(v3)
}
