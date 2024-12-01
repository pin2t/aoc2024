package main

import	"fmt"
import	"math"
import	"slices"

func main() {
	var left, right []int = make([]int, 0), make([]int, 0)
	for {
		var l, r int
		if n, _ := fmt.Scan(&l, &r); n < 2 {
			break
		}
		left = append(left, l)
		right = append(right, r)
	}
	slices.Sort(left)
	slices.Sort(right)
	var total, similarity int
	for i, _ := range left {
		total += int(math.Abs(float64(left[i] - right[i])))
		var c int
		for _, n := range right {
			if left[i] == n { c++ }
		}
		similarity += left[i] * c
	}
	fmt.Println(total, similarity)
}
