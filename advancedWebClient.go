package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ping http://www.any-site-name.com

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing: ", err)
		return
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		fmt.Println("GET: ", err)
		return
	}

	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in Do()", err)
		return
	}

	fmt.Println("Staus code: ", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))

	contentType := httpData.Header.Get("Content-type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character set : ", characterSet[1])
	}

	if httpData.ContentLength == -1 {
		fmt.Println("Content Lenght: ", httpData.ContentLength)
	}

	lenght := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read((buffer[0:]))
		if err != nil {
			fmt.Println(err)
			break
		}
		lenght = lenght + n
	}
	fmt.Println("Calculated response data length:", lenght)
}
