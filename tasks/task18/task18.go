package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct {
	value int
}

func V2(amount int) {
	type counterV2 struct {
		counter int64
	}

	cnt := counterV2{}
	wg := sync.WaitGroup{}
	wg.Add(amount)

	for i := 0; i < amount; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&cnt.counter, 1) // атомики безопасно использовать в конкурентной среде
		}()
	}
	wg.Wait()

	fmt.Println(cnt.counter)
}

func V1(amount int) {
	cnt := &counter{}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(amount)
	for i := 0; i < amount; i++ {
		go func(count *counter) {
			defer wg.Done()

			mu.Lock()
			count.value++
			mu.Unlock()
		}(cnt)
	}
	wg.Wait()

	fmt.Println(cnt.value)
}

func main() {
	V1(1000)
	V2(1000)
}
