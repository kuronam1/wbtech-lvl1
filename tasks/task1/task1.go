package main

import (
	"fmt"
	"io"
	"os"
)

type Action struct {
	Human
	Direct io.Writer
}

func (a *Action) Burk() {
	_, err := fmt.Fprintln(a.Direct, "гав-гав!")
	if err != nil {
		panic(err)
	}
}

func (a *Action) Talk() {
	_, err := fmt.Fprintln(a.Direct, "говорю..")
	if err != nil {
		panic(err)
	}
}

func (a *Action) Meow() {
	_, err := fmt.Fprintln(a.Direct, "мяу-мяу")
	if err != nil {
		panic(err)
	}
}

type Human struct {
	Height float64
	Weight float64
}

func (h *Human) Eat() {
	h.Weight += 1.5
}

func (h *Human) Train() {
	h.Weight -= 2
}

func (h *Human) Grow() {
	h.Height += 1
}

func main() {
	human := Human{
		Height: 150,
		Weight: 50,
	}

	action := Action{
		Human:  human,
		Direct: os.Stdout,
	}
	fmt.Println(action.Weight)
	action.Eat()
	fmt.Println(action.Weight)
	action.Train()
	fmt.Println(action.Weight)

	fmt.Println(action.Height)
	action.Grow()
	fmt.Println(action.Height)

	action.Burk()
	action.Meow()
	action.Talk()
}
