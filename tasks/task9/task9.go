package main

import (
	"fmt"
	"math/rand"
	"time"
)

func read(out chan int) {
	randomizer := New()
	numbers := randomizer.Perm(50) // подготавливаем случайный массив из 50-ти символов

	for i := 1; i <= len(numbers)-1; i++ {
		out <- numbers[i] // Читаем данные из массива
	}
	close(out) // Закрываем канал после прочтения, чтобы дать понять горутине, которая читает этот канал, что чтение прекращено
}

func square(in, out chan int) {
	for num := range in { // читаем данные в цикле, чтобы не залететь на панику в случае закрытия канала
		out <- num * num // запускаем число после преобразования дальше по конвееру
	}
	close(out) // закрываем канал, чтоб дать понять main-функции, что данная горутина закончила работу
}

func main() {
	readChan := make(chan int)
	squareChan := make(chan int)
	go read(readChan)
	go square(readChan, squareChan)

	counter := 0
	for num := range squareChan { // читаем данные в цикле, чтобы не залететь на панику в случае закрытия канала
		counter++
		fmt.Println("номер -", counter, "число -", num) // выводим полученные значения из конвеера
	}
}

func New() *rand.Rand {
	salt := rand.NewSource(time.Now().Unix())
	return rand.New(salt)
}
