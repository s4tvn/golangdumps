package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	p := person{name: "Steven", age: 22}
	fmt.Println(p)

}