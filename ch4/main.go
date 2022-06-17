package main

import (
	"fmt"
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
/* func main() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
} */

// ex4-8 for statement
func main() {
	fmt.Println("=== complate for statement ===")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("=== condition-only for statement ===")
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	fmt.Println("=== Infinite looping nostalgia for statement ===")
	count := 1
	for {
		fmt.Println("Hello")
		if count > 10 {
			break
		}
		count++
	}

	fmt.Println("=== The for-range loop for statement ===")
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	for _, v := range evenVals {
		fmt.Println(v)
	}

	uniqueNames := map[string]bool{
		"Fred":  true,
		"Raul":  true,
		"Wilma": true,
	}
	for k := range uniqueNames {
		fmt.Println(k)
	}
}
