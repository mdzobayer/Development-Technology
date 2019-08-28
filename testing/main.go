package main

import "fmt"

func main() {
	greetMessageJhon := hello("John")

	fmt.Println(greetMessageJhon)

	greetMessageEmpty := hello("")

	fmt.Println(greetMessageEmpty)
}
