package main

import (
	"fmt"
	"slices"
)

var slice = []string{"cat", "cat", "dog", "cat", "tree"}

// Mножество - набор или свовокупность неповторяющихся элементов

func V1() {
	var result []string //предствляем множество в виде масива
	for _, element := range slice {
		if !slices.Contains(result, element) { //проверяем, чтоб в массиве не было повторящихся элементов
			result = append(result, element)
		}
	}

	fmt.Println(result)
}

func V2() {
	result := make(map[string]struct{}) //Ключ в map не может повторяться, а значит map может реализовать определение множества

	for _, elem := range slice {
		result[elem] = struct{}{}
	}

	fmt.Print("[")
	for key := range result {
		fmt.Print(key, " ")
	}
	fmt.Print("]")
}

func main() {
	V1()
	V2()
}
