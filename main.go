package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstDigitString string
	var secondDigitString string
	var operator string
	var isArabicNumbers = false
	fmt.Println("Введите пример в формате *число* *операция* *число*")
	scanResult, err := fmt.Scanln(&firstDigitString, &operator, &secondDigitString)
	checkStringFormat(scanResult, err)
	firstDigit := stringToInteger(firstDigitString, &isArabicNumbers, true)
	secondDigit := stringToInteger(secondDigitString, &isArabicNumbers, false)
	solve(firstDigit, operator, secondDigit)
}
func checkStringFormat(n int, err error) {
	if n != 3 || err != nil {
		panic("Введённый пример не соответсвует формату")
	}
} // checks the format of string
func stringToInteger(digitString string, isArabicNumbers *bool, first bool) int {
	digit, getDigit := strconv.Atoi(digitString)
	if getDigit != nil {
		digit = romanDigitStringToInt(digitString)
		if *isArabicNumbers {
			panic("Используются одновременно разные системы счисления")
		}
	} else {
		if !*isArabicNumbers && !first { // if it is the second number it must be from the same digits family as the first one
			panic("Используются одновременно разные системы счисления")
		}
		*isArabicNumbers = true
	}
	checkDigit(digit)
	return digit
} // converts string to integer
func checkDigit(digit int) {
	if digit > 10 {
		panic(fmt.Sprintf("Число %d > 10", digit))
	} else if digit < 1 {
		panic(fmt.Sprintf("Число %d < 1", digit))
	}
} // checks if digit can be used in the task
func solve(firstDigit int, operator string, secondDigit int) {
	var answer int
	switch operator {
	case "+":
		answer = firstDigit + secondDigit
		break
	case "-":
		answer = firstDigit - secondDigit
		break
	case "*":
		answer = firstDigit * secondDigit
		break
	case "/":
		answer = firstDigit / secondDigit
		break
	default:
		panic("Невозможно распознать знак операции")
		return
	}
	print(answer)
} // solves the task
func romanDigitStringToInt(digitString string) int {
	var result int = 0
	for i := 0; i < len(digitString); i++ {
		fS := getRomanDigitFromString(digitString[i])
		if i+1 < len(digitString) {
			sS := getRomanDigitFromString(digitString[i+1])
			if fS >= sS {
				result += fS
			} else {
				result += sS - fS
				i++
			}
		} else {
			result += fS
		}
	}
	return result
} // converts roman number to int
func getRomanDigitFromString(c uint8) int {
	switch c {
	case 'I':
		return 1
		break
	case 'V':
		return 5
		break
	case 'X':
		return 10
		break
	case 'L':
		return 50
		break
	case 'C':
		return 100
		break
	case 'D':
		return 500
		break
	case 'M':
		return 1000
		break
	default:
		panic(fmt.Sprintf("Невозможно преобразовать число %d", c))
	}
	return -1
} // converts current roman number to int
