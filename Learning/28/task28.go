//Напишите простой пример использования горутин и передачи информации между ними чс помощью каналов.
package main

import (
	"fmt"
	"time"
)

func pinger(c chan int) {
	for i := 1; ; i++ {
		c <- i
	}
}
func printer(c chan int) {
	for {
		fmt.Println("прошло", <-c, "секунд")
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan int = make(chan int)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
