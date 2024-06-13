package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// timeout .. возможность отмены горутины, когда пройдет определенное кол-во времени (милисекунд)
func timeout(N int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer cancel()
	operation := make(chan float64, 2)

	go func(ctx context.Context, ch chan float64) {
		for {
			select { // ждем завершения контекста, читаем данные из канала
			case <-ctx.Done(): // жд
				fmt.Println("Done!")
				return
			case msg := <-ch:
				fmt.Println("message: ", msg)
			}
		}
	}(ctx, operation)

	randomizer := New() // создаем рандомайзер для случайных чисел

	ticker := time.NewTicker(time.Second) // создаем тикер чтоб проще было отследить работу
	defer ticker.Stop()                   // освобождаем ресурсы по окончании функции

Loop:
	for {
		select {
		case <-ctx.Done(): // ждем завершения контекста (срабатывания функции отмены)
			close(operation)
			break Loop
		case <-ticker.C:
			msg := randomizer.Float64()
			operation <- msg // посылаем в канал данные
		}
	}
}

// deadline .. почти эквивалентен WithTimeout. Функция отмены сработает, когда наступит метка времени
func deadline(N int) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(N)*time.Second))
	defer cancel()
	operation := make(chan float64, 2)

	go func(ctx context.Context, ch chan float64) {
		for {
			select { // ждем завершения контекста, читаем данные из канала
			case <-ctx.Done():
				fmt.Println("Done!")
				return
			case msg := <-ch:
				fmt.Println("message: ", msg)
			}
		}
	}(ctx, operation)

	randomizer := New()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

Loop:
	for {
		select {
		case <-ctx.Done(): // ждем завершения контекста (срабатывания функции отмены)
			close(operation)
			break Loop
		case <-ticker.C:
			msg := randomizer.Float64()
			operation <- msg // посылаем в канал данные
		}
	}
}

func Cancel(N int) {
	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст для остановки, но, в целом, можно обойтись и без него
	operation := make(chan float64, 2)                      // но это как другой, возможный вариант

	go func(ctx context.Context, ch chan float64) {
		for {
			select { // ждем завершения контекста, читаем данные из канала
			case <-ctx.Done():
				fmt.Println("Done!")
				return
			case msg := <-ch:
				fmt.Println("message: ", msg)
			}
		}
	}(ctx, operation)

	randomizer := New()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	timer := time.NewTimer(time.Duration(N) * time.Second) // саздаем таймер, чтоб отследить момент вермени, когда остановиться
	defer timer.Stop()                                     // освобождаем ресурсы по окончании функции

Loop:
	for {
		select {
		case <-timer.C: // Так же можно канал положить в аргументы к второй горутине и не использовать контекст
			close(operation)
			cancel()
			break Loop
		case <-ticker.C:
			msg := randomizer.Float64()
			operation <- msg
		}
	}
	time.Sleep(time.Second)
}

func main() {
	timeout(6)
	deadline(6)
	Cancel(6)
}

func New() *rand.Rand {
	salt := rand.NewSource(time.Now().Unix())
	return rand.New(salt)
}
