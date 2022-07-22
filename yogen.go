package main

import (
	"fmt"
	"math"
)

/*
Point Calculation
target -> x,y

> outer_circle -> 0 points
middle_circle < target < outer_circle -> 1
inner-cirle < middle_circle -> 5
0,0< target<inner_circle -> 10

*/

/*

Formula for Distance from Centre Calculation

Formula: √(x2−x1)^2+(y2−y1)^2
https://www.mathway.com/popular-problems/Algebra/960674

(x-center_x)^2 - (y-center_y)^2 = radius^2


*/

func main() {
	fmt.Println("Hello, World!")
	var distance = CalculateDistanceFromCentre(5, 10)
	fmt.Printf("Calculate : %v\n", distance)
}

func CalculateDistanceFromCentre(x, y float64) int {
	fmt.Println("5 points")
	distance := int(math.Sqrt((math.Pow(x, 2) + math.Pow(y, 2))))

	switch {
	case distance > 5 && distance < 10:
		return 1
	case distance > 1 && distance <= 5:
		return 5
	case distance > 1:
		return 10
	default:
		return 0
	}

}
