package main

import "fmt"

// Ducker ... интерфейс, к которому надо адоптировать
type Ducker interface {
	quack() string
}

type cater interface {
	Meow() string
}

// cat ... адаптируемая структура
type cat struct {
}

func (c cat) Meow() string {
	return "meow, meow"
}

// DuckAdaptee ... адаптер, который содержит адаптируемую структуру
type DuckAdaptee struct {
	cat *cat
}

// quack ... метод для реализации интерфейса Ducker
func (adapter *DuckAdaptee) quack() string {
	return adapter.cat.Meow()
}

func main() {
	c := &cat{} // создаем кошку

	adapter := &DuckAdaptee{cat: c} // создаем агента

	fmt.Println("adapter agent quack:", adapter.quack()) // проверяем, может ли он квакать

	fmt.Println("pool with ducks: ", DucksOnly(adapter)) // залезаем в пруд
	// => кошка = утка
}

// DucksOnly ... функция, которая принимает в пруд только уток
func DucksOnly(duck Ducker) string {
	return duck.quack()
}
