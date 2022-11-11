package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
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

	//path := "./upload/" + time.Now().Format(string("20060102"))
	//
	//err := os.MkdirAll(path, os.ModePerm)
	//
	//fmt.Println(err)

	var wait sync.WaitGroup
	wait.Add(50)

	for i := 0; i < 50; i++ {
		time.Sleep(time.Millisecond * 10)
		go func() {

			r, _ := http.Get("http://gsadmin.api.suiyidian.cn/api/common/captcha/info")
			b, _ := ioutil.ReadAll(r.Body)
			fmt.Println(string(b))

			wait.Done()
		}()

	}

	wait.Wait()

}
