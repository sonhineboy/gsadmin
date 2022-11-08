package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wp sync.WaitGroup

func main() {

	ch := make(chan int)
	ctx, ccc := context.WithTimeout(context.Background(), time.Second*5)
	defer ccc()
	wp.Add(1)
	go work(ctx, ch)
	for v := range ch {
		fmt.Println(v)
	}
	wp.Wait()
}

func work(ctx context.Context, ch chan int) {

	var i int = 1

loop:
	for i < 9 {
		select {
		case <-ctx.Done():
			fmt.Println("超时退出")
			close(ch)
			break loop
		default:
			fmt.Println("休息3秒")
			fmt.Println(i)
			ch <- i
			time.Sleep(time.Second * 3)
			i++
		}

	}

	fmt.Println("退出了")
	wp.Done()

}
