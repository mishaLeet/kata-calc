package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var arrRoman [10]string = [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var arrArab [10]string = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

// Функция, конвертирующая строку в число
func stringToInt(str string) int {
	number, _ := strconv.Atoi(str)
	return number
}

// Функция, конвертирующая число в строку
func intToString(num int) string {
	return strconv.Itoa(num)
}

// Перевод римских в арабские
func romanNum(roman string) int {
	var arab int
	for i := 0; i < 10; i++ {
		if roman == arrRoman[i] {
			arab = i + 1
			break
		}
	}
	return arab
}

// Ответ в виде римских
func romanAnswer(a, b int) string {
	var dozens, units string
	switch a {
	case 1:
		dozens = "X"
	case 2:
		dozens = "XX"
	case 3:
		dozens = "XXX"
	case 4:
		dozens = "XL"
	case 5:
		dozens = "L"
	case 6:
		dozens = "LX"
	case 7:
		dozens = "LXX"
	case 8:
		dozens = "LXXX"
	case 9:
		dozens = "XC"
	case 10:
		dozens = "C"
	}

	switch b {
	case 1:
		units = "I"
	case 2:
		units = "II"
	case 3:
		units = "III"
	case 4:
		units = "IV"
	case 5:
		units = "V"
	case 6:
		units = "VI"
	case 7:
		units = "VII"
	case 8:
		units = "VIII"
	case 9:
		units = "IX"
	}
	return dozens + units
}

// Функция с математическими операциями для калькулятора
func operations(a, b int, op string) (result int) {
	switch {
	case op == "+":
		result = a + b
	case op == "-":
		result = a - b
	case op == "*":
		result = a * b
	case op == "/":
		result = a / b
	}
	return result
}

// Функция обработки ошибок
func error(a, op, b string) {
	// Проверка на числа от 1 до 10
	var newA, newB int
	for i := 0; i < 10; i++ {
		if a == arrRoman[i] {
			newA = romanNum(a)
		}
		if b == arrRoman[i] {
			newB = romanNum(b)
		}
		if a == arrArab[i] {
			newA = stringToInt(a)
		}
		if b == arrArab[i] {
			newB = stringToInt(b)
		}
	}
	if newA < 1 || newA > 10 || newB < 1 || newB > 10 {
		fmt.Print(errors.New("Вне диапазона"))
		os.Exit(1)
	}
	// Проверка операции
	if op != "+" && op != "-" && op != "*" && op != "/" {
		fmt.Print(errors.New("Не является математической операцией"))
		os.Exit(1)
	}
	// Проверка на отрицательные числа в римской системе
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (a == arrRoman[i]) && (b == arrRoman[j]) {
				if (operations(romanNum(a), romanNum(b), op)) < 1 {
					fmt.Print(errors.New("В римской системе нет отрицательных чисел!"))
					os.Exit(1)
				}
			}
		}
	}
	// Проверка на один тип
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if (a == arrArab[i] && b == arrRoman[j]) || (a == arrRoman[i] && b == arrArab[j]) {
				fmt.Print(errors.New("Используются разные системы счисления"))
				os.Exit(1)
			}
		}
	}
}

// Функция вычислений
func calc(a, op, b string) (result string) {
	var x, y, dozens, units int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if a == arrArab[i] && b == arrArab[j] {
				x = stringToInt(a)
				y = stringToInt(b)
				result = intToString(operations(x, y, op))
				break
			}
			if a == arrRoman[i] && b == arrRoman[j] {
				x = romanNum(a)
				y = romanNum(b)
				dozens = operations(x, y, op) / 10
				units = operations(x, y, op) % 10
				result = romanAnswer(dozens, units)
				break
			}
		}
	}
	return result
}
func main() {
	var a, op, b string
	fmt.Print("Введите выражение: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	num := strings.Fields(text)
	if len(num) < 3 {
		fmt.Print(errors.New("Строка не является математической операцией"))
		os.Exit(1)
	}
	a = num[0]
	op = num[1]
	b = num[2]
	error(a, op, b)
	fmt.Print(calc(a, op, b))
}
