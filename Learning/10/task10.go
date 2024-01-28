//Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество чётных и нечётных чисел. Для ввода и подсчёта используйте разные циклы.
//Что оценивается Для введённых чисел 1, 1, 1, 2, 2, 2, 3, 3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6.

package main

import (
	"fmt"
)

func main() {

	var numbers [10]int
	var even, odd int
	for i, _ := range numbers {
		fmt.Print("Введите число: ")
		fmt.Scan(&numbers[i])
	}
	for i, _ := range numbers {
		if numbers[i]%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println("чётных —", even, ", нечётных —", odd)
}
