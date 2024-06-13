package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(slice []int, low, high int) int {
	pivot := slice[high]
	i := low

	for j := low; j < high; j++ {
		if slice[j] < pivot {
			slice[j], slice[i] = slice[i], slice[j]
			i++
		}
	}

	slice[i], slice[high] = slice[high], slice[i]
	return i
}

func quicksort(slice []int, low, high int) {
	if low < high {
		pivot := partition(slice, low, high)
		quicksort(slice, low, pivot-1)
		quicksort(slice, pivot+1, high)
	}
}

func main() {
	slice := generateSlice(100)

	fmt.Printf("unsorted: %v\n", slice)
	quicksort(slice, 0, len(slice)-1)
	fmt.Printf("sorted: %v", slice)
}

func generateSlice(size int) []int {
	randomizer := New()
	return randomizer.Perm(size)
}

func New() *rand.Rand {
	salt := rand.NewSource(time.Now().UnixNano())
	return rand.New(salt)
}
