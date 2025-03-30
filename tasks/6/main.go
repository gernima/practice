package main

import (
	"fmt"
	"os"
)

func countLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return lineCount, nil
}

func main() {
	count, err := countLines("file.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Количество строк: %d\n", count)
}
