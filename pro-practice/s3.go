package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Using default port number: ", PORT)
	} else {
		PORT = ":" + arguments[1]
	}
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	g := time.Now().Local().Format(time.Kitchen)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)

	KitchenTime := "Kitcehn time is: "
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", KitchenTime)
	fmt.Fprintf(w, "<h2 alieg=\"center\">%s</h2>", g)
	fmt.Fprintf(w, "Serving time: %s\n", r.URL.Path)
	fmt.Printf("Serving kitchen time: %s\n", r.Host)
}
