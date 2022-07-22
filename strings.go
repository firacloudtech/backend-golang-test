package main

import (
	"fmt"
	"sort"
)

func main() {

	// greeting := "hello there friends!"

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, "there"))

	ages := []int{45, 20, 30, 75, 60, 50, 25}

	sort.Ints(ages)

	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

}
