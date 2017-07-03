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
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

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

}