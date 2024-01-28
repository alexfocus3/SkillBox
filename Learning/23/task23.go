//Напишите программу, демонстрирующую HTTP запрос  к любому хосту в сети.

package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	httpRequest := "GET / HTTP/1.1\n" +
		"Host: waypay.ru\n\n"
	conn, err := net.Dial("tcp", "waypay.ru:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}
