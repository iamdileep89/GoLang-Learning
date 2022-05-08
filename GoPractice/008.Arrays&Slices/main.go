package main

import "fmt"

func main() {
	// ARRAYS -- fixed length
	var numbers [3]int
	fmt.Println(numbers)

	coral := [4]string{"blue coral", "staghorn coral", "pillar coral", "elkhorn coral"}
	fmt.Println(coral)
	fmt.Printf("%q\n", coral)
	fmt.Println(coral[1])
	// fmt.Println(coral[18]) //panic: runtime error: index out of range
	// fmt.Println(coral[-1]) //invalid array index -1 (index must be non-negative)
	fmt.Println("Sammy loves " + coral[0])

	fmt.Printf("%q\n", coral)
	coral[1] = "foliose coral"
	fmt.Printf("%q\n", coral)
	fmt.Printf("%d\n", len(coral))

	numbers1 := [13]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	fmt.Println(len(numbers1))

	// coral = append(coral, "black coral") //first argument to append must be slice; have [4]string

	// SLICES
	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp", "anemone"}
	fmt.Printf("%q\n", seaCreatures)

	oceans := make([]string, 3)
	fmt.Printf("%q\n", oceans)

	oceans1 := make([]string, 3, 5)
	fmt.Printf("%q\n", oceans1)

	// SLICING ARRAY INTO SLICES
	fmt.Println(coral[1:3])
	fmt.Println(coral[:3])

	// Convert array to slice
	coralSlice := coral[:]
	fmt.Printf("%q\n", coralSlice)
	coralSlice = append(coralSlice, "black coral")
	fmt.Printf("%q\n", coralSlice)
	coralSlice = append(coralSlice, "antipathes", "leptopsammia")
	fmt.Printf("%q\n", coralSlice)

	// combine slices
	moreCoral := []string{"massive coral", "soft coral"}
	coralSlice = append(coralSlice, moreCoral...)
	fmt.Printf("%q\n", coralSlice)

	// removing item from slice - If i is the index of the element to be removed, then the format of this process would look like the following:
	//slice = append(slice[:i], slice[i+1:]...)
	//From coralSlice, let’s remove the item "elkhorn coral". This item is located at the index position of 3.
	coralSlice = append(coralSlice[:3], coralSlice[4:]...)
	fmt.Printf("%q\n", coralSlice)

	// remove or delete a range of items
	coralSlice = append(coralSlice[:3], coralSlice[6:]...)
	fmt.Printf("%q\n", coralSlice)

	// measure capacity cap() of slice
	numbers2 := []int{}
	for i := 0; i < 4; i++ {
		numbers2 = append(numbers2, i)
	}
	fmt.Printf("%d\n", numbers2)

	numbers3 := make([]int, 4)
	// fmt.Println(cap(numbers3)) //4
	for i := 0; i < cap(numbers3); i++ {
		numbers3[i] = i
	}
	fmt.Printf("%d\n", numbers3)

	// MULTI-DIMENSIONAL SLICES
	seaNames := [][]string{{"shark", "octopus", "squid", "mantis shrimp"}, {"Sammy", "Jesse", "Drew", "Jamie"}}
	fmt.Println(seaNames[1][0])
	fmt.Println(seaNames[0][0])

}
