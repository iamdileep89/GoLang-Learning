package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Sammy Shark"
	fmt.Println(strings.ToUpper(s))
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.HasPrefix(s, "Sammy"))
	fmt.Println(strings.HasSuffix(s, "Shark"))
	fmt.Println(strings.Contains(s, "Sh"))
	fmt.Println(strings.Count(s, "S"))
	fmt.Println(len(s))
	fmt.Println(strings.Join([]string{"sharks", "oceans", "powerful"}, ", "))
	fmt.Println(strings.Split(s, " "))
	data := " username password    email     date"
	fields := strings.Fields(data)
	fmt.Printf("%q\n", fields)
	balloon := "Sammy has a balloon."
	fmt.Println(strings.ReplaceAll(balloon, "has", "had"))
}
