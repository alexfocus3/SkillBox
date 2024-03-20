//Напишите программу, которая на вход получала бы строку, введённую пользователем,
//а в файл писала № строки, дату и сообщение в формате: 2020-02-10 15:00:00 продам гараж.
//При вводе слова exit программа завершает работу.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var str string
	var count int = 1

	file, err := os.Create("res.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	for {
		fmt.Print("Введите строку: ")
		fmt.Fscan(os.Stdin, &str)
		if str == "exit" {
			break
		} else {
			fmt.Fprintf(file, "%d %s %s\n", count, time.Now().Format("01-02-2006 15:04:05"), "продам гараж!!")
		}
		count++
	}
}
