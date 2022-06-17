package main

import (
	"fmt"
	"math/rand"
)

// ex4-1. Shadowing variables
/* func main() {
	x := 10
	if x > 5 {
		fmt.Println(x)
		x := 5
		fmt.Println(x)
	}
	fmt.Println(x)
} */

// ex4-2. Shadowing with multiple assignment
/* func main() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)
} */

// ex4-3. Shadowing package name
/* func main() {
	x := 10
	fmt.Println(x)
	fmt := "oops"
	fmt.Println(fmt)
} */

//ex4-6 Scoping a variable to an if statement
func main() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}
