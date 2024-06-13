package main

import "fmt"

type Ducker interface {
	quack() string
}

type cater interface {
	Meow() string
}

type cat struct {
}

func (c cat) Meow() string {
	return "meow, meow"
}

type DuckAdaptee struct {
	cat *cat
}

func (adapter *DuckAdaptee) quack() string {
	return adapter.cat.Meow()
}

func main() {
	c := &cat{}

	adapter := &DuckAdaptee{cat: c}

	fmt.Println("adapter agent quack:", adapter.quack())

	fmt.Println("pool with ducks: ", DucksOnly(adapter))
}

func DucksOnly(duck Ducker) string {
	return duck.quack()
}
