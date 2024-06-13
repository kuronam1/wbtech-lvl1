package main

import (
	"fmt"
	"time"
)

func CustomSleep(dur time.Duration) {
	timer := time.NewTimer(dur)
	defer timer.Stop()

	<-timer.C
}

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
