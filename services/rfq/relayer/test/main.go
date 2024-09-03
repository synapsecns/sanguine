package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 10; i++ {
		go func() {
			select {
			case <-ctx.Done():
				//cancel()
				fmt.Println("cancel")
			case <-time.After(10 * time.Second):
				return
			}
		}()
	}

	time.Sleep(1 * time.Second)
	cancel()
}
