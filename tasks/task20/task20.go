package main

import (
	"fmt"
	"slices"
	"strings"
)

// revertWords ... Вызовы отложенных функций выполняются в порядке LIFO (Last in First Out). но мы не сможем вернуть значение из функции
func revertWords(str string) {
	strConvert := strings.Fields(str)

	defer fmt.Println()
	for _, word := range strConvert {
		defer fmt.Print(word, " ")
	}
}

// revertWordsV2 ... разворачиваем слайс с помощью пакета slices
func revertWordsV2(str string) string {
	strConvert := strings.Fields(str)
	slices.Reverse(strConvert)
	return strings.Join(strConvert, " ")
}

// revertWordsV3 .. не очень эффективный, но простой способ
func revertWordsV3(str string) string {
	strConvert := strings.Fields(str)

	sl := make([]string, 0)
	for i := len(strConvert) - 1; i != -1; i-- {
		sl = append(sl, strConvert[i])
	}
	return strings.Join(sl, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println(str)
	revertWords(str)
	v2 := revertWordsV2(str)
	fmt.Println(v2)

	v3 := revertWordsV3(str)
	fmt.Println(v3)
}
