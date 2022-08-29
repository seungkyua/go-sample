package main

import "fmt"

// struct set value zeon
func main() {
	type person struct {
		name string
		age  int
	}

	var jina person
	fmt.Println(jina)

	jamie := person{}
	fmt.Println(jamie)

	/* jina := person{
		name: "Jina",
		age:  21,
	}
	fmt.Println(jina)

	jamie := person{
		"Jamie",
		18,
	}
	fmt.Println(jamie) */

	/* var jina person
	jina.name = "Jina"
	jina.age = 21
	// jina["age"] = 21  // 인덱싱은 지원하지 않음

	fmt.Println(jina.name) */

	/* var jina struct {
		name string
		age  int
	}

	jina.name = "Jina"
	jina.age = 21
	fmt.Println(jina)

	jamie := struct {
		name string
		age  int
	}{
		name: "Jamie",
		age:  18,
	}
	fmt.Println(jamie) */

	/* jina := person{
		name: "Jina",
		age:  21,
	}

	jina2 := person{
		name: "Jina",
		age:  21,
	}
	fmt.Println(jina == jina2)

	jamie := person{
		name: "Jamie",
		age:  18,
	}
	fmt.Println(jina == jamie) */

	/* type anotherPerson struct {
		name string
		age  int
	}

	jina := person{
		name: "Jina",
		age:  21,
	}
	fmt.Println(jina)

	jamie := anotherPerson{
		name: "Jamie",
		age:  18,
	}

	// fmt.Println(jina == jamie)  // mismatched types person and anotherPerson

	jina = person(jamie)
	fmt.Println(jina) */

	/* jina := person{
		name: "Jina",
		age:  21,
	}
	fmt.Println(jina)

	var jamie struct {
		name string
		age  int
	}

	jamie = jina
	fmt.Println(jina == jamie) */

	/* jamie := struct {
		name string
		age  int
	}{
		name: "Jamie",
		age:  18,
	}

	var jina person

	jina = person(jamie)
	// jina = jamie
	fmt.Println(jina)
	fmt.Println(jina == jamie) */

}
