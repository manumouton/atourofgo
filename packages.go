package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Each time you run this random value, it will return the same number (deterministic environment)
	fmt.Println("My favorite number is", rand.Intn(10))
}
