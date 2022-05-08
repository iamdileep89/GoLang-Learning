package main

import "fmt"

var g = "global"

func printLocal() {
	l := "local"
	fmt.Println(l)
	fmt.Println(g)
}

const (
	year     = 365
	leapYear = int32(366)
)

func main() {
	i := 656875647
	i = 65457686764
	fmt.Println(i)
	fmt.Println(i - 765)

	// s := "Hello, World!"
	// f := 45.57
	// b := 5 > 9
	// array := [4]string{"jhg", "kj", "jhg", "jh"}
	// slice := []string{"kjb", "jhg", "jhg", "kj"}
	// m := map[string]string{"letter": "g", "number": "seven", "symbol": "%"}

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a %T = %d\n", a, a)
	fmt.Printf("var a %T = %q\n", b, b)
	fmt.Printf("var a %T = %+v\n", c, c)
	fmt.Printf("var a %T = %+v\n", d, d)

	names := []string{"mary", "bob", "Anna"}
	for i, n := range names {
		fmt.Printf("index: %d = %q\n", i, n)
	}

	j, k, l := "shark", 2.045, true
	fmt.Println(j, k, l)

	printLocal()
	fmt.Println(g)
	fmt.Println(l)

	const shark = "sammy"
	fmt.Println(shark)

}
