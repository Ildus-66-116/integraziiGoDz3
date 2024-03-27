package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Задание 1. Конвейер
// -Реализуйте паттерн-конвейер:
// Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
// -Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
// -Произведение: следующая горутина умножает квадрат числа на 2.
// -При вводе «стоп» выполнение программы останавливается.

func square(numbers chan int, results chan int) {
	for num := range numbers {
		results <- num * num
		fmt.Printf("Возведение в квадрат: %d\n", num*num)
	}
}
func multiplyByTwo(results chan int) {
	for result := range results {
		fmt.Printf("Умножение на 2: %d\n", result*2)
	}
}

func main() {
	numbers := make(chan int)
	results := make(chan int)
	go square(numbers, results)
	go multiplyByTwo(results)

	for {
		fmt.Println("Введите число")
		var text string
		fmt.Scan(&text)
		if strings.ToLower(text) == "стоп" {
			fmt.Println("До свидания")
			break
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Нужно число!!!")
			continue
		}
		numbers <- num
	}
	close(numbers)
	close(results)
}
