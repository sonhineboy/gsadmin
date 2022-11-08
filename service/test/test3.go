package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {

	type Test struct {
		A string `aa:"xss,table:aaa,where:xxx"`
	}

	var c Test

	ct := reflect.ValueOf(&c)

	t := reflect.TypeOf(c)

	b, _ := t.FieldByName("A")

	tasArr := strings.Split(b.Tag.Get("aa"), ",")

	for _, v := range tasArr {
		fmt.Println(v)
	}

	ct.Elem().FieldByName("A")

	fmt.Println("v%", ct.Interface())
}
