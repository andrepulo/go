package main

import "fmt"

func testInt() {

	// значение по умолчанию
	var num0 int

	// значение при инициализации
	var num1 int = 1

	// пропуск типа переменных
	var num2 = 20
	fmt.Println(num0, num1, num2)

	// короткое объявление переменной
	num := 30
	fmt.Println(num)
	// только для новых переменных
	/*	num := 31

		num += 1
		fmt.Println("+=", num)

		// ++num нет
		num++
		fmt.Println("++", num)
	*/

	// camelCase - принятый стиль
	userIndex := 10
	// under_score - непринятый стиль
	user_index := 10
	fmt.Println(userIndex, user_index)

	// объявление нескольких переменных
	var weight, height int = 10, 20

	// присваивание существующих переменных
	weight, height = 11, 21

	// короткое присваивание
	// хотя бы одна переменная должна быть новой
	weight, age := 12, 22

	fmt.Println(weight, height, age)

	// int - платформозависимый тип
	var i int = 10

	// автоматический выбранный int
	var autoInt = -10

	// int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// платформозависимый тип, 32/64
	var unsignedInt uint = 100500

	// uint8, unt16, uint32, uint64
	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// float32, float64
	var pi float32 = 3.141
	var e = 2.718
	goldenRatio := 1.618

	fmt.Println(pi, e, goldenRatio)

	//bool
	var b bool // false по-умолчанию
	var isOk bool = true
	var success = true
	cond := true

	fmt.Println(b, isOk, success, cond)

	//complex64, complex128
	var c complex128 = -1.1 + 7.12i
	c2 := -1.1 + 7.12i

	fmt.Println(c, c2)

}
