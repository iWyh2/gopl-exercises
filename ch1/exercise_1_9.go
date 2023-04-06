package ch1

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func FetchModified3() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			n, err := fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			if err != nil {
				fmt.Printf("%d bytes written.\n", n)
			}
			os.Exit(1)
		}
		fmt.Println(resp.Status)
		n, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Printf("%d bytes written.\n", n)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		err = resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
		}
		if err != nil {
			n, err := fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			if err != nil {
				fmt.Printf("%d bytes written.\n", n)
			}
			os.Exit(1)
		}
		//fmt.Printf("%s", b)
	}
}
