package main

import "fmt"

func TestString() {
	/*
		// пустая строка
		var str string

		// со спец символами
		var hello string = "Привет\n\t"

		// без спецсимволов
		var world string = `Мир\n\t`
	*/
	helloWorld := "Привет Мир"
	//конкатенация строк
	andGoodMorning := helloWorld + " и доброе утро!"

	fmt.Println(andGoodMorning)

	//строки неизменяемы
	//cannot assign to helloWorld[0]

}
