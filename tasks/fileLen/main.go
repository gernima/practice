package main

import (
	"fmt"
	"os"
)

func CountLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	lineCount := 0
	var char byte
	var buf [1]byte

	for {
		_, err := file.Read(buf[:])
		if err != nil {
			break
		}
		char = buf[0]
		if char == '\n' {
			lineCount++
		}
	}

	_, err = file.Seek(-1, 2)
	if err == nil {
		file.Read(buf[:])
		if buf[0] != '\n' {
			lineCount++
		}
	}

	return lineCount, nil
}

func main() {
	count, err := CountLines("file.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Количество строк: %d\n", count)
}
