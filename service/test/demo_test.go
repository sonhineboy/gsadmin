package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestStringTo(t *testing.T) {

	b := sync.WaitGroup{}
	a := "xxx"
	b.Add(1)
	go func() {
		fmt.Println(a)
		b.Done()
	}()

	b.Wait()
}

type b struct {
	Name string
}

func TestP(t *testing.T) {
	var a *b
	a = &b{"asfdfd"}

	//fmt.Println(a)
	fmt.Printf("%p", a)

}
