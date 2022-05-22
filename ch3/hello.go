package main

import "fmt"

/* // General slice
func main() {
	var x []int
	fmt.Println(x == nil)

	x = append(x, 10)
	fmt.Println(x)

	x = append(x, 5, 6, 7)
	fmt.Println(x)

	y := []int{20, 30, 40}
	x = append(x, y...)
	fmt.Println(x)
} */

/* // Example 3-1
func main() {
	var x []int
	fmt.Println(x, len(x), cap(x))

	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
} */

/* // make
func main() {
	x := make([]int, 5)
	fmt.Println(x, len(x), cap(x))

	x = make([]int, 0, 10)
	fmt.Println(x == nil)
	fmt.Println(x, len(x), cap(x))
	// fmt.Println(x[0])  // It's paninc
} */

/* // Example 3-2
func main() {
	var x = []int{}
	fmt.Println(x == nil)
	fmt.Println(x, len(x), cap(x))

} */

/* // Example 3-4
func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
} */

/* // Example 3-5
func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	x[1] = 20
	y[0] = 10
	z[1] = 30
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
} */

/* // Example 3-6
func main() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	fmt.Println(len(x), cap(x))
	fmt.Println(len(y), cap(y))

	y = append(y, 30)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
} */

/* // Example 3-7
func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))

	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
} */

/* // Example 3-8
func main() {
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println(cap(x), cap(y), cap(z))

	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
} */

/* // rune, byte convert to string
func main() {
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(s)
	fmt.Println(s2)

	// Go 1.15 이상 부터는 int 를 string 으로 conversion 하면 go vet 에서 에러 발생
	var x int = 65
	var y = string(x)
	fmt.Println(y)
} */

// // Converting string to run and byte
// func main() {
// 	var s string = "Hello, 🌞"
// 	var bs []byte = []byte(s)
// 	var rs []rune = []rune(s)
// 	fmt.Println(bs)
// 	fmt.Println(rs)
// }

/* // map nil value
func main() {
	var nilMap map[string]int
	fmt.Println(nilMap == nil)

	ages := make(map[int][]string, 10)
	fmt.Println(len(ages))
} */

/* //  Example 3-10 map
func main() {
	// var x []int
	// fmt.Println(x == nil)
	// x = append(x, 10)
	// fmt.Println(x)

	// // error: nil map 에는 key=value 를 추가할 수 없음
	// var totalWins map[string]int
	// fmt.Println(totalWins == nil)
	// totalWins["Orcas"] = 1

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lion"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])

	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])

	totalWins["Lion"] = 3
	fmt.Println(totalWins["Lion"])
} */

/* // map: comma ok Idion
func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodby"]
	fmt.Println(v, ok)

	delete(m, "hello")
	fmt.Println(m)
} */

/* // Set - bool
func main() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
} */

/* // Set - struct{}
func main() {
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = struct{}{}
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5 in in the set")
	}
} */

/* // struct set value zeon
func main() {
	type person struct {
		name string
		age  int
		pet  string
	}

	var fred person
	fmt.Println(fred)

	bob := person{}
	fmt.Println(bob)
} */

/* // anonymous struct
func main() {
	var person struct {
		name string
		age  int
		pet  string
	}

	person.name = "bob"
	person.age = 50
	person.pet = "dog"
	fmt.Println(person)

	pet := struct {
		name string
		kind string
	}{
		name: "Fido",
		kind: "dog",
	}
	fmt.Println(pet)
}
*/

// struct can asign and compare only anonymous struct
func main() {
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	g = f
	fmt.Println(f == g)
}
