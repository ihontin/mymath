package greet

import "fmt"

const Testik = "jopa"

type Octopus struct {
	Name string
	Color string
}

func Hello() {
	fmt.Println("Hello, World!")
	
}

func (o Octopus) String() string {
	return fmt.Sprintf("The octopus's name is %q and the color is %s\n", o.Name, o.Color) 
}
