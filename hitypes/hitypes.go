package main

import (
	"fmt"
	"strings"
	"math"
)

// struct : is a collection of fields
type Vertex struct {
	X int
	Y int
}

// struct literals
var (
        v1 = Vertex{1,2}        
        v2 = Vertex{X:1}        // Y:0 is implicit
        v3 = Vertex{}           // x:0 and Y:0 implicit
        pv1 = &Vertex{1,2}

)

// map
type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

// map-literals
var mliteral = map[string]Vertex2{
	"K Labs": Vertex2{
		40.2313, -78.2131,
	},
	"Google": Vertex2{
		231.1312, -0.3123123,
	},
}

// slice length and capacity
func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v \n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v \n", s, len(x), cap(x), x)
}

// function-values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3,4)
}

func main() {

	// pointers
	i, j := 42, 2701

	p := &i 	// point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21 	// set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j 		// point to j
	*p = *p / 37 	// divide j through the pointer
	fmt.Println(j)	// see the new value of j

	// struct
	fmt.Println(Vertex{1,23})

	// accessing struct fields
	v := Vertex{1,2}
	v.X = 4
	fmt.Println(v)

	// pointers to stucts
	pv := &v
	pv.X = 1e9
	fmt.Println(v)

	// struct literals
	fmt.Println(v1, pv1, v2, v3)

	// array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0])
	fmt.Println(a)

	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)

	// slices : are like references to arrays jadi beda sama operand di python (?? ya kayaknya hhe)
	fmt.Println(primes[1:4])

	beatles := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(beatles)
	beatlesa := beatles[0:2]
	beatlesb := beatles[1:3]
	fmt.Println(beatlesa, beatlesb)
	beatlesb[0] = "XXX"
	fmt.Println(beatlesa, beatlesb)
	fmt.Println(beatles)

	// slice literal : like an array literal without the length
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	fmt.Println(q[1:4])
	fmt.Println(q[:2])
	fmt.Println(q[1:])

	// slices length and capacity
	// slice length : the number of elements it contains
	// slice capacity : the number of elements IN THE UNDERLYING ARRAY counting from the FIRST ELEMENT IN THE SLICE

	/// slice to give it zero length
	q = q[:0]
	printSlice(q)
	/// extend its length
	q = q[:4]
	printSlice(q)
	/// drop it first to values
	q = q[2:]
	printSlice(q)

	// none nya tuh nil hahahah
	var k []int
	fmt.Println(k, len(k), cap(k))
	if k == nil {
		fmt.Println("nil si k")
	}

	// creating slice with make, make function allocates a zeroed array and returns a slice that refers to that array
	aslice := make([]int, 5)
	printSlice2("a", aslice)
	bslice := make([]int, 5)
	printSlice2("b", bslice)
	cslice := bslice[:2]
	printSlice2("c", cslice)
	dslice := cslice[2:5]
	printSlice2("d", dslice)


	// slice of slice
	/// create a tic-tac toe board
	board := [][]string {
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	/// the players take turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i:=0; i<len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}

	// append
	var sappend []int
	printSlice(sappend)

	sappend = append(sappend, 0)
	sappend = append(sappend, 1, 2, 3)
	printSlice(sappend)

	// range
	var pow = []int{1,2,4,8,16,32,64,128}
	for i:= range pow {
		fmt.Println(i)
	}
	for i, v:= range pow {
		fmt.Println(v,i)
	}

	pow2 := make([]int, 10)
	for  i:= range pow2 {
		pow2[i] = 1 << uint(i)  	// == 2**i
	}
	for _, value:= range pow2 {
		fmt.Printf("%d\n", value)
	}

	// maps
	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.4392, -79.29931,
	}
	fmt.Println(m["Bell Labs"])

	// map-literals
	fmt.Println(mliteral)

	// mutating-maps
	mmutating := make(map[string]int)
	mmutating["An"] = 42
	fmt.Println("The value:", mmutating["An"])
	mmutating["An"] = 48
	fmt.Println("The value:", mmutating["An"])
	vmm, ok := mmutating["An"]
	fmt.Println("The value:", vmm, "Present?", ok)
	delete(mmutating, "An")
	fmt.Println("The value:", mmutating["An"])
	vmm, ok = mmutating["An"]
	fmt.Println("The value:", vmm, "Present?", ok)

	// function-values
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))

	fmt.Println(compute(hypot))
}
