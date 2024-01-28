// Найти самое длинное слово в предложении без помощи len()
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print("Введите строку:")
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}

	trim := " "
	slice := strings.Split(string(line), trim)

	maxword := ""
	maxlen := 0

	for _, i := range slice {
		count := 0
		for range i {
			count++
		}
		if count > maxlen {
			maxlen = count
			maxword = i
		}
	}
	fmt.Print(maxword)
}
