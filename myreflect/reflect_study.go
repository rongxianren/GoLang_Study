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
	peroson := Person{
		Name: "kloud",
		Age:  20,
		City: "shenzhen",
	}
	t := reflect.TypeOf(peroson)
	v := reflect.ValueOf(peroson)
	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf(" %s(%T) tag=%s \n", field.Name, field.Type.Name(), field.Tag.Get("attr"))
	}
}
