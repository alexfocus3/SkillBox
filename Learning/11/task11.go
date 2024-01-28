//Напишите функцию, принимающую на вход массив и возвращающую массив, в котором элементы идут в обратном порядке по сравнению с исходным.
//Напишите программу, демонстрирующую работу этого метода.
//Что оценивается При вводе 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 программа должна выводить при помощи дополнительной функции, реверсировав массив: 10, 9, 8, 7, 6, 5, 4, 3, 2, 1.

package main

import (
	"fmt"
)

func reverse(array [10]int) {

	size := len(array)
	for i := 0; i < size/2; i++ {
		j := size - 1 - i
		array[i], array[j] = array[j], array[i]
	}
	fmt.Println(array)
}

func main() {

	var numbers [10]int
	for i, _ := range numbers {
		fmt.Print("Введите число: ")
		fmt.Scan(&numbers[i])
	}
	reverse(numbers)
}
