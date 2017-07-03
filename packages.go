package main

import (
	"fmt"
	"math/rand"
	"math"
)

//A function can take several arguments
//The variable type is coming after the var identifier
//func add(x int, y int) int {
//could also be writter
func add(x, y int) int {
	return x + y
}

//a function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

//named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, java, scala bool
var i, j int = 1, 2

func main() {
	//Each time you run this random value, it will return the same number (deterministic environment)
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Now you have %g problems.", math.Sqrt(7))
	//In Go, an identifier is exported if it starts with an uppercase letter, e.g. pi is not exported, Pi or PI is
	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Println(i, j, c, scala, java)


}
