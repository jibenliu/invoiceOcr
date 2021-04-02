package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func hardWork(job interface{}) error {
	time.Sleep(time.Minute)
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	//done := make(chan error) // 这一行会出现内存泄漏
	done := make(chan error, 1)
	go func() {
		done <- hardWork(job)
	}()
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	const total = 1000
	var wg sync.WaitGroup
	wg.Add(total)
	now := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			requestWork(context.Background(), "any")
		}()
	}
	wg.Wait()
	fmt.Println("elapsed:", time.Since(now))
	fmt.Println("number of goroutines:", runtime.NumGoroutine())
}
