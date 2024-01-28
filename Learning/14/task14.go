//Напишите функцию, которая на вход принимает функцию вида A func (int, int) int, а внутри оборачивает и вызывает её при выходе (через defer).
//Вызовите эту функцию с тремя разными анонимными функциями A. Тела функций могут быть любыми, но главное, чтобы все три выполняли разное действие.
//Рекомендация В качестве среды разработки может помочь Goland или VScode.

package main

import (
	"fmt"
)

func f(n1, n2 int, A func(int, int) int) int {
	var result int

	defer func() {
		result = A(n1, n2)
		fmt.Println("Результат=", result)
	}()

	return result
}
func main() {

	fmt.Println("Результат defer =", f(10, 20, func(n1, n2 int) int { return n1 * n2 }))
	fmt.Println("Результат defer =", f(10, 20, func(n1, n2 int) int { return n1 + n2 }))
	fmt.Println("Результат defer =", f(10, 20, func(n1, n2 int) int { return n1 - n2 }))
}
