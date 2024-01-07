package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("The time is", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))

	fmt.Printf("Now you have %v problems.\n", math.Sqrt(7))

	fmt.Println(math.Pi)

	fmt.Println(add(42, 13))

	var i, j int = 1, 2

	fmt.Println(i, j)

	var c, python, java = true, false, "no!" // type inference
	// without type inference
	// var c bool = true
	// var python bool = false
	// var java string = "no!"

	fmt.Println(c, python, java)

	const Pi float32 = 3.14
	fmt.Println(Pi)

	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

}
