package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y  * v.Y)
}

// another way to write it as function, but basically methods are function with a receiver argument
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// methods with a non-struct types
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f<0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v:= Vertex{3,4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f:= MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
