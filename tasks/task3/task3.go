package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var text = "2,4,6,8,10"

func V1() {

	start := time.Now()
	subsequence := strings.Split(text, ",")
	resChan := make(chan int)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(subsequence))
		for idx := range subsequence {
			//Проблема решенша в Go task1.22, передаю в параметры функции значение из-за warning от IDE
			go func(value string, ch chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				num, err := strconv.Atoi(value)
				if err != nil {
					fmt.Println("error while converting type")
					return
				}

				ch <- num * num

			}(subsequence[idx], resChan, &wg)
		}
		wg.Wait()
		close(resChan)
	}()

	var sum int
	for SqNum := range resChan {
		sum += SqNum
	}
	fmt.Printf("sum = %d\n", sum)
	fmt.Println(time.Since(start).Nanoseconds())
}

type Worker struct {
	sync.Mutex
	value string
}

func (w *Worker) AddAndSquare(wg *sync.WaitGroup, result *int) {
	defer wg.Done()
	num, err := strconv.Atoi(w.value)
	if err != nil {
		fmt.Println("error while converting type")
		return
	}

	w.Lock()
	*result += num * num
	w.Unlock()
}

func V2() {
	start := time.Now()
	subsequence := strings.Split(text, ",")
	var (
		result int
		wg     sync.WaitGroup
	)
	wg.Add(len(subsequence))

	for idx := range subsequence {
		worker := Worker{value: subsequence[idx]}
		go worker.AddAndSquare(&wg, &result)
	}
	wg.Wait()

	fmt.Println(result)
	fmt.Println(time.Since(start).Nanoseconds())
}

func main() {
	V1()
	V2()
}
