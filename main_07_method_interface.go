package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s, %s, %d", p.FirstName, p.LastName, p.Age)
}

func main() {
	/* p := Person{
		FirstName: "Seungkyu",
		LastName:  "Ahn",
		Age:       52,
	}
	fmt.Println(p.String()) */

	/* var t = Timer{scheduled: time.Now()}
	fmt.Println(t.String())
	t.AddMinute(10)
	fmt.Println(t.String())
	(&t).AddMinute(10)
	fmt.Println(t.String()) */

	/* adder := Adder{start: 10}
	fmt.Println(adder.Add(10))

	func1 := adder.Add
	fmt.Println(func1(10))

	func2 := Adder.Add
	fmt.Println(func2(adder, 10)) */

	/* type Animal int
	const (
		dog Animal = iota
		cat
		cow
		pig
	)
	fmt.Println(dog, cat, cow, pig) */

	/* c := Client{
		L: MyLogic{},
	}
	c.Program() */

	/* var s *string
	fmt.Println(s == nil)
	var i interface{}
	fmt.Println(i == nil)
	i = s
	fmt.Println(i == nil) */

	/* var i interface{}
	var mi MyInt = 10
	i = mi
	fmt.Println(i.(MyInt) + 1)

	i2, ok := i.(int)
	if !ok {
		fmt.Println(fmt.Errorf("unexpected type for %v", i))
		return
	}
	fmt.Println(i2 + 1) */

	/* var a interface{}
	check(a)

	var b int
	check(b)

	var c MyInt
	check(c)

	var d bool
	check(d)

	var e rune
	check(e)

	var f MyRead
	check(f)

	var g string
	check(g) */

	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.Greeting)
	http.ListenAndServe(":7777", nil)

}

type Timer struct {
	scheduled time.Time
}

func (t *Timer) AddMinute(m int) {
	t.scheduled = t.scheduled.Add(time.Minute * time.Duration(m))
}

func (t Timer) String() string {
	return fmt.Sprintf("scheduled: %v", t.scheduled)
}

type Adder struct {
	start int
}

func (a Adder) Add(val int) int {
	return a.start + val
}

type Logic interface {
	Process(s string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	p := c.L.Process("My Process")
	fmt.Println(p)
}

type MyLogic struct{}

func (ml MyLogic) Process(s string) string {
	return fmt.Sprintf("%s is working", s)
}

type MyInt int

type MyRead struct{}

func (m MyRead) Read(p []byte) (n int, err error) {
	return 0, nil
}

func check(i interface{}) {
	switch j := i.(type) {
	case nil:
		fmt.Printf("j: %T\n", j)
	case int:
		fmt.Printf("j: %T\n", j)
	case MyInt:
		fmt.Printf("j: %T\n", j)
	case io.Reader:
		fmt.Printf("j: %T\n", j)
	case string:
		fmt.Printf("j: %T\n", j)
	case bool, rune:
		fmt.Printf("j: %T\n", j)
	default:
		fmt.Printf("j: %T\n", j)
	}
}

// Using Interface
type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

func LogOutput(message string) {
	fmt.Println(message)
}

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{
		userData: map[string]string{
			"1": "Fred",
			"2": "Mary",
			"3": "Pat",
		},
	}
}

// business logic
type BusinessLogic interface {
	SayHello(userID string) (string, error)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello, " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye, " + name, nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type Controller struct {
	l     Logger
	logic BusinessLogic
}

func (c Controller) Greeting(w http.ResponseWriter, r *http.Request) {
	c.l.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, bl BusinessLogic) Controller {
	return Controller{
		l:     l,
		logic: bl,
	}
}
