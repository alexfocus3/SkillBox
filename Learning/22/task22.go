//Напишите программу, демонстрирующую работу каналов.
Создайте буферизованный канал, заполните его числами, выведите его длину и емкость, после чего выведите его значения.

package main

import "fmt"

func main() {

	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 7
	intCh <- 3
	fmt.Println(cap(intCh))
	fmt.Println(len(intCh))

	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
	fmt.Println(<-intCh)
}
