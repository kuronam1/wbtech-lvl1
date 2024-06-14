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

// CheckUniquenessV1 .. способ через массив
func CheckUniquenessV1(s string) bool {
	s = strings.ToLower(s) // приводим все к одному регистру
	var slice []rune
	for _, symbol := range s {
		if slices.Contains(slice, symbol) { // проверяем - есть ли в массиве такой символ и если нет добавляем
			return false // иначе можем завершать выполнение
		}
		slice = append(slice, symbol)
	}
	return true
}

// CheckUniquenessV2 .. способ через map
func CheckUniquenessV2(s string) bool {
	s = strings.ToLower(s) // приводим все к одному регистру
	mp := make(map[rune]rune)
	for _, symbol := range s { // по ключу записываем все символы в map
		mp[symbol] = symbol
	}
	return len(mp) == len(s) // если len(mp) == len(s) - все символы уникальны
}

// CheckUniquenessV3 .. способ через строку - логика как с массивом
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
