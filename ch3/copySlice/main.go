package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)

	y = make([]int, 2)
	num = copy(y, x)
	fmt.Println(y, num)

	y = make([]int, 2)
	num = copy(y, x[2:])
	fmt.Println(y, num)

	num = copy(x[:3], x[1:])
	fmt.Println(x, num)
}
