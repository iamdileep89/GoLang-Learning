package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(1 + 5)
	a := 5.5
	b := 2.5
	fmt.Printf("%.2f\n", a+b)

	i := 2.2
	j := -19
	k := 3.6
	l := -30
	fmt.Println(+i)
	fmt.Println(+j)
	fmt.Println(-k)
	fmt.Println(-l)

	s := 80
	t := 6
	r := float64(s) / float64(t)
	fmt.Println(r)

	o := 85
	p := 15
	fmt.Println(o % p)

	m := 36.0
	n := 8.0
	u := math.Mod(m, n)
	fmt.Println(u)

	w := 5
	w += 1
	fmt.Println(w)

	values := []int{0, 1, 2, 3, 4, 5}
	for _, x := range values {
		w := x
		w *= 2
		fmt.Println(w)
	}
}
