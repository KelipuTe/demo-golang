package context

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func f8WaitForSignal(i9ctx context.Context, name string) {
	for {
		select {
		case <-i9ctx.Done():
			fmt.Println(name, "i9ctx.Done()")
			return
		default:
			fmt.Println(name, "keep wait")
			time.Sleep(1 * time.Second)
		}
	}
}

func Test_f8ContextWithValue(p7s6t *testing.T) {
	i9ctx := context.Background()

	i9ctx = context.WithValue(i9ctx, "key", "a")
	i9ctx = context.WithValue(i9ctx, "key2", 1)

	keyValue := i9ctx.Value("key").(string)
	key2Value := i9ctx.Value("key2").(int)

	fmt.Println("key type is:", reflect.TypeOf(keyValue), "value is:", keyValue)
	fmt.Println("key2 type is:", reflect.TypeOf(key2Value), "value is:", key2Value)
}

func Test_f8ContextWithCancel(p7s6t *testing.T) {
	i9ctx := context.Background()

	i9ctx, f8cancel := context.WithCancel(i9ctx)

	go f8WaitForSignal(i9ctx, "go")
	go f8WaitForSignal(i9ctx, "go2")

	time.Sleep(2 * time.Second)
	fmt.Println("before cancel:", i9ctx.Err())

	f8cancel()

	time.Sleep(2 * time.Second)
	fmt.Println("after cancel:", i9ctx.Err())
}

func f8ContextWithDeadline() {
	fmt.Println("f8ContextWithDeadline")
	i9ctx := context.Background()

	i9ctx, f8cancel := context.WithDeadline(i9ctx, time.Now().Add(4*time.Second))

	go f8WaitForSignal(i9ctx, "go")
	go f8WaitForSignal(i9ctx, "go2")

	time.Sleep(2 * time.Second)
	fmt.Println("before deadline:", i9ctx.Err())

	time.Sleep(4 * time.Second)
	fmt.Println("after deadline:", i9ctx.Err())

	f8cancel()
	fmt.Println("after cancel:", i9ctx.Err())
}

func f8ContextWithDeadlineV2() {
	fmt.Println("f8ContextWithDeadlineV2")
	i9ctx := context.Background()

	i9ctx, f8cancel := context.WithDeadline(i9ctx, time.Now().Add(4*time.Second))

	go f8WaitForSignal(i9ctx, "go")
	go f8WaitForSignal(i9ctx, "go2")

	time.Sleep(2 * time.Second)
	fmt.Println("before deadline:", i9ctx.Err())

	f8cancel()
	fmt.Println("after cancel:", i9ctx.Err())

	time.Sleep(4 * time.Second)
	fmt.Println("after deadline:", i9ctx.Err())
}

func Test_f8ContextWithDeadline(p7s6t *testing.T) {
	f8ContextWithDeadline()
	f8ContextWithDeadlineV2()
}

func f8ContextWithTimeout() {
	fmt.Println("f8ContextWithTimeout")
	i9ctx := context.Background()

	i9ctx, f8cancel := context.WithTimeout(i9ctx, 4*time.Second)

	go f8WaitForSignal(i9ctx, "go")
	go f8WaitForSignal(i9ctx, "go2")

	time.Sleep(2 * time.Second)
	fmt.Println("before timeout:", i9ctx.Err())

	time.Sleep(4 * time.Second)
	fmt.Println("after timeout:", i9ctx.Err())

	f8cancel()
	fmt.Println("after cancel:", i9ctx.Err())
}

func f8ContextWithTimeoutV2() {
	fmt.Println("f8ContextWithTimeoutV2")
	i9ctx := context.Background()

	// WithTimeout 里面其实是调了 WithDeadline
	i9ctx, f8cancel := context.WithTimeout(i9ctx, 4*time.Second)

	go f8WaitForSignal(i9ctx, "go")
	go f8WaitForSignal(i9ctx, "go2")

	time.Sleep(2 * time.Second)
	fmt.Println("before timeout:", i9ctx.Err())

	f8cancel()
	fmt.Println("after cancel:", i9ctx.Err())

	time.Sleep(4 * time.Second)
	fmt.Println("after timeout:", i9ctx.Err())
}

func Test_f8ContextWithTimeout(p7s6t *testing.T) {
	f8ContextWithTimeout()
	f8ContextWithTimeoutV2()
}
