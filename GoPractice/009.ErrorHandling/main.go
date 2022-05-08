package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func capitalize(name string) (string, error) {
	if name == "" {
		return "", errors.New("no name provided")
	}
	return strings.ToTitle(name), nil
}

func capitalize1(name string) (string, int, error) {
	handle := func(err error) (string, int, error) {
		return "", 0, err
	}
	if name == "" {
		return handle(errors.New("no name provided"))
	}
	return strings.Title(name), len(name), nil
}
func main() {
	err := errors.New("ThisIsError")
	fmt.Println("Sammy says: ", err)

	err = fmt.Errorf("error occured at: %v", time.Now())
	fmt.Println("An error happend:", err)

	name, err := capitalize("samy")
	if err != nil {
		fmt.Println("Couldn't capitalize:", err)
		return
	}
	fmt.Println("Capitalized name:", name)

	name, size, err := capitalize1("sammy")
	if err != nil {
		fmt.Println("An error occured:", err)
	}
	fmt.Printf("Capitalized name: %s, length: %d\n", name, size)
}
