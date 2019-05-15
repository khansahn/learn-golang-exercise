package main

import (
	"fmt"
	"math/cmplx"
	"math"
)

func add(x, y int) int {
    return x + y
}

func swap(x, y string) (string, string) {
    return y, x
}

func split(sum int) (x, y int) {
    x = (sum * 4) / 9
    y = sum - x
    return
}

func needInt(x int) int {
    return x*10 + 1
}

func needFloat(x float64) float64 {
    return x * 0.1
}

var c, python, java bool
var i, j int = 1, 2
var (
    ToBe bool = false
    MaxInt  uint64 = 1<<64 - 1
    z complex128 = cmplx.Sqrt(-5 + 12i)
    )

const Pi = 3.14
const (
	Big = 1 << 100
	Small = Big >> 99
      )


func main() {
    fmt.Printf("hello, world\n")
    fmt.Println(add(12,12))

    a, b := swap("ikan", "inka")
    fmt.Println(a,b)

    fmt.Println(split(17))

    var k int
    fmt.Println(k, c, python, java)

    var c, d, e = true, false, "yeeeay!"
    fmt.Println(c, d, e)

    fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
    fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
    fmt.Printf("Type: %T Value: %v \n", z, z)

    var m, n int = 3, 4
    var o float64 = math.Sqrt(float64(m*m +n*n))
    var p uint = uint(o)
    fmt.Println(m, n, o, p)

    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")

    const truth = true
    fmt.Println("Go rules?", truth)

    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
