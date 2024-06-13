package main

import (
	"fmt"
	"sync"
)

func V2(n int) {
	var wg sync.WaitGroup
	mp := make(map[string]int)
	mu := sync.Mutex{}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(value int) {
			defer wg.Done()                       //говорим, что горутина завершила работа (счетчик работаующих горутин понижается на 1)
			key := fmt.Sprintf("worker%d", value) // произвольный ключ
			mu.Lock()                             // блокируем для всех горутин возможность работать с map
			mp[key] = value                       // записываем значение по созданному ключу
			mu.Unlock()                           //
		}(i)
	}
	wg.Wait()

	fmt.Println(mp)
}

func V1(n int) {
	var (
		mp sync.Map       // используем map из пакета sync, она дает возможность безопасно работать в конкурентной среде
		wg sync.WaitGroup // без мьютексов
	)

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(num int) {
			defer wg.Done()                     //говорим, что горутина завершила работа (счетчик работаующих горутин понижается на 1)
			key := fmt.Sprintf("worker%d", num) // произвольный ключ
			mp.Store(key, num)                  //сохраняем значение по ключу
		}(i)
	}

	wg.Wait()
	mp.Range(func(key, value any) bool { //выводим значения map в определенном формате, применяя функцию printf ко всем парам
		fmt.Printf("key: %s, value: %v;\n", key, value)
		return true
	})
}

func main() {
	V1(10)
	V2(10)
}
