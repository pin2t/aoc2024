package main

import "fmt"

func main() {
	var sum = int64(0)
	var prices = make(map[int]int)
	for {
		var n int64
		if _, err := fmt.Scanf("%d", &n); err != nil { break }
		var changes []int
		var processed = make(map[int]bool)
		for i := 0; i < 2000; i++ {
			var nx = next(n)
			changes = append(changes, int(nx % 10 - n % 10))
			if i > 2 {
				var key = changes[i - 3] * 1000000 + changes[i - 2] * 10000 + changes[i - 1] * 100 + changes[i]
				if !processed[key] {
					prices[key] = prices[key] + int(nx % 10)
					processed[key] = true
				}
			}
			n = nx
		}
		sum += n
	}
	var m = 0
	for _, p := range prices {
		if p > m { m = p }
	}
	fmt.Println(sum, m)
}

func next(n int64) int64 {
	var n1 = ((n * 64) ^ n) % 16777216
	var n2 = ((n1 / 32) ^ n1) % 16777216
	return ((n2 * 2048) ^ n2) % 16777216
}