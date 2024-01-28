//Программа должна получать на вход имена двух файлов, необходимо конкатенировать их содержимое, используя strings.Join.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {

	file1, err := os.OpenFile("1.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer file1.Close()
	wr1 := bytes.Buffer{}
	sc1 := bufio.NewScanner(file1)
	for sc1.Scan() {
		if wr1.Len() > 0 {
			wr1.WriteString("\n")
		}
		wr1.WriteString(sc1.Text())
	}

	file2, err := os.OpenFile("2.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Unable to open file:", err)
		os.Exit(1)
	}
	defer file2.Close()
	wr2 := bytes.Buffer{}
	sc2 := bufio.NewScanner(file2)
	for sc2.Scan() {
		if wr1.Len() > 0 {
			wr2.WriteString("\n")
		}
		wr2.WriteString(sc2.Text())
	}

	file3, err := os.Create("3.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file3.Close()

	file3.WriteString(wr1.String())
	file3.WriteString(wr2.String())

	fmt.Println("created file ", file3.Name())
}
