package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(elem int, slice []int) (int, bool) {
	lowPtr := 0
	highPtr := len(slice) - 1

	for lowPtr <= highPtr {
		median := (lowPtr + highPtr) / 2

		switch {
		case slice[median] == elem:
			return median, true
		case slice[median] < elem:
			lowPtr = median + 1
		case slice[median] > elem:
			highPtr = median - 1
		}
	}

	if slice[lowPtr] != elem || lowPtr == len(slice) || highPtr == 0 {
		return 0, false
	}

	return lowPtr, true
}

func main() {
	slice := generateSlice(20)

	fmt.Printf("slice: %v\n", slice)
	idx, inSlice := binarySearch(44, slice)
	if inSlice == false {
		fmt.Println("element not in slice")
		return
	}

	fmt.Println("elem idx =", idx)
}

func generateSlice(size int) []int {
	salt := rand.NewSource(time.Now().Unix())
	random := rand.New(salt)

	slice := make([]int, 0)
	for i := 0; i < size; i++ {
		slice = append(slice, random.Intn(100))
	}
	//slice = append(slice, 44)

	sort.Ints(slice)

	return slice
}
