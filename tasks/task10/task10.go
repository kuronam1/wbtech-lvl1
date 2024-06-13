package main

import "fmt"

var subsequence = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

func main() {
	mp := make(map[int][]float64)

	for _, temp := range subsequence {
		mp[round(temp)] = append(mp[round(temp)], temp)
	}

	fmt.Println(mp)
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

// roundPositive .. округление положительных чисел
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

// roundNegative .. округление отрицательных чисел
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
