package main

import "fmt"

func main() {

	// var ages = [3]int{20, 5, 20}
	// names := [3]string{"yogen", "audrey", "kyle"}

	// fmt.Println(ages)
	// fmt.Println(names, len(names))

	// slices
	var scores = []int{100, 50, 6}

	fmt.Println(scores, len(scores))

	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	rangeOne := scores[0:3]

	var newArray = scores[len(scores)-2:]

	fmt.Println(rangeOne)
	fmt.Println(newArray)

}
