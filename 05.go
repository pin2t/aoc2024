package main

import "bufio"
import "fmt"
import "slices"
import "os"
import "strings"
import "regexp"
import "strconv"

func numbers(s string) []int {
	var re = regexp.MustCompile("\\d+")
	var result []int
	for _, ss := range re.FindAllString(s, -1) {
		n, _ := strconv.ParseInt(ss, 10, 32)
		result = append(result, int(n))
	}
	return result
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var orders [][]int
	var sums [2]int
	for scanner.Scan() {
		var line = scanner.Text()
		if strings.Contains(line, "|") {
			orders = append(orders, numbers(line))
		} else {
			var nns = numbers(line)
			if len(nns) == 0 { continue } // skip empty line
			var cmp = func (left, right int) int {
				for _, order := range orders {
					if order[0] == right && order[1] == left { return 1 }
					if order[0] == left && order[1] == right { return -1 }
				}
				return 0
			}
			if slices.IsSortedFunc(nns, cmp) {
				sums[0] += nns[len(nns) / 2]
			} else {
				slices.SortFunc(nns, cmp)
				sums[1] += nns[len(nns) / 2]
			}
		}
	}
	fmt.Println(sums)
}
