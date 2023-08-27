package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка в прочтении ввода: %s %s", input, err)
	}

	if strings.ContainsAny(input, "+-*/") == false {
		fmt.Println("Ошибка: введена неправильная строка")
		return
	}

	numbers, operation := splitIntoTwoNumbersAndOperation(input)
	if operation == "" {
		fmt.Println("Ошибка: введена неправильная строка")
		return
	}

	a, isRoman1 := parseNum(numbers[0])
	b, isRoman2 := parseNum(numbers[1])

	if isRoman1 != isRoman2 {
		fmt.Println("Ошибка: введены числа разных типов")
		return
	}

	var result int

	switch operation {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	// Выводим результат
	if isRoman1 {
		if result <= 0 {
			fmt.Println("Ошибка: результат меньше единицы")
			return
		}
		fmt.Println(toRoman(result))
		return
	}
	fmt.Println(result)

}

func parseNum(numStr string) (int, bool) {
	romanMap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	if num, ok := romanMap[numStr]; ok {
		return num, true

	}

	return 0, false

}

func toRoman(num int) string {
	romanMap := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	result := ""
	for value, symbol := range romanMap {
		for num >= value {
			result += symbol
			num -= value
		}
	}
	return result
}

func splitIntoTwoNumbersAndOperation(input string) ([]string, string) {

	switch {
	case strings.Contains(input, "+"):
		input = strings.Replace(input, " ", "", -1)
		input = strings.Replace(input, "\r\n", "", -1)
		nums := strings.Split(input, "+")
		return nums, "+"
	case strings.Contains(input, "-"):
		input = strings.Replace(input, " ", "", -1)
		input = strings.Replace(input, "\r\n", "", -1)
		nums := strings.Split(input, "-")
		return nums, "-"
	case strings.Contains(input, "/"):
		input = strings.Replace(input, " ", "", -1)
		input = strings.Replace(input, "\r\n", "", -1)
		nums := strings.Split(input, "/")
		return nums, "/"
	case strings.Contains(input, "*"):
		input = strings.Replace(input, " ", "", -1)
		input = strings.Replace(input, "\r\n", "", -1)
		nums := strings.Split(input, "*")
		return nums, "*"
	}
	return nil, ""
}
