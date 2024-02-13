package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[Handler] request received")
		// retrieve the context of the request
		rCtx := r.Context()
		// create the result channel
		resChan := make(chan int)
		// launch the func doWork in a goroutine
		go doWork(rCtx, resChan)
		// Wait for
		// 1. the client drops the connection.
		// 2. the function doWork to finish it works

		select {
		case <-rCtx.Done():
			log.Println("[Handler] context canceled in main handler, client has diconnected")
			return
		case result := <-resChan:
			log.Println("[Handler] Received 1000")
			log.Println("[Handler] Send response")
			// send data to client side
			fmt.Fprintln(w, "Response %d", result)
			return
		}
	})
	err := http.ListenAndServe("127.0.0.1:8989", nil)
	if err != nil {
		panic(err)
	}
}

func doWork(ctx context.Context, resChan chan int) {
	log.Println("[doWork] launch the doWork")
	sum := 0
	for {
		log.Println("[doWork] one iteration")
		time.Sleep(time.Microsecond)
		select {
		case <-ctx.Done():
			log.Println("[doWork] ctx Done is received inside doWork")
			return
		default:
			sum++
			if sum > 1000 {
				log.Println("[doWork] sum has reached 1000")
				resChan <- sum
				return
			}
		}
	}
}
