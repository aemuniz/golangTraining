package main

import (
	"fmt"
)

var x int = 42

func main() {

	a := 20
	b := 300
	c := "Arthur"
	d := "Jessyca"
	e := 3.14
	f := 2000.23

	fmt.Println("Hello World!!")
	fmt.Printf("%d - %b \n", 42, 42)
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	fmt.Printf("%T - %T - %T - %T - %T - %T \n", a, b, c, d, e, f)
	fmt.Printf("%v - %v - %v - %v - %v - %v \n", a, b, c, d, e, f)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}

	foo("Arthur")
	y := 17
	fmt.Printf("I'm %v", y, " years old!")
	x := 42
}


func foo(str string){
	fmt.Printf("%T", str)
	fmt.Println("%b")
	fmt.Printf("%v", x) // out of scope;;
}



