package main

import "fmt"

type person struct {
	fName   string
	lName   string
	favFood []string
}

func (g person) walk() string {
	return fmt.Sprintln(g.fName, "is walking")
}

func main() {
	p1 := person{
		"Todd",
		"McLeod",
		[]string{"carrots", "oranges", "peaches", "strawberries"},
	}

	s := p1.walk()
	fmt.Println(s)

}
