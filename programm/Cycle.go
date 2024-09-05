package main

import "fmt"

func testCycle() {
	testCycle1()
}

/*
Произведение всех
На вход подается целое число. Вам необходимо вывести произведение всех чисел от 1 до данного числа (включительно).

Например, на вход подается число 5, вам нужно найти - 1*2*3*4*5 = 120
*/
func testCycle1() {
	var a int
	fmt.Scan(&a)
	res := 1
	for i := 1; i <= a; i++ {
		res *= i
	}
	fmt.Println(res)
}
