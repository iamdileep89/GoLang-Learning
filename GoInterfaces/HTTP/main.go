package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)
	// io.Copy(os.Stdout, resp.Body)
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// reader := strings.NewReader("Clear is better than clever")
	// p := make([]byte, 4)
	// for {
	// 	n, err := reader.Read(p)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(string(p[:n]))
	// }
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
