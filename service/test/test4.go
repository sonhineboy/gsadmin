package main

import (
	"fmt"
	"os"
	"time"
)

type Demo struct {
	a, b int
}

func (d *Demo) Add() {
	fmt.Println(d.a + d.b)
}

func (d *Demo) New(a, b int) *Demo {
	return &Demo{a: a, b: b}
}

type Demo2 struct {
	a *Demo
}

func (d2 *Demo2) Cell() {

	d2.a.Add()
}

func main() {
	//
	//c := &Demo2{}
	//
	//
	//c.a = c.a.New(1,2)
	//c.Cell()

	path := "./upload/" + time.Now().Format(string("20060102"))

	err := os.MkdirAll(path, os.ModePerm)

	fmt.Println(err)

}
