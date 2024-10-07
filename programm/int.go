package main

import "fmt"

func testInt() {
	testInt1()
	testInt2()
	testInt3()
	testInt4()
	testInt5()
	testInt6()
	testInt7()
	testInt8()
	testInt9()
	testInt10()
	testInt11()
	testInt12()
	testInt13()
	testInt14()
	testInt15()
	testInt16()
	testInt17()

}

/*
По данному целому числу найдите его квадрат.

Формат входных данных
На вход дается одно целое число, по модулю не превосходящее
100.

Формат выходных данных
Программа должна вывести квадрат данного числа.
*/
func testInt1() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a * a)
}

/*
Расстояние в километрах
Напишите программу, которая находит полное число километров по заданному числу метров.

Формат входных данных
На вход программе подаётся натуральное число – количество метров.

Формат выходных данных
Программа должна вывести одно число – полное число километров.
*/

func testInt2() {
	var b int
	fmt.Scan(&b)
	fmt.Println(b / 1000)
}

/*
Дано число 𝑥. Вычислите число 𝑥^6 при помощи трех операций умножения.

Формат входных данных
Программа получает на вход одно число 𝑥

Формат выходных данных
Программа должна вывести 𝑥^6 при помощи трех операций.
*/
func testInt3() {
	var x int
	fmt.Scan(&x)
	c := x * x * x
	d := c * c
	fmt.Println(d)
	// put your code here
}

/*
Напишите программу, которая считывает три целых числа и выводит на экран их произведение. Каждое число записано в отдельной строке.

Формат входных данных
На вход программе подаётся три целых числа, каждое на отдельной строке.

Формат выходных данных
Программа должна вывести произведение введенных чисел.
*/
func testInt4() {
	var a int
	var b int
	var c int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	x := a * b * c
	fmt.Println(x)
}

/*
𝑁 школьников делят 𝐾 яблок поровну, неделящийся остаток остается в корзинке. Сколько яблок достанется каждому школьнику?

Формат входных данных
На вход дается два целых положительных числа 𝑁 и 𝐾, каждое из которых не превышает 10000.

Формат выходных данных
Программа должна вывести количество яблок, которое достанется каждому школьнику.
*/
func testInt5() {
	var N, K int
	fmt.Scan(&N, &K)
	z := K / N
	fmt.Println(z)
}

/*
𝑁 школьников делят 𝐾 яблок поровну, неделящийся остаток остается в корзинке. Сколько яблок останется в корзинке?
 Формат входных данных
 На вход дается два целых положительных числа 𝑁 и 𝐾, каждое из которых не превышает 10000.

 Формат выходных данных
 Программа должна вывести искомое количество яблок.
*/

func testInt6() {
	var N, K uint16
	fmt.Scan(&N, &K)
	y := K % N
	fmt.Println(y)
}

/*
Напишите программу вывода на экран трех последовательно идущих чисел, каждое на отдельной строке.
Первое число вводит пользователь, остальные числа вычисляются в программе.

Формат входных данных
На вход программе подается одно целое число.

Формат выходных данных
Программа должна вывести три последовательно идущих числа в соответствии с условием задачи.
*/

func testInt7() {
	var a int
	fmt.Scan(&a)
	b := a
	b++
	c := b
	c++
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

/*
Напишите программу, которая считает стоимость трех телефонов, состоящих из самого устройства, чехла, зарядки и наушников.

Формат входных данных
На вход программе подаётся четыре целых числа, каждое на отдельной строке.
В первой строке — стоимость самого устройства, во второй строке — стоимость чехла,
в третьей строке — стоимость зарядки и в четвертой строке — стоимость наушников.

Формат выходных данных
Программа должна вывести одно число – стоимость покупки (трех телефонов).
*/
func testInt8() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	S := 3 * (a + b + c + d)
	fmt.Println(S)
}

/*
Пирожок в столовой стоит 𝑎 рублей и 𝑏 копеек. Определите, сколько рублей и копеек нужно заплатить за
𝑛 пирожков.

Формат входных данных
Программа получает на вход три целых числа:
𝑎,𝑏,𝑛.

Формат выходных данных
Программа должна вывести два числа: стоимость покупки в рублях и копейках.
*/
func testInt9() {
	var a, b, n int
	fmt.Scan(&a, &b, &n)
	rvk := a * 100 // рубли в копейки
	smk := b + rvk // сумма всех копеек
	sp := smk * n  // стоимость пирожков
	r := sp / 100
	k := sp % 100
	fmt.Println(r, k)
}

/*
Напишите программу для пересчёта величины временного интервала, заданного в минутах, в величину, выраженную в часах и минутах.

Формат входных данных
На вход программе подаётся целое число – количество минут.

Формат выходных данных
Программа должна вывести текст в соответствии с условием задачи.
*/
func testInt10() {
	var a int
	fmt.Scan(&a)
	h := a / 60
	m := a % 60
	fmt.Println(a, "мин - это", h, "час", m, "минут.")
}

/*
Напишите программу, которая считывает целое число, после чего на экран выводится следующее и предыдущее целое число с пояснительным текстом.

Формат входных данных
На вход программе подаётся целое число.

Формат выходных данных
Программа должна вывести текст согласно условию задачи.
*/
func testInt11() {
	var a int
	fmt.Scan(&a)
	b := a
	b++
	c := a
	c--
	fmt.Println("Следующее за числом", a, "число:", b)
	fmt.Println("Для числа", a, "предыдущее число:", c)
}

/*
По данному целому числу, найдите его квадрат.
*/
func testInt12() {
	var a int
	fmt.Scan(&a)
	a = a * a
	fmt.Println(a)
}

/*
Дано натуральное число, выведите его последнюю цифру.

Формат входных данных
На вход дается натуральное число N, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - ответ на задачу.
*/
func testInt13() {
	var a int
	fmt.Scan(&a)
	a = a % 10
	fmt.Println(a)
}

/*
Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).

Формат входных данных
На вход дается натуральное число, не превосходящее 10000.

Формат выходных данных
Выведите одно целое число - число десятков.
*/
func testInt14() {
	var a int
	fmt.Scan(&a)
	a = (a % 100) / 10
	fmt.Println(a)
}

/*
Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

Входные данные
На вход программе подается целое число d (0 < d < 360).

Выходные данные
Выведите на экран фразу:
It is ... hours ... minutes.
Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.X
*/
func testInt15() {
	var d int
	fmt.Scan(&d)
	h := d / 30
	m := 2 * (d % 30)
	fmt.Println("It is", h, "hours", m, "minutes.")
}

/*
Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число — сумму цифр введенного числа.
*/
func testInt16() {
	var a int
	fmt.Scan(&a)
	sum := a/100 + (a%100)/10 + a%10
	fmt.Println(sum)
}

/*
Дано трехзначное число. Переверните число и выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.
*/
func testInt17() {
	var a int
	fmt.Scan(&a)
	b := a % 10
	c := (a % 100) / 10
	d := a / 100
	fmt.Print(b)
	fmt.Print(c)
	fmt.Print(d)
}
