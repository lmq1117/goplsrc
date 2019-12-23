package main

import "fmt"

func main() {
	fmt.Println(f(), f(), f(), f())
}

func f() *int {
	v := 1
	return &v
}
