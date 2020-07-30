package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	//var b [2]int = [2]int{ 100, 200 }
	//b = append(b, 300) // エラー出る
	//fmt.Println(b)

	var b []int = []int{ 100, 200 }
	b = append(b, 300)
	fmt.Println(b)
	n := []int{ 1, 2, 3, 4, 5}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])
	n[2] = 100
	fmt.Println(n)

	var board = [][]int{
		[]int{ 0, 1, 2},
		[]int{ 3, 4, 5},
		[]int{ 6, 7, 8},
	}
	fmt.Println(board)
	n = append(n, 100, 200, 300)
	fmt.Println(n)
}
