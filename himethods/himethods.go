package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// methods
func (v Vertex) Abs() float64 {  // method dengan receiver Vertex, tanpa input, output float64
	return math.Sqrt(v.X * v.X + v.Y  * v.Y)
}

// another way to write it as function, but basically methods are function with a receiver argument
func Abs(v Vertex) float64 { 	// method dengan input Vertex, output float64
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// methods with a non-struct types
type MyFloat float64
func (f MyFloat) Abs() float64 {  // method dengan receiver MyFloat, tanpa input, output float64
	if f<0 {
		return float64(-f)
	}
	return float64(f)
}

// methods with pointer-receivers
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// personal-note
/*
cara format nulis function di golang itu,
func //namafunction// (/tipe data yang masuk ke function, OPT/) /tipe data keluaran function, kalau ga nge-return apa2 berarti ga usah diisi, OPT/ {

}

sedangkan kalau format nulis method itu,
func (/type data atau struct yang bakal di-associate sama method ini/) //namamethod// (/tipe data yang masuk ke method, OPT/) /tipe data keluaran method, OPT/ {

}

*/

func main() {
	// methods
	v:= Vertex{3,4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
	
	// methods with non-struct type
	f:= MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// methods with pointer receivers
	vpr := Vertex{3,4}
	vpr.Scale(10)
	fmt.Println(vpr)
	fmt.Println(vpr.Abs())
}
