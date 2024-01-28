//Напишите программу, которая читает и выводит в консоль строки из файла,
//созданного в предыдущей практике, без использования ioutil.
//Если файл отсутствует или пуст, выведите в консоль соответствующее сообщение.
//Рекомендация Для получения размера файла воспользуйтесь методом Stat(), который возвращает информацию о файле и ошибку.

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("res.txt")
	if err != nil {
		fmt.Println("Unable to open file:", err)
		return
	}
	defer file.Close()

	fileDataStruct, err := os.Stat("res.txt")
	if err != nil {
		fmt.Println(err)
	} else if fileDataStruct.Size() == 0 {
		fmt.Println("File is empty")
		return
	}

	buf := make([]byte, fileDataStruct.Size())
	file.Read(buf)
	fmt.Print(string(buf))
}
