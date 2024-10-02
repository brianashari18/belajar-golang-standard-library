package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name, Address, Email, Gender string
}

type Animal struct {
	Name string `required:"true"`
	Type string `required:"true"`
	Age  int    `required:"true"`
}

func readField(value any) {
	var structType reflect.Type = reflect.TypeOf(value)
	var dataValue reflect.Value = reflect.ValueOf(value)
	fmt.Println("Type/struct name: ", structType.Name())
	fmt.Println("Data name: ", dataValue)
	for i := 0; i < structType.NumField(); i++ {
		var structField = structType.Field(i)
		var dataField = dataValue.Field(i)
		fmt.Println("Struct field: ", structField)
		if tagValue, ok := structField.Tag.Lookup("max"); ok {
			fmt.Println("Max: ", tagValue)
		}
		fmt.Println(structField.Name, ": ", dataField.Interface())

	}
	fmt.Println()
}

func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			if data := reflect.ValueOf(value).Field(i).Interface(); data == "" {
				return false
			}

		}
	}
	return true
}

func main() {
	readField(Sample{"Brian"})
	readField(Person{"Brian", "Jakarta", "----gmail.com", "Male"})
	fmt.Println(isValid(Sample{""}))
	fmt.Println(isValid(Animal{"", "Dog", 5}))
}
