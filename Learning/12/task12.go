//Напишите функцию, которая производит слияние двух отсортированных массивов длиной четыре и пять в один массив длиной девять.
//Советы и рекомендации Обратите внимание на размеры массивов. В качестве среды разработки может помочь GoLand или VS Code.
//Что оценивается Правильность размеров. Правильный порядок элементов в конечном массиве.

package main

import (
	"fmt"
)

func main() {

	const size1 = 4
	const size2 = 5
	const size3 = size1 + size2

	var arr1 [size1]int = [size1]int{1, 3, 5, 7}
	var arr2 [size2]int = [size2]int{2, 4, 6, 8, 9}
	var arr3 [size3]int

	for i := 0; i < size1; i++ {
		s := i * 2
		if arr1[i] < arr2[i] {
			arr3[s] = arr1[i]
			arr3[s+1] = arr2[i]
		} else {
			arr3[s] = arr3[i]
			arr3[s+1] = arr1[i]
		}
	}

	k := size1
	for i := (size1 * 2); i < size3; i++ {
		arr3[i] = arr2[k]
		k++
	}

	fmt.Println(arr3)
}
