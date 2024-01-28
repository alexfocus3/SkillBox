//  Пользователь вводит два ряда чисел. Необходимо собрать их в одной структуре, удалив дубликаты.
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

func main() {

	var res = map[string]int{}
	var rslice []int

	reader := bufio.NewReader(os.Stdin)
	str1 := read(reader)
	str2 := read(reader)
	str3 := str1 + str2

	for _, s := range str3 {
		num, _ := strconv.Atoi(string(s))
		if _, ok := res[string(s)]; !ok {
			res[string(s)] = num
			rslice = append(rslice, num)
		}
	}
	fmt.Print(rslice)

}
