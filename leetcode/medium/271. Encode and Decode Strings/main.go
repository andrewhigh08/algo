// https://leetcode.com/problems/encode-and-decode-strings/description/
// https://www.lintcode.com/problem/659/
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	encodedString := encodeString([]string{"Вася", "code", "love", "you"})
	fmt.Println(encodedString)
	fmt.Println(decodeString(encodedString))

	encodedString2 := encodeString([]string{"we", "say", ":", "yes"})
	fmt.Println(encodedString2)
	fmt.Println(decodeString(encodedString2))
}

// Time & space complexity O(n)
// Временная и пространственная сложность O(n)
func encodeString(inString []string) string {
	// Создаем strings.Builder для эффективной конкатенации строк
	// Create strings.Builder for efficient string concatenation
	var builder strings.Builder

	// Итерируемся по каждой строке в массиве
	// Iterate through each string in the array
	for _, str := range inString {
		// Получаем количество Unicode символов в строке
		// Get the number of Unicode characters in the string
		runeCount := utf8.RuneCountInString(str)

		// Добавляем длину строки в символах
		// Add the string length in characters
		builder.WriteString(strconv.Itoa(runeCount))

		// Добавляем разделитель '#'
		// Add the '#' delimiter
		builder.WriteString(string('#'))

		// Добавляем саму строку
		// Add the string itself
		builder.WriteString(str)
	}

	// Возвращаем закодированную строку
	// Return the encoded string
	return builder.String()
}

// Time & space complexity O(n)
func decodeString(inString string) []string {
	// Слайс для хранения результата декодирования
	// Slice to store the decoding result
	var res []string

	// Переменные для хранения текущего блока и его размера
	// Variables to store the current block and its size
	block, sizeBlock := "", 0

	// Итерируемся по каждому символу в закодированной строке
	// Iterate through each character in the encoded string
	for _, char := range inString {
		// Если символ не '#' и мы еще не начали читать содержимое строки,
		// то это часть числа, указывающего на длину
		// If the character is not '#' and we haven't started reading the string content yet,
		// then it's part of the number indicating length
		if char != 35 && sizeBlock == 0 { // '#' in UTF8
			block += string(char)
			continue
		} else if char == 35 {
			// Если встретили '#', то преобразуем накопленное число в int
			// If we encounter '#', convert the accumulated number to int
			sizeBlock, _ = strconv.Atoi(block)
			block = ""
			continue
		}

		// Добавляем символ к текущему блоку
		// Add the character to the current block
		block += string(char)

		// Уменьшаем счетчик оставшихся символов
		// Decrease the counter of remaining characters
		sizeBlock--

		// Если счетчик достиг нуля, строка закончилась
		// If the counter reached zero, the string has ended
		if sizeBlock == 0 {
			// Добавляем строку в результат
			// Add the string to the result
			res = append(res, block)
			block = ""
		}
	}

	// Возвращаем массив декодированных строк
	// Return the array of decoded strings
	return res
}
