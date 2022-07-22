package main

import "fmt"

func main() {

	names := []string{"yogen", "audrey", "kathir"}

	for _, value := range names {
		fmt.Printf("Hello %v \n", value)
	}

}
