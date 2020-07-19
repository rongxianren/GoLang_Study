package myreflect

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `attr:"name"`
	Age  int    `attr:"age"`
	City string `attr:"city"`
}

func PrintStructInfo() {
	person := Person{
		Name: "kloud",
		Age:  20,
		City: "shenzhen",
	}
	t := reflect.TypeOf(person)
	v := reflect.ValueOf(person)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf(" %s(%T) tag=%s \n", field.Name, field.Type.Name(), field.Tag.Get("attr"))
	}
}
