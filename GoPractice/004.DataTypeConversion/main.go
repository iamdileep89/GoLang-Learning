package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// little to big
	var index int8 = 15
	var bigIndex int32 = int32(index)

	fmt.Println(bigIndex)

	fmt.Printf("index data type: %T\n", index)
	fmt.Printf("bigIndex data type: %T\n", bigIndex)

	// big to little
	var big int64 = 64
	fmt.Println(int8(big))

	// int to float
	var x int64 = 57
	var y float64 = float64(x)
	fmt.Printf("%.2f\n", y)
	fmt.Printf("%.5f\n", y)

	//float to int
	var f float64 = 657.567
	var i int = int(f)
	fmt.Printf("f = %.2f\n", f)
	fmt.Printf("i = %d\n", i)

	// convert by division
	a := 5 / 2
	fmt.Println(a)
	b := 5.0 / 2
	fmt.Println(b)

	// numbers to strings
	j := strconv.Itoa(12)
	fmt.Printf("%q\n", j)

	user := "Sammy"
	lines := 50
	fmt.Println("Congratulations, " + user + "! You just wrote " + strconv.Itoa(lines) + " lines of code.")

	k := 43657.65 // float value
	fmt.Println("Sammy has " + fmt.Sprint(k) + " points")

	// strings to numbers
	lines_yesterday := "50"
	lines_today := "10823"

	yesterday, err := strconv.Atoi(lines_yesterday)
	if err != nil {
		log.Fatal(err)
	}
	today, err := strconv.Atoi(lines_today)
	if err != nil {
		log.Fatal(err)
	}
	lines_more := today - yesterday
	fmt.Println(lines_more)

	// convering strings and bytes, strings in go are stored as slice of bytes
	l := "my string"
	m := []byte(l)
	n := string(l)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
}
