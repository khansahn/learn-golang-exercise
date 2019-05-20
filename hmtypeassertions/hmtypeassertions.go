package main

import (
	"fmt"
	"time"
	"io"
	"strings"
	"image"
)

// type-switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v \n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long \n", v, len(v))
	default:
		fmt.Printf("Idk, man %T \n", v)
	}
}

// stringers
type Person struct {
	Name string
	Age int
}
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

// errors
type MyError struct {
	When time.Time
	What string
}
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}
func run() error {
	return &MyError {
		time.Now(),
		"it didn't work",
	}
}

// images
/// package image defines the image interface,
/// package image
/// type Image interface {
/// 	ColorModel() color.Model
///	Bounds() Rectangle
///	At(x, y int) color.Color
///}

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) 	// panic
	//fmt.Println(f)

	// type switches
	do(21)
	do("hello")
	do(true)

	// stringers
	a := Person{"Archie", 18}
	z := Person{"Zathura", 1901}
	fmt.Println(a, z)

	// error
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// readers
	/// Read method from interface io.Reader is:
	/// func (T) Read(b []byte) (n int, err error)
	r := strings.NewReader("Hello, R!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v \n", n, err, b)
		fmt.Printf("b[:n] = %q \n", b[:n])
		if err == io.EOF {
			break
		}
	}

	// images
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
