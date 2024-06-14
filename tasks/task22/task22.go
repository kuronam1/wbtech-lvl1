package main

import (
	"fmt"
	"math/big"
)

var a string = "24000000000000000000"
var b string = "23000000000000000000"

// division .. деление
func division(num1, num2 *big.Int) *big.Int {
	result := new(big.Int)
	return result.Div(num1, num2)
}

// multiplication .. умножение
func multiplication(num1, num2 *big.Int) *big.Int {
	result := new(big.Int)
	return result.Mul(num1, num2)
}

// addition .. сложение
func addition(num1, num2 *big.Int) *big.Int {
	result := new(big.Int)
	return result.Add(num1, num2)
}

// subtraction .. вычитание
func subtraction(num1, num2 *big.Int) *big.Int {
	result := new(big.Int)
	return result.Sub(num1, num2)
}

func main() {
	num1, num2 := new(big.Int), new(big.Int)
	num1.SetString(a, 10)
	num2.SetString(b, 10)
	fmt.Printf("num1: %v, num2: %v\n", num1, num2)
	fmt.Printf("division num1/num2: %v\n", division(num1, num2)) // результат округлен до целых
	fmt.Printf("multiplication num1/num2: %v\n", multiplication(num1, num2))
	fmt.Printf("addition num1/num2: %v\n", addition(num1, num2))
	fmt.Printf("substration num1/num2: %v\n", subtraction(num1, num2))
}
