package main

import (
	"fmt"
	"math/rand"
	"time"
)

func DeleteElemV1(i int, slice []int) []int {
	return append(slice[0:i], slice[i+1:]...)
}

func DeleteElemV2(i int, slice []int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	slice := generateSlice(20)
	fmt.Printf("slice: %v\n", slice)
	fmt.Println("enter a number of element that needs to be deleted")
	var i int
	_, err := fmt.Scan(&i)
	if err != nil {
		fmt.Println("error while scanning")
		return
	}
	newSlice := DeleteElemV1(i, slice)
	fmt.Printf("V1 new slice: %v\n", newSlice)

	newSlice = DeleteElemV2(i, slice)
	fmt.Printf("V2 new slice: %v\n", newSlice)
}

func generateSlice(size int) []int {
	salt := rand.NewSource(time.Now().UnixNano())
	random := rand.New(salt)

	slice := make([]int, 0)
	for i := 0; i < size; i++ {
		slice = append(slice, random.Intn(100))
	}

	return slice
}
