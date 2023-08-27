package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка в прочтении ввода: %s %s", input, err)
	}

	if strings.ContainsAny(input, "+-*/") == false {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	numbers, operation := splitIntoTwoNumbersAndOperation(input)
	if operation == "" {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}
	if len(numbers) > 2 {
		fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		return
	}

	a, isRoman1 := parseNum(numbers[0])
	b, isRoman2 := parseNum(numbers[1])

	if isRoman1 != isRoman2 {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
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
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
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
	num, err := strconv.Atoi(numStr)
	if err != nil {
		log.Printf("Не смог конвертировать стринг %s в инт, err: %s", numStr, err)
		return 1, false
	}

	return num, false

}

func toRoman(num int) string {
	romanMap := map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}

	var keys []int
	for k := range romanMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	result := ""
	for _, value := range keys {
		symbol := romanMap[value]
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
