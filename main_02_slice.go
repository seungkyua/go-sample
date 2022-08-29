package main

import "fmt"

func main() {
	/* var s = []int{1, 2, 3, 4, 5}
	fmt.Println(s) */
	// var s1 []int
	/* fmt.Println(s1 == nil)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s1[0]) */
	// s1 = append(s1, 1)
	// fmt.Println(s1)

	/* s2 := []int{2, 3, 4, 5}
	s1 = append(s1, s2...)
	// s1 = append(s1, 2, 3, 4, 5)
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
	fmt.Println(s, len(s), cap(s))
	fmt.Println(s3, len(s3), cap(s3)) */

	/* s := make([]int, 0, 5)
	fmt.Println(s == nil)
	fmt.Println(s, len(s), cap(s))

	s1 := make([]int, 2)
	fmt.Println(s1) */

	/* var s = []int{}
	fmt.Println(s == nil)
	fmt.Println(s, len(s), cap(s)) */

	/* a := []int{1, 2, 3, 4, 5}
	b := a[:3:3]
	c := a[2:]
	fmt.Println("cap(a):", cap(a), "cap(b):", cap(b), "cap(c):", cap(c))
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("===================")

	c[0] = 30
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("===================")

	b = append(b, 40)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c) */

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
