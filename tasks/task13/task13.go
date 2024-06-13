package main

import "fmt"

// Swap .. множественное присваивание
func Swap(num1, num2 int) {
	fmt.Println("первое:", num1, "второе:", num2)
	num1, num2 = num2, num1
	fmt.Println("первое:", num1, "второе:", num2)
}

// SwapV2 .. Арифмитический метод
func SwapV2(num1, num2 int) {
	fmt.Println("первое:", num1, "второе:", num2)
	num1 += num2
	num2 = num1 - num2
	num1 -= num2
	fmt.Println("первое:", num1, "второе:", num2)
}

func main() {
	Swap(21, 45)
	Swap(100, 54)
}
