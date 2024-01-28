//Напишите программу, которая с помощью функции генерирует три случайные точки в двумерном пространстве (две координаты),
//а затем с помощью другой функции преобразует эти координаты по формулам: x = 2 × x + 10, y = −3 × y − 5.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func transform(x, y int) (z, n int) {

	z = x*2 + 10
	n = -3*y - 5
	return
}

func main() {

	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	y := rand.Intn(100)
	fmt.Println("x1 before =", x)
	fmt.Println("y1 before =", y)
	x, y = transform(x, y)
	fmt.Println("x1 after =", x)
	fmt.Println("y1 after =", y)
	fmt.Println("----------------------------")

	x = rand.Intn(100)
	y = rand.Intn(100)
	fmt.Println("x2 before =", x)
	fmt.Println("y2 before =", y)
	x, y = transform(x, y)
	fmt.Println("x2 after =", x)
	fmt.Println("y2 after =", y)
	fmt.Println("----------------------------")

	x = rand.Intn(100)
	y = rand.Intn(100)
	fmt.Println("x3 before =", x)
	fmt.Println("y3 before =", y)
	x, y = transform(x, y)
	fmt.Println("x3 after =", x)
	fmt.Println("y3 after =", y)
	fmt.Println("----------------------------")

}
