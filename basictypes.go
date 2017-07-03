//bool
//
//string
//
//int  int8  int16  int32  int64
//uint uint8 uint16 uint32 uint64 uintptr
//
//byte => alias for uint8
//
//rune => alias for int32 , représente un "code point" Unicode
//
//float32 float64
//
//complex64 complex128

package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

//func giveMeYourType(v){
//	fmt.Printf("v is of type %T\n", v)
//}

func sumTill(x int){
	sum:=0
	for i:=0; i< x; i++{
		sum+=i
	}
	fmt.Println(sum)
}

func sumTillXWhile(x int) {
	sum:=1
	for sum < x {
		sum += sum
	}

	fmt.Println(sum)
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}


func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	//Zero value for primitive types
	var i int
	var fl float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, fl, b, s)

	sumTill(11)

	sumTillXWhile(110)

	//How to generate an infinite loop in go
	//for{
	//
	//}

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}