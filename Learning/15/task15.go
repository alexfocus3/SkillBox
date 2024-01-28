//Заполните массив неупорядоченными числами на основе генератора случайных чисел.
//Введите число. Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого.
//При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.
//Что оценивается Программа должна корректно считать числа в массиве после заданного.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const size = 10
	var array [size]int
	var n int

	rand.Seed(time.Now().UnixNano())
	for i, _ := range array {
		array[i] = rand.Intn(10)
	}

	fmt.Print("Введите целое число: ")
	fmt.Scan(&n)

	res := 0
	for i, _ := range array {
		if array[i] == n {
			res = size - i - 1
			break
		}
	}

	fmt.Println(array)
	fmt.Println(n)
	fmt.Println(res)
}
