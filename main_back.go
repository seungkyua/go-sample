package main

import "fmt"

func main() {
	// fmt.Println("Hello World !!!")

	/* var isDev bool
	var isTrue = true
	fmt.Println(isDev)
	fmt.Println(isTrue)

	var x int
	var y int = 1
	var z = 2
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	var p float64
	var q = 0.0
	fmt.Println(p)
	fmt.Println(q) */

	/* var x int = 10
	var p float64 = 20.2
	fmt.Println(x + int(p))
	fmt.Println(float64(x) + p) */

	/* var r rune = 'r'
	var s string = string(r)
	var b byte = 'b'
	var s2 string = string(b)
	fmt.Println(r)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(s2) */

	/* var s string = "Hello, 😀"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs) */

	/* var a, b int = 1, 2
	var c, d = 3, "a"

	i := 10
	x, s := 20.5, "Hello World"

	fmt.Println(a, b, c, d)
	fmt.Println(i, x, s) */

	/* var s = []int{1, 2, 3, 4, 5}
	fmt.Println(s)

	var s1 []int
	fmt.Println(s1 == nil)
	fmt.Println(len(s1), cap(s1))

	// fmt.Println(s1[0])   // panic 에러
	// fmt.Println(len(s1), cap(s1))

	fmt.Printf("address s1: %p\n", &s1)
	s1 = append(s1, 1)
	fmt.Printf("address s1: %p\n", &s1)
	fmt.Println(s1)
	s2 := []int{2, 3, 4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1) */

	/* var s []int
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 2)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 3)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 4)
	fmt.Println(s, len(s), cap(s))
	s3 := append(s, 5)
	fmt.Println("s:", s, len(s), cap(s))
	fmt.Println("s3", s3, len(s3), cap(s3)) */

	/* s := make([]int, 0, 5)
	fmt.Println(s == nil)
	fmt.Println(s, len(s), cap(s))

	s1 := make([]int, 2)
	fmt.Println(s1)

	s1 = append(s1, 1)
	fmt.Println(s1) */

	/* a := []int{1, 2, 3, 4, 5}
	b := a[:3:3]
	c := a[2:]
	fmt.Println("cap(a):", cap(a), "cap(b):", cap(b), "cap(c):", cap(c))
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Println("===== (c[0] = 30) =====")
	c[0] = 30
	fmt.Println("cap(a):", cap(a), "cap(b):", cap(b), "cap(c):", cap(c))
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	fmt.Println("===== append(b, 40) ======")
	b = append(b, 40)
	b[0] = 10
	fmt.Println("cap(a):", cap(a), "cap(b):", cap(b), "cap(c):", cap(c))
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c) */

	/* var m map[string]int
	// m["one"] = 0
	fmt.Println(m == nil)

	var m1 = map[string]int{}
	fmt.Println(m1 == nil)
	m1["one"] = 0
	fmt.Println(m1)

	m2 := map[string]int{
		"one":  1,
		"two": 2,
	}
	fmt.Println(m2) */

	/* m := make(map[string]int, 5)
	fmt.Println(len(m))

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	fmt.Println(m, len(m))

	m["six"] = 6
	fmt.Println(m, len(m)) */

	/* m := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m["three"])

	v, ok := m["one"]
	fmt.Println(v, ok)

	v, ok = m["three"]
	fmt.Println(v, ok)

	delete(m, "one")
	fmt.Println(m)
	delete(m, "three") */

	/* m := map[string]int{
		"one": 1,
		"two": 2,
	}
	m["two"] = 3
	fmt.Println(m) */

	n := []int{1, 2, 3, 4}
	fmt.Println(len(n), cap(n))

	AddOne(n)
	fmt.Println("AddOne:", n)

	n = []int{1, 2, 3, 4}
	AddOneWithAppend(n)
	fmt.Println("AddOneWithAppend:", n)

	n = make([]int, 4, 5)
	n[0] = 1
	n[1] = 2
	n[2] = 3
	n[3] = 4
	AddOneWithAppend(n)
	fmt.Println("AddOneWithAppend2:", n)

	n = []int{1, 2, 3, 4}
	AddOneWithPorinter(&n)
	fmt.Println("AddOneWithPorinter:", n)

}

func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

func AddOneWithAppend(nums []int) {
	nums = append(nums, 5)
	for i := range nums {
		nums[i]++
	}
}

func AddOneWithPorinter(nums *[]int) {
	*nums = append(*nums, 5)
	for i := range *nums {
		(*nums)[i]++
	}
}
