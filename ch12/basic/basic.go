package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)

	fmt.Println("==========>")
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	t = v.Type()
	fmt.Println(t.String())
	x := v.Interface()
	i := x.(int)
	fmt.Println(i)
	v = reflect.ValueOf("3")
	fmt.Println(v.String())

	fmt.Println("==========>")

}
