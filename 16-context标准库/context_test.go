package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestWithValue(t *testing.T) {
	f := func(ctx context.Context, k string) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	ctx := context.WithValue(context.Background(), "key1", "Golang")
	f(ctx, "key1")
	f(ctx, "key2")
}

func handelrequest(ctx context.Context) {
	go writeredis(ctx)
	go writedatabase(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("handelrequest done.")
			return
		default:
			fmt.Println("handelrequest running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writeredis(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writeredis done.")
			return
		default:
			fmt.Println("writeredis running")
			time.Sleep(2 * time.Second)
		}
	}
}

func writedatabase(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("writedatabase done.")
			return
		default:
			fmt.Println("writedatabase running")
			time.Sleep(2 * time.Second)
		}
	}
}
func TestCacel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go handelrequest(ctx)

	time.Sleep(5 * time.Second)
	fmt.Println("it's time to stop all sub goroutines!")
	cancel()

	//just for test whether sub goroutines exit or not
	time.Sleep(5 * time.Second)
}

const shortDuration = 4 * time.Second

func TestDeadLine(t *testing.T) {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("在截止时间之后停止")
	case <-ctx.Done():
		fmt.Println("在截止时间停止")
	}

}

func TestDeadLineTime(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(7 * time.Second):
		fmt.Println("在超时时间之后结束")
	case <-ctx.Done():
		fmt.Println("在超时时间结束")
	}
}
