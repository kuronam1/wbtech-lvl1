package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"
)

// stopV1 .. возможость остановить горутину через "stop" канал передав туда пустую структуру (структура весит 0, что хорошо сказывается на памяти)
func stopV1() {
	stop := make(chan struct{})

	go func() {
		for {
			select { // ждем сообщения из канала стоп, иначе работаем
			case <-stop:
				fmt.Println("stopping..") // горутина останавливает раоту
				return
			default:
				fmt.Println("working..")
				time.Sleep(500 * time.Millisecond) // имитируем работу горутины
			}
		}
	}()

	time.Sleep(3 * time.Second) // имитируем работу главного потока
	stop <- struct{}{}          // останавливаем горутину
	fmt.Println("waiting for goroutine")
	time.Sleep(1 * time.Second) // дожидаемся окончания горутины
}

// stopV2 .. возможность остановить горутину с помощью контекста (так же можно использовать WithDeadline, WithTimeOut и др)
func stopV2() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select { // ждем сообщения из канала стоп, иначе работаем
			case <-ctx.Done():
				fmt.Println("stopping..") // горутина останавливает раоту
				return
			default:
				fmt.Println("working..")
				time.Sleep(500 * time.Millisecond) // имитируем работу горутины
			}
		}
	}()

	time.Sleep(3 * time.Second) // имитируем работу главного потока
	cancel()                    // останавливаем горутину
	fmt.Println("waiting for goroutine")
	time.Sleep(1 * time.Second) // дожидаемся окончания горутины
}

func stopV3() {
	var flag int32

	go func(stop *int32) {

		for {
			if atomic.LoadInt32(stop) == 1 { // проверяем значение флага (ждет), если значнение == 1 (true) - останавливаемся и выходим
				fmt.Println("stopping")
				return
			}

			fmt.Println("working..")           // имитируем работу горутины
			time.Sleep(500 * time.Millisecond) // дожидаемся окончания горутины
		}
	}(&flag)

	time.Sleep(3 * time.Second) // имитируем работу главного потока
	atomic.AddInt32(&flag, 1)   // останавливаем горутину путем изменения "булевого" флага на true

	fmt.Println("waiting for goroutine")
	time.Sleep(1 * time.Second) // дожидаемся окончания горутины
}

func main() {
	fmt.Println("V1")
	stopV1()
	fmt.Println("V2")
	stopV2()
	fmt.Println("V3")
	stopV3()
}
