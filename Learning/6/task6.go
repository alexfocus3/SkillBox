//Напишите функцию, которая на вход получает число и возвращает true, если число четное, и false, если нечётное.
//Рекомендация Программа запрашивает у пользователя или генерирует случайное число, передает в функцию в качестве аргумента и выводит в консоль результат её работы.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func check(x int) bool {
	if x%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())
	fmt.Println(check(rand.Intn(100)))
}
