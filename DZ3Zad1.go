package main

import (
	"bufio"
	"fmt"
	"os"
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
		fmt.Printf("Квадрат: %d\n", num*num)
	}
}
func multiplyByTwo(results chan int) {
	for result := range results {
		fmt.Printf("Прозведение: %d\n", result*2)
	}
}

func main() {
	numbers := make(chan int)
	results := make(chan int)
	go square(numbers, results)
	go multiplyByTwo(results)

	fmt.Println("Введите число")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToLower(input) == "стоп" {
			fmt.Println("До свидания")
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Нужно число!!!")
			continue
		}
		numbers <- num
	}

	close(numbers)
	close(results)
}
