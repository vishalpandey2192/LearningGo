package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (y truck) transportationDevice() string {
	return fmt.Sprintln("I am a truck with", y.doors, "doors and I am durable.")
}

func (z sedan) transportationDevice() string {
	return fmt.Sprintln("I am a sedan with", z.doors, "doors and I am luxurious.")
}

func main() {
	t1 := truck{
		vehicle{
			2,
			"red",
		},
		true,
	}
	fmt.Println(t1.transportationDevice())

	s1 := sedan{
		vehicle{
			4,
			"white",
		},
		false,
	}
	fmt.Println(s1.transportationDevice())
}
