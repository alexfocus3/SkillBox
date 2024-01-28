//Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения заданного числа в массив.
//Сложность алгоритма должна быть минимальная.
//Что оценивается Верность индекса.
//При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.

package main

import (
	"fmt"
)

func main() {

	const size = 12
	var array [size]int = [size]int{1, 2, 2, 2, 2, 6, 7, 9, 9, 9, 10, 11}
	var zn, n, pos, start, end int

	fmt.Print("Введите число:")
	fmt.Scan(&n)

	start = 0
	end = size - 1
	for {
		zn = int((end-start)/2) + start
		if array[zn] == n {
			pos = zn
			break
		} else if array[zn] > n {
			end = zn - 1
		} else {
			start = zn + 1
		}
		if array[end] < n {
			fmt.Println("результат не найден")
			return
		} else if array[start] > n {
			fmt.Println("результат не найден")
			return
		}
	}

	i := pos - 1
	for i > 0 {
		if array[i] == n {
			pos--
			i--
		} else {
			break
		}
	}
	fmt.Println("результат=", pos)

}
