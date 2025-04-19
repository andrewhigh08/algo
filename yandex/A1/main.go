package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputFilePath := "yandex/A1/input.txt"
	outputFilePath := "yandex/A1/output.txt"

	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла для чтения: %v\n", err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var numbers []int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		for _, part := range parts {
			number, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Ошибка при преобразовании строки в число: %v\n", err)
				return
			}
			numbers = append(numbers, number)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении из файла: %v\n", err)
		return
	}

	if len(numbers) != 4 {
		fmt.Println("В файле не четыре числа")
		return
	}

	result := cntThrees(numbers[0], numbers[1], numbers[2], numbers[3])

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Ошибка при создании файла для записи: %v\n", err)
		return
	}
	defer outputFile.Close()

	fmt.Fprintf(outputFile, "%d\n", result)
}

func cntThrees(p int, v int, m int, q int) int {
	v1, v2 := p-v, p+v
	q1, q2 := m-q, m+q

	if v1 > q1 {
		v1 = q1
	}
	if q2 > v2 {
		v2 = q2
	}

	return absInt(v2) - v1

}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
