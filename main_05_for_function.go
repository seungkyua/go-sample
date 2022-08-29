package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	/* if num := 5; num == 0 {
		fmt.Println("False")
	} else if num < 10 {
		fmt.Println("True")
	} else {
		fmt.Println("Big number")
	}
	fmt.Println(num) */

	/* num := 10
	if num > 0 {
		fmt.Println("outer num =", num)
		num := 0
		fmt.Println("inner num =", num)
	}
	fmt.Println("outer num =", num) */

	/* nums := []int{10, 20, 30, 40, 50}
	for i, v := range nums {
		fmt.Println("index =", i, ",", "value =", v)
	}
	for _, v := range nums {
		fmt.Println("value =", v)
	}
	for i := range nums {
		fmt.Println("index =", i)
	} */

	/* m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for k, v := range m {
		fmt.Println("key =", k, ",", "value =", v)
	}
	for _, v := range m {
		fmt.Println("value =", v)
	}
	for k := range m {
		fmt.Println("key =", k)
	} */

	/* alpha := []string{"a", "ab", "abc", "abcd", "abcde", "abcdef", "abedefg"}
	for _, a := range alpha {
		switch l := len(a); l {
		case 1, 2:
			fmt.Println(a, "length = 2 or 3")
		case 3:
			fmt.Println(a, "length = 3")
		case 4, 5, 6:
			fmt.Println(a, "length = 4, 5 or 6")
		default:
			fmt.Println(a, "length > 6")
		}
	} */

	// ===================================================
	// function
	// ===================================================
	/* result := div(4, 2)
	fmt.Println(result) */

	/* fmt.Println(addTo(1))
	fmt.Println(addTo(1, 2))
	fmt.Println(addTo(1, 2, 3))
	i := []int{4, 5}
	fmt.Println(addTo(3, i...))
	fmt.Println(addTo(6, []int{7, 8, 9}...)) */

	/* result, remainder, err := multiReturnValues(10, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result, remainder) */

	/* // function is variable
	exp := []string{"3", "+", "4"}
	result, err := cal(exp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result) */

	/* // Anonymous Function
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println(j)
		}(i)
	} */

	/* // function 안에 function 을 선언한 것을 closure 라 한다.
	// high-order function 은 input 파라미터나 리턴값으로 function 이 사용되는 것을 의미한다.
	funcA := multiple(2)
	fmt.Println(funcA(3))
	funcA = multiple(4)
	fmt.Println(funcA(5)) */

	/* // Call by Value
	i := 1
	s := "Hello"
	p := person{}
	modify(i, s, p)
	fmt.Println(i, s, p) */

	// Call by Value but slice and map like pointer
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2}
	modSlice(s)
	fmt.Println(s)

}

// func div(n int, d int) int {
func div(n, d int) int {
	if d == 0 {
		return 0
	}
	return n / d
}

// 가변인자 (variadic parameter)
func addTo(base int, vals ...int) []int {
	result := make([]int, 0, len(vals))
	for _, v := range vals {
		result = append(result, base+v)
	}
	return result
}

// multiple retrun values
func multiReturnValues(n, d int) (int, int, error) {
	if d == 0 {
		return 0, 0, errors.New("cannot divid by zero")
	}
	return (n / d), (n % d), nil
}

func cal(exp []string) (int, error) {
	/* opMap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
	} */

	// Function Type
	type opFuncType func(int, int) int
	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
	}

	a, err := strconv.Atoi(exp[0])
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	b, err := strconv.Atoi(exp[2])
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	op := exp[1]
	opFunc, ok := opMap[op]
	if !ok {
		fmt.Println("unsupported operator: ", op)
		return 0, err
	}
	return opFunc(a, b), nil
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// high-ooder fuctnion
func multiple(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

// Call by Value
type person struct {
	name string
	age  int
}

func modify(i int, s string, p person) {
	i = i + 1
	s = "modified"
	p.name = "ask"
	p.age = 100
}

// Call by Value but slice and map like pointer
func modMap(m map[string]int) {
	m["two"] = 20
	m["three"] = 3
	delete(m, "one")
}

func modSlice(s []int) {
	s[0] = s[0] * 10
	s = append(s, 3)
	s[1] = s[1] * 10
}
