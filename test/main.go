package main

import "fmt"

var update = func(x *int) {
	*x = 12
}

func main() {
	x := 10
	var array = []int{1,2,3}
	update(&x)
	array = append(array, 9);
	fmt.Println(x, array)
}