package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/liuwangchen/executor"
)

func before(ctx context.Context) error {
	log.Println("before begin")
	defer log.Println("before end")
	return nil
}

func after(ctx context.Context) error {
	log.Println("after begin")
	defer log.Println("after end")
	return nil
}

func foo(ctx context.Context) error {
	log.Println("foo begin")
	defer log.Println("foo end")
	time.Sleep(time.Second * 2)
	return nil
}

func bar(ctx context.Context) error {
	log.Println("bar begin")
	defer log.Println("bar end")
	for i := 0; i < 10; i++ {
		log.Println(i)
	}
	return nil
}

func main() {
	err := executor.Execute(
		context.Background(),
		executor.Timeout(time.Second*3, executor.ExecutorFunc(func(ctx context.Context) error {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("time out")
					return ctx.Err()
				default:
					fmt.Println("zhixing")
					time.Sleep(time.Second * 1)
				}
			}
		})),
		executor.WithBefore(executor.ExecutorFunc(before)),
		executor.WithAfter(executor.ExecutorFunc(after)),
		executor.WithSignal(syscall.SIGINT, executor.ExecutorFunc(func(ctx context.Context) error {
			return errors.New("haha")
		})),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(123)
}
