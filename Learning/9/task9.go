//Напишите программу, в которой будет три функции, попарно использующие разные глобальные переменные.
//Функции должны прибавлять к поданному на вход числу глобальную переменную и возвращать результат.
//Затем вызовите по очереди три функции, передавая результат из одной в другую.

package main

import (
	"fmt"
	"os"
)

var first int = 10
var second int = 7
var third int = 15

func _1(x int) int {

	b := x + first
	return b
}

func _2(x int) int {

	b := x + second
	return b
}

func _3(x int) int {

	b := x + third
	return b
}

func main() {

	var x int
	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &x)
	fmt.Println(_3(_2(_1(x))))
}
