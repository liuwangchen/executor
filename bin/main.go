package main

import (
	"context"
	"fmt"
	"os"
	"syscall"
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
	os.Exit(1)
	return nil
}

func main() {
	err := executor.Execute(
		context.Background(),
		executor.Timeout(time.Second*10, executor.ExecutorFunc(do)),
		executor.WithBefore(executor.ExecutorFunc(before)),
		executor.WithAfter(executor.ExecutorFunc(after)),
		executor.WithSignal(syscall.SIGINT, executor.ExecutorFunc(handleSig)),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("done")
}
