package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Fullname string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	typeof := reflect.TypeOf(data)

	for i := 0; i < typeof.NumField(); i++ {
		if typeof.Field(i).Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			return value != ""
		}
	}

	return true
}

func main() {
	sample := Sample{"eko"}

	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min"))

	sample2 := Sample{}
	fmt.Println(IsValid(sample2))
	fmt.Println(IsValid(sample))
}
