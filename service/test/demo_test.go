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
