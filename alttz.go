package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение:")
	var input string
	fmt.Scanln(&input)

	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		panic("Ошибка! Неизвестный оператор.")
	}

	parts := strings.SplitN(input, operator, 2)
	if len(parts) != 2 {
		panic("Ошибка! Неверный формат ввода.")
	}

	aStr := strings.TrimSpace(parts[0])
	bStr := strings.TrimSpace(parts[1])

	if !(strings.HasPrefix(aStr, "\"") && strings.HasSuffix(aStr, "\"")) {
		panic("Ошибка! Первым аргументом должна быть строка в кавычках.")
	}

	aStr = aStr[1 : len(aStr)-1]

	if len(aStr) > 10 {
		panic("Ошибка! Длина строки не должна превышать 10 символов.")
	}

	var result string

	switch operator {
	case "+":
		if !(strings.HasPrefix(bStr, "\"") && strings.HasSuffix(bStr, "\"")) {
			panic("Ошибка! Для сложения вторым аргументом должна быть строка в кавычках.")
		}
		bStr = bStr[1 : len(bStr)-1]

		if len(bStr) > 10 {
			panic("Ошибка! Длина строки не должна превышать 10 символов.")
		}

		result = aStr + bStr

	case "-":
		if !(strings.HasPrefix(bStr, "\"") && strings.HasSuffix(bStr, "\"")) {
			panic("Ошибка! Для вычитания вторым аргументом должна быть строка в кавычках.")
		}
		bStr = bStr[1 : len(bStr)-1]

		if strings.Contains(aStr, bStr) {
			result = strings.Replace(aStr, bStr, "", -1)
		} else {
			result = aStr
		}

	case "*":
		b, err := strconv.Atoi(bStr)
		if err != nil || b < 1 || b > 10 {
			panic("Ошибка! Вторым аргументом должно быть число от 1 до 10.")
		}

		result = strings.Repeat(aStr, b)

	case "/":
		b, err := strconv.Atoi(bStr)
		if err != nil || b < 1 || b > 10 {
			panic("Ошибка! Вторым аргументом должно быть число от 1 до 10.")
		}

		if len(aStr)/b == 0 {
			result = ""
		} else {
			result = aStr[:len(aStr)/b]
		}
	default:
		panic("Ошибка! Неизвестная операция.")
	}

	if len(result) > 40 {
		result = result[:40] + "..."
	}

	fmt.Println("Результат:", result)
}
