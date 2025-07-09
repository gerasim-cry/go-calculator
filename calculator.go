package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Ошибка: деление на ноль")
		return 0
	}
	return x / y
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите первое число: ")
		aStr, _ := reader.ReadString('\n')
		aStr = strings.TrimSpace(aStr)
		a, err := strconv.ParseFloat(aStr, 64)
		if err != nil {
			fmt.Println("Неверный ввод!")
			continue
		}

		fmt.Println("Введите операцию (+, -, *, /): ")
		op, _ := reader.ReadString('\n')
		op = strings.TrimSpace(op)

		fmt.Println("Введите второе число: ")
		bStr, _ := reader.ReadString('\n')
		bStr = strings.TrimSpace(bStr)
		b, err := strconv.ParseFloat(bStr, 64)
		if err != nil {
			fmt.Println("Неверный ввод!")
			continue
		}

		var result float64

		switch op {
		case "+":
			result = add(a, b)
		case "-":
			result = subtract(a, b)
		case "*":
			result = multiply(a, b)
		case "/":
			result = divide(a, b)
		default:
			fmt.Println("Неизвестная операция")
			continue
		}

		fmt.Println("Резултат:", result)

		fmt.Println("Хотите продолжить? (y/n): ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if strings.ToLower(choice) != "y" {
			fmt.Println("Пока")
			break
		}
	}
}
