package main

import (
	"fmt"
	"sort"
)

var subsequence = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

func main() {
	mp := make(map[int][]float64)

	for _, temp := range subsequence {
		roundElem := round(temp)
		mp[roundElem] = append(mp[roundElem], temp)
	}

	keys := make([]int, 0, len(mp))
	for key := range mp {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	//Вывод в нужном виде
	for i := 0; i < len(keys); i++ {
		fmt.Printf("%d{", keys[i])
		for idx, value := range mp[keys[i]] {
			if idx == len(mp[keys[i]])-1 {
				fmt.Printf("%.1f", value)
				continue
			}

			fmt.Printf("%.1f, ", value)
		}
		fmt.Print("}, ")
	}
}

func round(num float64) int {
	switch {
	case num > 0:
		return roundPositive(num)
	case num < 0:
		return roundNegative(num)
	default:
		return 0
	}
}

//Идея заключается в округлении каждого числа и записи его как ключ map (или добавлении в массив по ключу map)

// roundPositive ... округление положительных чисел
func roundPositive(num float64) int {
	result := int(num)
	cnt := 0
	for result > 10 {
		cnt++
		result /= 10
	}
	result *= cnt * 10
	return result
}

// roundNegative ...  округление отрицательных чисел
func roundNegative(num float64) int {
	result := int(num)
	cnt := 0
	for result < -10 {
		cnt++
		result /= 10
	}
	result *= cnt * 10
	return result
}
