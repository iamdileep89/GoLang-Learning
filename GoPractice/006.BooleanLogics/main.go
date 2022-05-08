package main

import "fmt"

func main() {
	x := 5
	y := 8
	fmt.Println("x==y: ", x == y)
	fmt.Println("x!=y: ", x != y)
	fmt.Println("x<y: ", x < y)
	fmt.Println("x>y: ", x > y)
	fmt.Println("x<=y: ", x <= y)
	fmt.Println("x>=y: ", x >= y)

	Sammy := "Sammy"
	sammy := "sammy"
	fmt.Println("Sammy === sammy: ", Sammy == sammy)

	fmt.Println((9 > 7) && (2 < 4))   // Both original expressions are true
	fmt.Println((8 == 8) || (6 != 6)) // One original expression is true
	fmt.Println(!(3 <= 1))            // The original expression is false

	grade := 80
	if grade >= 65 {
		fmt.Println("Passing Grade!")
	} else {
		fmt.Println("Failing Grade!")
	}
}
