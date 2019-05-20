package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type MyFloat float64
type Vertex struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	if f <0 {
		return float64(-f)
	}
	return float64(f)
}
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// interfaces are implemented implicitly
type I interface {
	M()
}
type T struct {
	S string
}
/// this method means type T implements the interface I, implicitly
func (t T) M() {
	fmt.Println(t.S)
}

// interfaces values
type I2 interface {
	M2()
}
func (t *T) M2() {
	fmt.Println(t.S)
}
type F float64
func (f F) M2() {
	fmt.Println(f)
}
func describe2(i I2) {
	fmt.Printf("(%v, %T) \n", i, i)
}

// interface values with nil underlying values
/// if the concrete value inside the interface itself is nil, the method will be called with a nil receiver
type I3 interface {
	M3()
}
func (t *T) M3() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
func describe3(i I3) {
	fmt.Printf("(%v, %T) \n", i, i)
}

// nil interface values
//type I4 interface {
//	M4()
//}
//func describe4( i I4) {
//	fmt.Printf("(%v, %T) \n", i, i)
//}

// empty interface
func describe5(i interface{}) {
	fmt.Printf("(%v, %T) \n", i, i)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//	a = v

	fmt.Println(a.Abs())

	// interfaces are implemented implicitly
	var i I = T{"hello"}
	i.M()

	// interfaces values
	var i2 I2
	i2 = &T{"haiiii"}
	describe2(i2)
	i2.M2()

	i2 = F(math.Pi)
	describe2(i2)
	i2.M2()

	// interface value with nil underlying values
	var i3 I3
	var t3 *T
	i3 = t3
	describe3(i3)
	i3.M3()

	i3 = &T{"hellll-o"}
	describe3(i3)
	i3.M3()

	// nil interface values
	//var i4 I4
	//describe4(i4)
	//i4.M4()

	// empty interface
	var i5 interface{}
	describe5(i5)

	i5 = 42
	describe5(i5)

	i5 = "hhehh"
	describe5(i5)
}
