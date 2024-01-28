//Напишите функцию, которая принимает в качестве аргументов указатели на два типа int и меняет их значения местами.
//Рекомендация В методе main создайте и присвойте значения двум переменным типа int, выведите значения этих переменных в консоль до вызова функции и после

package main

import (
	"fmt"
)

func changeValue(x, y *int) {
	*x, *y = *y, *x
}

func main() {

	a := 10
	b := 20
	fmt.Println("a before =", a)
	fmt.Println("b before =", b)
	changeValue(&a, &b)
	fmt.Println("a after =", a)
	fmt.Println("b after =", b)

}
