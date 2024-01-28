//Напишите программу, которая на вход получает число, затем с помощью двух функций преобразует его.
//Первая умножает, а вторая прибавляет число, используя именованные возвращаемые значения.

package main

import (
	"fmt"
	"os"
)

func multiplication(x int) (res int) {
	res = x * x
	return
}

func addition(x int) (res int) {
	res = x + x
	return
}

func main() {

	var x int
	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &x)

	x = multiplication(x)
	x = addition(x)

	fmt.Println(x)
}
