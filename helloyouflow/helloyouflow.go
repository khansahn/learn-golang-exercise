package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g \n", v, lim)
	}
	return lim
}

func main() {
	sum := 0
	for i:= 0; i< 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// for-loop
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// while- alike
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// rekursif
	fmt.Println(sqrt(2), sqrt(-4))

	// if- with a short statement 
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// switch
	fmt.Print("Go runs on ")
	switch os:= runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os)
	}

	  defer fmt.Println("yaaaaaaaaaaaaaaaaaaaaaa!!!")


	// switch lagi
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Wednesday {
	case today +0:
		fmt.Println("Today.")
	case today +1:
		fmt.Println("Tomorrow.")
	case today +2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Idk, man.")
	}

	// defer
	fmt.Println("start")
	for i:=0; i<10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
