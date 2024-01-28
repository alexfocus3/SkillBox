// Напишите программу, реверсирующую цифры заданного числа.

package main

import (
	"fmt"
	//"unicode/utf8"
)

func main() {

	var x int = 1357
	for x > 0 {
		fmt.Print(x % 10)
		x /= 10
	}

}
