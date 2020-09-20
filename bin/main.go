package main

import (
	"context"
	"fmt"
	"time"

	"github.com/liuwangchen/executor"
)

func before(ctx context.Context) error {
	fmt.Println("before begin")
	defer fmt.Println("before end")
	return nil
}

func after(ctx context.Context) error {
	fmt.Println("after begin")
	defer fmt.Println("after end")
	return nil
}

func foo(ctx context.Context) error {
	fmt.Println("foo begin")
	defer fmt.Println("foo end")
	time.Sleep(time.Second * 2)
	return nil
}

func bar(ctx context.Context) error {
	fmt.Println("bar begin")
	defer fmt.Println("bar end")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	return nil
}

func do(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			fmt.Println("zhixing")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func handleSig(ctx context.Context) error {
	fmt.Println("handle sig")
	return nil
}

func haha(fn func(int) int) {
	fmt.Println(fn)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	err := executor.Execute(
		ctx,
		executor.Ants(1, executor.Func(foo), executor.Func(foo), executor.Func(foo), executor.Func(foo)),
	)
	defer cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("done")
}
