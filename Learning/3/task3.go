//Напишите программу, создающую текстовый файл только для чтения, и проверьте, что в него нельзя записать данные.
//Рекомендация. Для проверки создайте файл, установите режим только для чтения, закройте его, а затем, открыв, попытайтесь прочесть из него данные.

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("res.txt")
	if err != nil {
		fmt.Println("Unable to create file: ", err)
		os.Exit(1)
	}
	fmt.Println("Создан файл: ", file.Name())
	os.Chmod("res.txt", 0444)
	file.Close()
	file, err = os.OpenFile("res.txt", os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Unable to open file: ", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString("Новая текстовая строка")
	if err != nil {
		fmt.Println("Unable to add string in the file: ", err)
		os.Exit(1)
	}

}
