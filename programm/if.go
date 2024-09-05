package main

import "fmt"

func testIf() {
	testIf1()
	testIf2()
	testIf3()
	testIf4()
	testIf5()
	testIf6()
	testIf7()
	testIf8()
	testIf9()

}

/*
Богач или бедняк?
На вход подается целое число, сумма денег, которые у вас есть. Ваша задача - вывести марку телефона, которую вы можете себе позволить купить.

Если сумма больше 1000 - вывести Apple
Если сумма от 500 до 1000 (включительно) - вывести Samsung
Если сумма меньше 500 - вывести Nokia с фонариком
*/
func testIf1() {
	var a int
	fmt.Scan(&a)
	if a > 1000 {
		fmt.Println("Apple")
	} else if a >= 500 && a <= 1000 {
		fmt.Println("Samsung")
	} else {
		fmt.Println("Nokia с фонариком")
	}
}

/*
Давайте исправим еще одну программу. Нам нужно, чтобы получился правильный оператор if...else if...else, который проверяет переменную height и если она больше 185,
должен выводиться текст Высокий рост, если она в пределах от 170 до 185 (включительно), должен выводить текст Средний рост,
а если ее значение меньше 170 - текст Небольшой рост
*/
func testIf2() {
	var height int
	fmt.Scan(&height)
	if height > 185 {
		fmt.Println("Высокий рост")
	} else if height >= 170 && height <= 185 {
		fmt.Println("Средний рост")
	} else {
		fmt.Println("Небольшой рост")
	}
}

/*
Вы создаете робота, который может говорить числа.
Ваш робот должен принимать на вход 3 целых числа (каждое с новой строки) в диапазоне от 0 до 10 (включительно) и выдавать соответствующий текст на русском языке.
*/
func testIf3() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	switch a {
	case 0:
		fmt.Println("Ноль")
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
	case 3:
		fmt.Println("Три")
	case 4:
		fmt.Println("Четыре")
	case 5:
		fmt.Println("Пять")
	case 6:
		fmt.Println("Шесть")
	case 7:
		fmt.Println("Семь")
	case 8:
		fmt.Println("Восемь")
	case 9:
		fmt.Println("Девять")
	case 10:
		fmt.Println("Десять")
	default:
		fmt.Println(a)
	}
	switch b {
	case 0:
		fmt.Println("Ноль")
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
	case 3:
		fmt.Println("Три")
	case 4:
		fmt.Println("Четыре")
	case 5:
		fmt.Println("Пять")
	case 6:
		fmt.Println("Шесть")
	case 7:
		fmt.Println("Семь")
	case 8:
		fmt.Println("Восемь")
	case 9:
		fmt.Println("Девять")
	case 10:
		fmt.Println("Десять")
	default:
		fmt.Println(b)
	}
	switch c {
	case 0:
		fmt.Println("Ноль")
	case 1:
		fmt.Println("Один")
	case 2:
		fmt.Println("Два")
	case 3:
		fmt.Println("Три")
	case 4:
		fmt.Println("Четыре")
	case 5:
		fmt.Println("Пять")
	case 6:
		fmt.Println("Шесть")
	case 7:
		fmt.Println("Семь")
	case 8:
		fmt.Println("Восемь")
	case 9:
		fmt.Println("Девять")
	case 10:
		fmt.Println("Десять")
	default:
		fmt.Println(c)
	}
	// Тут у меня возник вопрос. А что будет делать программа когда у нас введут больше 10? Как нам этот момент учитывать?
	//Или раз в требованиях не было, значит этот момент мы не учитываем) Степик у меня принял такое решение
}

/*
Вывести максимальное из двух чисел. Если числа равны, вывести любое из них.

Формат входных данных
Даны два целых числа, каждое из которых не превышает 100000.

Формат выходных данных
Выведите максимальное из них.
*/
func testIf4() {
	var a, b int
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println(a)
	} else if a < b {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}

/*
При регистрации на сайтах требуется вводить пароль дважды. Это сделано для безопасности,
поскольку такой подход уменьшает возможность неверного ввода пароля.
Напишите программу, которая сравнивает пароль и его подтверждение. Если они совпадают,
то программа выводит: «Пароль принят», иначе: «Пароль не принят».

Формат входных данных
На вход программе подаются две строки.

Формат выходных данных
Программа должна вывести одну строку в соответствии с условием задачи.
*/
func testIf5() {
	var a, b string
	fmt.Scan(&a, &b)
	if b == a {
		fmt.Println("Пароль принят")
	} else if b != a {
		fmt.Println("Пароль не принят")
	}
}

/*
Напишите программу, которая определяет, разрешен пользователю доступ к курсам школы или нет.

Формат входных данных
На вход программе подаётся целое число — возраст пользователя.

Формат выходных данных
Программа должна вывести текст «Доступ разрешен» если возраст не менее 12 лет, и «Доступ запрещен» в противном случае.
*/
func testIf6() {
	var a int
	fmt.Scan(&a)
	if a >= 12 {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}
}

/*
Дается одно целое число. Определите, является ли число четным.

Формат входных данных
На вход подается одно целое число, по модулю не превышающее 10^7.

Формат выходных данных
Выведите "YES", если число четное, в противном случае — "NO".
*/
func testIf7() {
	var a int
	fmt.Scan(&a)
	b := a % 2
	if b == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*
Проверить делится ли 𝑎 на 𝑏 нацело.

Формат входных данных
На вход подаются два целых числа
𝑎 и 𝑏, каждое из которых по модулю не превышает
10^6. Гарантируется что b не равна нулю.

Формат выходных данных
Выведите "YES", если
𝑎 делится на 𝑏 нацело, в противном случае — "NO".
*/
func testIf8() {
	var a, b int
	fmt.Scan(&a, &b)
	c := a % b
	if c == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

/*
Говоря простым языком:
Если вводимое число меньше нуля, то выводим
−1;
Если число равно нулю, то выводим 0;
Если число больше нуля, то выводим 1.
*/
func testIf9() {
	var a int
	fmt.Scan(&a)
	if a > 0 {
		fmt.Println("1")
	} else if a == 0 {
		fmt.Println("0")
	} else if a < 0 {
		fmt.Println("-1")
	} else {
		fmt.Println(a)
	}
}
