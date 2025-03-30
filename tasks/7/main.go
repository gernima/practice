package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func printPerson(p Person) string {
	return fmt.Sprintf("Имя: %s, Возраст: %d", p.Name, p.Age)
}

func main() {
	jsonFile, err := os.ReadFile("file.json")
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	var persons []Person
	err = json.Unmarshal(jsonFile, &persons)
	if err != nil {
		fmt.Println("Ошибка при разборе JSON:", err)
		return
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer outputFile.Close()

	for _, person := range persons {
		result := printPerson(person)
		fmt.Println(result)

		_, err := outputFile.WriteString(result + "\n")
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
}
