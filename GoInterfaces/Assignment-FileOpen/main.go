package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 99999)
	count, err1 := file.Read(data)
	if err != nil {
		log.Fatal(err1)
	}
	fmt.Println(string(data))
	fmt.Println(count)
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
	// io.Copy(os.Stdout, file)
}
