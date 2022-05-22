package main

import "fmt"

/* Example 2-1

func main() {
	fmt.Println("Hello, world!")
} */

/* Example 2-3
const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
} */

/* go vet & golangci-lint run
func main() {
	x := 10
	x = 20
	fmt.Println(x)
	x = 30
} */

// latin a and alphabet a is a different
func main() {
	ａ := "hello"   // Unicode U+FF41
	a := "goodbye" // standard lowercase a (Unicode U+0061)
	fmt.Println(ａ)
	fmt.Println(a)
}
