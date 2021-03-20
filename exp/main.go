package main

import "fmt"

type MyStruct struct {
	a int
}

func (m *MyStruct) ChangeA(a int) {
	m.a = a
}

func main() {
	m := MyStruct{a: 10}
	m.ChangeA(15)

	fmt.Println(m)

}
