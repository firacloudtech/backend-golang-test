package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {

	fmt.Printf("Good morning %v \n", name)

}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)

	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius

}

func main() {
	// sayGreeting("mario")
	// sayGreeting("luigi")
	// sayBye("mario")

	// cycleNames([]string{"mario", "luigi", "barret"}, sayGreeting)
	// cycleNames([]string{"mario", "luigi", "barret"}, sayBye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Println("Area for a1", a1)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)
}
