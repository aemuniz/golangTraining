package main

import "fmt"

//I don't have to worry about imports // the webstorm ide does it automatically for me!!

func main() {

	fmt.Println("Hello World!!")
	fmt.Printf("%d - %b \n", 42, 42)
	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}

	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}
}




