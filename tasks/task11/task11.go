package main

import (
	"fmt"
	"slices"
)

var first = []string{"Andrey", "Vasya", "Marat", "Ivan", "Dmitriy", "Maksim"}
var second = []string{"Sasha", "Maksim", "Nikolay", "Andrey", "Ivan"}

// Mножество - набор или свовокупность неповторяющихся элементов

func V1() []string {
	var result []string //предствляем множество в виде масива
	for _, value := range first {
		if slices.Contains(second, value) { //проверяем, чтоб в массиве не было повторящихся элементов
			result = append(result, value)
		}
	}

	return result
}

func V2() []string {
	mp := make(map[string]struct{}) //Ключ в map не может повторяться, а значит map может реализовать определение множества
	for _, elem := range first {
		mp[elem] = struct{}{}
	}

	var res []string

	for _, elem := range second {
		if _, inMap := mp[elem]; inMap {
			res = append(res, elem)
		}
	}

	return res
}

func Eleven() {
	fmt.Println(V1())
	fmt.Println(V2())
}
