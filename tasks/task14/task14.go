package main

import "fmt"

func main() {
	determineType(14)
	determineType("value")
	determineType(false)
	determineType(make(chan int))
}

func determineType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int") //после опреления мы можем сохранить флаг (в случае, если это структура)
	case string:
		fmt.Println("string") //или реализовать какой-либо метод связанный с этим типом данных
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("channel")
	}
}
