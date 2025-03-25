package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func PrintPerson(p Person) {
	fmt.Printf("Имя: %s, Возраст: %d\n", p.Name, p.Age)
}

func main() {
	p := Person{Name: "Иван", Age: 30}
	PrintPerson(p)
}
