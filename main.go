package main

import "fmt"

func main() {
	var i int = 1
	fmt.Println(&i, i)
	incr(&i)
	fmt.Println(&i, i)
	incr(&i)
	fmt.Println(&i, i)
	incr(&i)
	fmt.Println(&i, i)

}

func incr(p *int) {
	*p++
}
