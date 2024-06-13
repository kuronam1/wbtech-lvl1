package main

import (
	"fmt"
	"sync"
)

var numbers = []int{2, 4, 6, 8, 10}

func V1() {
	wg := sync.WaitGroup{}
	wg.Add(len(numbers))
	for _, num := range numbers {
		go func(num int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Print(num*num, " ")
		}(num, &wg)
	}
	wg.Wait()
}

func V2() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(numbers))

	for idx := range numbers {
		go func(num *int, wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()

			*num *= *num
		}(&numbers[idx], &wg, &mu)
	}
	wg.Wait()

	fmt.Println(numbers)
}

func main() {
	V1()
	V2()
}
