package main

import (
	"fmt"
	"sync"
)

var numbers = []int{2, 4, 6, 8, 10}

func V1() {
	wg := sync.WaitGroup{}
	wg.Add(len(numbers))

	defer fmt.Println()
	for _, num := range numbers {
		go func(num int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Print(num*num, " ")
		}(num, &wg)
	}
	wg.Wait()
}

func V2() {
	wg := sync.WaitGroup{}
	wg.Add(len(numbers))

	for idx := range numbers {
		//Проблема решена в Go 1.22, передаю в параметры функции значение из-за warning от IDE
		go func(slice []int, i int) {
			defer wg.Done()

			slice[i] *= slice[i]
		}(numbers, idx)
	}
	wg.Wait()

	fmt.Println(numbers)
}

func main() {
	V1()
	V2()
}
