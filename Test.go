package _Golang_

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func romToAr(romanNum string) int {
	var romanNumerals = map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
		"VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	return romanNumerals[romanNum]
}

func arToRom(arabicNum int) string {
	var result strings.Builder

	romanSymbols := []struct {
		value  int
		symbol string
	}{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	for _, symbol := range romanSymbols {
		for arabicNum >= symbol.value {
			result.WriteString(symbol.symbol)
			arabicNum -= symbol.value
		}
	}

	return result.String()
}

func isRoman(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) ||
			(unicode.ToUpper(char) != 'I' &&
				unicode.ToUpper(char) != 'V' &&
				unicode.ToUpper(char) != 'X') {
			return false
		}
	}
	return true
}

func main() {
	var a, b string
	var s string
	var err = ""
	fmt.Println("Введите пример для решения " + "Введите выражение:")
	fmt.Fscanln(os.Stdin, &a, &s, &b, &err)

	if err != "" {
		panic("Вы ввели что-то не так\"число операнд число\"")
		return
	}
	if isRoman(a) != isRoman(b) {
		panic("Выдача паники, так как используются одновременно разные системы счисления. ")
		return
	}

	var numA, numB int
	if isRoman(a) {
		numA = romToAr(a)
	} else {
		numA, _ = strconv.Atoi(a)
	}
	if isRoman(b) {
		numB = romToAr(b)
	} else {
		numB, _ = strconv.Atoi(b)
	}

	if numB > 10 || numB < 1 || numA > 10 || numA < 1 {
		panic("Используйте только числа от 1 до 10")
		return
	}

	var result int

	switch s {
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "*":
		result = numA * numB
	case "/":
		result = numA / numB
	default:
		panic("Выдача паники, так как строка не является математической операцией.")
		return
	}

	if isRoman(a) && isRoman(b) {
		var e = arToRom(result)
		if e == "" {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.\n")
			return
		}
		fmt.Println("Результат:", arToRom(result))
	} else {
		fmt.Println("Результат:", result)
	}
}
