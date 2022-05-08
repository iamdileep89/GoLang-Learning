package main

import (
	"fmt"
	"sort"
)

func main() {
	sammy := map[string]string{
		"name":     "Sammy",
		"animal":   "Shark",
		"color":    "Blue",
		"location": "ocean",
	}
	fmt.Println(sammy)
	fmt.Println(len(sammy))
	fmt.Println(sammy["color"])
	fmt.Println(sammy["animal"])
	fmt.Println(sammy["location"])
	fmt.Println(sammy["name"])

	for key, value := range sammy {
		fmt.Printf("%q is the key for the value %q\n", key, value)
	}

	keys := []string{}

	for key := range sammy {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	fmt.Println(len(keys))

	sort.Strings(keys)
	fmt.Println(keys)

	items := make([]string, len(sammy))
	var i int
	for _, v := range sammy {
		items[i] = v
		i++
	}
	fmt.Printf("%q\n", items)

	counts := map[string]int{}
	fmt.Println(counts["sammy"])

	count, ok := counts["sammy"]
	if ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy has not found")
	}

	if count, ok := counts["sammy"]; ok {
		fmt.Printf("Sammy has a count of %d\n", count)
	} else {
		fmt.Println("Sammy has not found")
	}

	// adding or modifying maps
	usernames := map[string]string{"Sammy": "sammy-shark", "Jamie": "mantisshrimp54"}
	usernames["Drew"] = "squidly"
	fmt.Println(usernames)

	//deleting
	permissions := map[int]string{1: "read", 2: "write", 4: "delete", 8: "create", 16: "modify"}
	fmt.Println(permissions)
	delete(permissions, 16)
	fmt.Println(permissions)
	// to remove/delete all items
	permissions = map[int]string{}
	fmt.Println(permissions)
}
