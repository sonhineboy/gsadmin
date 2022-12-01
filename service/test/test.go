package main

import (
	"fmt"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"reflect"
)

type Test struct {
	A string
}

func (t Test) Add() {
	fmt.Println(t.A)
}

func (t Test) Add2(s models.AdminMenu) {
	s.Test()
	fmt.Println(t.A)
}
func main() {

	var t map[string]interface{}

	t = make(map[string]interface{})
	t["AdminMenu"] = models.AdminMenu{}

	var a Test = Test{
		A: "xxxx",
	}

	ra := reflect.ValueOf(a)

	fmt.Println(reflect.TypeOf(ra.Interface()))

	add2 := ra.Method(1)
	o := t[add2.Type().In(0).Name()]

	var args []reflect.Value
	args = append(args, reflect.ValueOf(o))
	add2.Call(args)

}
