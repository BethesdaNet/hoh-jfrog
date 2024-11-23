package main

import (
	"fmt"
	"github.com/psychics-of-alderaan/ktoad/pkg/things"
)

func main() {
	fmt.Println("Hello, World!")

	// Create a new Ktoad
	k := things.Ktoad{
		Name:     "Ktoad",
		Greeting: "Hello, World!"}

	fmt.Printf("%s says %q\n", k.Name, k.Greeting)

}
