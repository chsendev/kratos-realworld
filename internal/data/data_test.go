package data

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ctx1 := context.WithValue(ctx, "k1", 1)

	ctx2 := context.WithValue(ctx1, "k2", 2)
	ctx3 := context.WithoutCancel(ctx2)

	fmt.Println(ctx2.Value("k1"))
	fmt.Println(ctx2.Value("k2"))

	go func() {
		select {
		case <-ctx3.Done():
			fmt.Println("down")
		}
	}()

	time.Sleep(time.Second * 2)
	cancel()
	time.Sleep(time.Second * 2)
}
