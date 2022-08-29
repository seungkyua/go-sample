package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	/* var i int64 = 10
	var b bool = true
	pointerI := &i
	pointerB := &b
	var pointerS *string

	fmt.Println("pointerI =", *pointerI)
	fmt.Println("pointerB =", *pointerB)
	fmt.Println("pointerS =", pointerS) */

	/* i := 10
	pointerI := &i
	fmt.Println("memory address =", pointerI)
	fmt.Println("value =", *pointerI)

	var i2 *int
	fmt.Println(i2 == nil)
	fmt.Println(*i2) // panic */

	/* var i = new(int)      // int 의 zero value 를 pointer 로 생성함
	fmt.Println(i == nil) // false
	fmt.Println(*i) */

	/* type Person struct {
		Name string
		Age  int
	}
	s := &Person{Name: "Seungkyu", Age: 50}
	fmt.Println(s)
	fmt.Println(*s)
	fmt.Println(s.Name)
	fmt.Println((*s).Name)

	p := &Person{}
	fmt.Println(p.Name, p.Age) */

	/* type Person struct {
		Name *string
		Age  int
	}
	// s := &Person{Name: &"Seungkyu"} // 에러
	// fmt.Println(s.Name)

	// n := "seungkyu"
	// s := &Person{Name: &n}
	// fmt.Println(*s.Name)

	s := &Person{Name: stringToPointer("seungkyu")}
	fmt.Println(*s.Name) */

	/* i := 10
	updateAddress(&i)
	fmt.Println(i)
	updateValue(&i)
	fmt.Println(i) */

	/* p := Person{}
	newPersonBad(&p)
	fmt.Println(p)

	p, err := newPersonGood()
	if err != nil {
		panic(nil)

	}
	fmt.Println(p) */

	p := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	if err := json.Unmarshal([]byte(`{"name": "seungkyu", "age": 52}`), &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

}

func stringToPointer(s string) *string {
	return &s
}

func updateAddress(i *int) {
	a := 50
	i = &a
}

func updateValue(i *int) {
	*i = 50
}

type Person struct {
	Name string
	Age  int
}

func newPersonBad(p *Person) error {
	p.Name = "seungkyu"
	p.Age = 52
	return nil
}

func newPersonGood() (Person, error) {
	p := Person{
		Name: "seungkyu",
		Age:  52,
	}
	return p, nil
}
