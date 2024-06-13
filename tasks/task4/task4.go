package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type worker struct { //создаем сущность worker, которая будет работать
	Number   int
	WorkChan chan int
}

func (w *worker) Run(ctx context.Context) { // непосредственный запуск работы воркера
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d ended work\n", w.Number)
			return
		case msg := <-w.WorkChan:
			fmt.Printf("worker: %d, msg: %d\n", w.Number, msg)
		}
	}
}

func main() {
	randomizer := New()                                     // создаем рандомайзер для случайных данных
	ctx, cancel := context.WithCancel(context.Background()) // создаем ручное управление горутинами
	fmt.Println("enter a number of workers:")
	var workersAmount int
	_, err := fmt.Scan(&workersAmount) // выбираем количество рабочих
	if err != nil {
		fmt.Println("error while converting")
		return
	}

	data := make(chan int, 10)
	for i := 0; i < workersAmount; i++ { // Запускаем воркеров работать
		wk := worker{
			Number:   i + 1,
			WorkChan: data,
		}
		go wk.Run(ctx)
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM) // ждем notify от os
Loop:
	for {
		select {
		case <-interrupt: // если программа прервана - завершаем рабочий день всех рабочих
			cancel()
			break Loop
		default: // посылаем данные на обработку
			num := randomizer.Intn(100)
			data <- num
		}
	}

	fmt.Println("ending program...")
	time.Sleep(1 * time.Second) // дожидаемся завершения горутин
}

func New() *rand.Rand {
	salt := rand.NewSource(time.Now().UnixNano())
	return rand.New(salt)
}
