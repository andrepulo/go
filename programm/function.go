package main

import "fmt"

func sum(arr []int) int {
	b := 0
	for i := 0; i < len(arr); i++ {
		b += arr[i]
	}
	return b
}

func prod(arr []int) int {
	c := 1
	for i := 0; i < len(arr); i++ {
		c *= arr[i]
	}
	return c
}

func max(arr []int) int {
	d := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > d {
			d = arr[i]
		}
	}
	return d
}
func min(arr []int) int {
	e := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < e {
			e = arr[i]
		}
	}
	return e
}
func main() {
	arr := []int{1, 2, 3, 4, 5}
	result1 := sum(arr)
	result2 := prod(arr)
	result3 := min(arr)
	result4 := max(arr)
	fmt.Println("Сумма всех символов", result1)
	fmt.Println("Произведение всех символов", result2)
	fmt.Println("Минимальное значение", result3)
	fmt.Println("Максимальное значение", result4)

}
