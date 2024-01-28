//  Пользователь вводит два ряда чисел. Необходимо собрать их в одном срезе, удалив дубликаты.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func read(reader *bufio.Reader) string {

	fmt.Print("Введите строку:")
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	return string(line)

}

func check(str []int, n int) bool {

	for _, i := range str {
		if i == n {
			return true
		}
	}
	return false
}

func main() {

	var res []int

	reader := bufio.NewReader(os.Stdin)
	str1 := read(reader)
	str2 := read(reader)
	str3 := str1 + str2

	for _, s := range str3 {
		num, _ := strconv.Atoi(string(s))
		if check(res, num) == false {
			res = append(res, num)
		}
	}
	fmt.Println(res)

}
