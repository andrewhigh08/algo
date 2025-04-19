package main

import (
	"strconv"
	"strings"
)

// Encode превращает массив строк в одну строку
func Encode(strs []string) string {
	var sb strings.Builder

	for _, str := range strs {
		length := strconv.Itoa(len(str))   // Преобразуем длину строки в строковое представление
		sb.WriteString(length + "#" + str) // Добавляем длину, разделитель и саму строку
	}

	return sb.String()
}

// Decode превращает закодированную строку обратно в массив строк
func Decode(s string) []string {
	var result []string
	i := 0

	for i < len(s) {
		// Найдем позицию разделителя '#'
		j := i
		for s[j] != '#' {
			j++
		}

		// Извлекаем длину строки
		length, _ := strconv.Atoi(s[i:j])
		// Извлекаем строку указанной длины
		result = append(result, s[j+1:j+1+length])
		// Сдвигаем указатель дальше на длину строки и на сам разделитель
		i = j + 1 + length
	}

	return result
}

func main() {
	// Пример 1
	strs1 := []string{"neet", "code", "love", "you"}
	encoded1 := Encode(strs1)
	println("Encoded:", encoded1)

	decoded1 := Decode(encoded1)
	println("Decoded:")
	for _, str := range decoded1 {
		println(str)
	}

	// Пример 2
	strs2 := []string{"we", "say", ":", "yes"}
	encoded2 := Encode(strs2)
	println("Encoded:", encoded2)

	decoded2 := Decode(encoded2)
	println("Decoded:")
	for _, str := range decoded2 {
		println(str)
	}
}
