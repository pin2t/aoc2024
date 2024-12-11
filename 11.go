package main

import "bufio"
import "os"
import "strings"
import "strconv"
import "fmt"

func blink(numbers map[int64]int64) map[int64]int64 {
	var result = make(map[int64]int64)
	for n, count := range numbers {
		if n == 0 {
			result[1] = result[1] + count
		} else {
			var s = strconv.FormatInt(n, 10)
			if len(s) % 2 == 0 {
				var left, _ = strconv.ParseInt(s[:len(s) / 2], 10, 64)
				var right, _ = strconv.ParseInt(s[len(s) / 2:], 10, 64)
				result[left] = result[left] + count
				result[right] = result[right] + count
			} else {
				result[n * 2024] = result[n * 2024] + count
			}
		}
	}
	return result
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input = make(map[int64]int64)
	for _, s := range strings.Split(scanner.Text(), " ") {
		var n, _ = strconv.ParseInt(s, 10, 64)
		input[n] = 1
	}
	var stones = make(map[int64]int64)
	for n, c := range input { stones[n] = c }
	for b := 0; b < 25; b++ { stones = blink(stones) }
	var total int64
	for _, c := range stones { total += c }
	fmt.Print(total)
	stones = make(map[int64]int64)
	for n, c := range input { stones[n] = c }
	for b := 0; b < 75; b++ { stones = blink(stones) }
	total = 0
	for _, c := range stones { total += c }
	fmt.Println("", total)
}
