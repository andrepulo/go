package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	testSlice1()
	testSlice2()
	testSlice3()
}

/*Напишите программу, которая:

Создает слайс (срез) из трех чисел.
Определяет функцию, которая находит максимальное число в слайсе.
Вызывает эту функцию и выводит результат.
Использует цикл for для вывода всех чисел из слайса.*/

func testSlice1() {
	// Создайте слайс из трех чисел
	numbers := []int{15, 5, 10}
	// Вызовите функцию findMax и выведите результат
	max := findMax(numbers)
	fmt.Printf("Максимальное значение: %d\n", max)
	// Используйте цикл for для вывода всех чисел
	for _, num := range numbers {
		fmt.Println(num)
	}
}
func findMax(nums []int) int {
	// Ваш код здесь
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max

}

/*Напишите программу, которая:

Создает слайс из 10 случайных чисел от 1 до 100.
Определяет функцию, которая находит сумму всех четных чисел в слайсе.
Вызывает эту функцию и выводит результат.
Выводит все числа из слайса и отмечает, какие из них четные.*/

func testSlice2() {
	rand.Seed(time.Now().UnixNano())

	// Ваш код здесь
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = rand.Intn(100) + 1
	}
	fmt.Println("Вывод всех чисел")
	for _, num := range numbers {
		if num%2 == 0 {
			fmt.Printf("%d(четное)\n", num)
		} else {
			fmt.Printf("%d(нечетное)\n", num)
		}
	}

	evenSum := sumEven(numbers)
	fmt.Printf("Сумма четных чисел %d\n", evenSum)
}

// Функция для суммирования четных чисел
func sumEven(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum

}

/*
Задача:

Реализуйте функцию processNumbers.
Для четных чисел используйте evenHandler (в данном случае, удвоение).
Для нечетных чисел используйте oddHandler (в данном случае, возведение в квадрат).
Верните новый слайс с обработанными числами.
*/

func processNumbers(numbers []int, evenHandler func(int) int, oddHandler func(int) int) []int {
	// Ваш код здесь
	result := make([]int, len(numbers))
	for i, num := range numbers {
		if num%2 == 0 {
			result[i] = evenHandler(num)
		} else {
			result[i] = oddHandler(num)
		}
	}
	return result
}

func testSlice3() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	double := func(x int) int {
		return x * 2
	}

	square := func(x int) int {
		return x * x
	}

	result := processNumbers(numbers, double, square)
	fmt.Println(result)
}

/*
Напишите функцию applyOperations, которая принимает слайс целых чисел
и произвольное количество функций-операций.
Функция должна применить все операции к каждому числу в слайсе по порядку
и вернуть новый слайс с результатами.
*/

func applyOperations(numbers []int, operations ...func(int) int) []int {
	result := make([]int, len(numbers))
	for i, num := range numbers {
		temp := num
		for _, operation := range operations {
			temp = operation(temp)

		}
		result[i] = temp
	}
	return result
}

func testSlice4() {
	numbers := []int{1, 2, 3, 4, 5}

	double := func(x int) int {
		return x * 2

	}

	addOne := func(x int) int {
		return x + 1
	}

	square := func(x int) int {
		return x * x

	}
	result := applyOperations(numbers, double, addOne, square)
	fmt.Println(result)

}
