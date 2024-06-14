package main

import (
	"fmt"
	"time"
)

// CustomSleep ... устанавливаем таймер и блокируем горутину на это время
// можно было еще просто через time.After, но логика там почти одна и та же
func CustomSleep(dur time.Duration) {
	timer := time.NewTimer(dur) // создаем таймер для внутреннего отсчета времени
	defer timer.Stop()

	<-timer.C
}

// Здесь мы просто представляем возможность выбрать время на которое уснуть. А потом выводим его для наглядности
func main() {
	fmt.Println("enter time for sleep in format: \"HHhSSs\"")

	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		fmt.Println("err while scanning")
		return
	}

	dur, err := time.ParseDuration(str)
	if err != nil {
		fmt.Println("err while converting")
		return
	}

	fmt.Println("sleeping...")
	start := time.Now()
	CustomSleep(dur)
	end := time.Since(start)
	fmt.Println("woke up. We were sleeping for", end)
}
