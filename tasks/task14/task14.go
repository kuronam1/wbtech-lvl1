package main

import (
	"fmt"
	"reflect"
)

func main() {
	determineType(14)
	determineType("value")
	determineType(false)
	determineType(make(chan string))
	determineTypeV2(make(chan int))
}

func determineType(value interface{}) {
	switch reflect.TypeOf(value).Kind() {
	case reflect.Chan:
		fmt.Println("channel")
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	default:
		fmt.Println("unknown type")
	}
}

func determineTypeV2(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("int") //после определения мы можем сохранить флаг (в случае, если это структура)
	case string:
		fmt.Println("string") //или реализовать какой-либо метод связанный с этим типом данных
	case bool:
		fmt.Println("bool")
	case chan int: // тип канала может быть любым.
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("unknown type")
	}
}
