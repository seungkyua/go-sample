package main

func main() {
	/* var m map[string]int
	fmt.Println(m == nil) */

	/* // nil error
	var m map[string]int
	m["one"] = 0 */

	/* // empty map 생성 => 리터럴
	var m1 = map[string]int{}
	fmt.Println(m1 == nil)
	m1["one"] = 0
	fmt.Println(m1) */

	/* // map 초기값 생성
	m2 := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println(m2) */

	/* // empty map 생성 => make, 길이:0, 크기:5
	m := make(map[string]int, 5)
	fmt.Println(len(m))
	fmt.Println("====================")

	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	m["four"] = 4
	m["five"] = 5
	fmt.Println(m, len(m))
	fmt.Println("====================")

	m["six"] = 6
	fmt.Println(m, len(m))
	fmt.Println("====================")
	// 없는 key 값 호출 => zero value return
	fmt.Println(m["seven"]) */

	/* // comma ok idiom
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	v, ok := m["one"]
	fmt.Println(v, ok)
	fmt.Println("====================")

	v, ok = m["three"]
	fmt.Println(v, ok) */

	/* // map item 삭제, 없는 key 삭제 => no error
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	delete(m, "one")
	fmt.Println(m)
	fmt.Println("====================")

	delete(m, "three")
	fmt.Println(m) */

	/* // map item 값 변경
	m := map[string]int{
		"one": 1,
		"two": 2,
	}
	m["two"] = 3
	fmt.Println(m) */

}
