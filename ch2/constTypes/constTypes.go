package main

import "fmt"

func main() {
	const x = 10

	var y = x
	var z float64 = x
	var d byte = x

	fmt.Println(x, y, z, d)

	const typedX int = 10

	y = typedX
	z = typedX
	d = typedX

	fmt.Println(x, y, z, d)
}
