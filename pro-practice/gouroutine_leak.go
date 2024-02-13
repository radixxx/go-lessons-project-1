package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

func main() {
	log.Println("begin program")
	go launch()
	time.Sleep(time.Millisecond)
	log.Printf("Gouroutine count: %d\n", runtime.NumGoroutine())
	for {
	}
}

func doGouroutineA(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("first goroutine return")
		return
	}
}

func doGouroutineB(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("second goroutine return")
		return
	}
}

func launch() {

	//	How to create an empty context?

	ctx := context.Background()
	ctx, _ = context.WithCancel(ctx)
	log.Println("launch first goroutine")
	go doGouroutineA(ctx)
	go doGouroutineA(ctx)
	go doGouroutineA(ctx)
	log.Println("launch second goroutine")
	go doGouroutineB(ctx)
	go doGouroutineB(ctx)
	go doGouroutineB(ctx)
}
