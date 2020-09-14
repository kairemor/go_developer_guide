package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.fr")

	log := logWriter{}
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	// io.Copy(os.Stdout, resp.Body)

	// bs := make([]byte, 24*1000)
	// n, err := resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// fmt.Println("The number of line", n)
	io.Copy(log, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
