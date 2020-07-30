package main

import "fmt"

var (
	i int = 1
	f64 float64 = 1.2
	s string = "test"
	t, f bool = true, false
)

func main() {
	fmt.Println(i, f64, s, t, f)

	xi := 1
	xi = 2
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Printf("%T\n", xf32)
	fmt.Println(xi, xf32, xs, xt, xf)
}
