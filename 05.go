package main

import "bufio"
import "fmt"
import "slices"
import "os"
import "strings"
import "regexp"
import "strconv"

func main() {
	var orders [][]int
	var cmp = func (left, right int) int {
		for _, order := range orders {
			if order[0] == right && order[1] == left { return 1 }
			if order[0] == left && order[1] == right { return -1 }
		}
		return 0
	}
	var re = regexp.MustCompile("\\d+")
	var scanner = bufio.NewScanner(os.Stdin)
	var sums [2]int
	for scanner.Scan() {
		var line = scanner.Text()
		var numbers []int
		for _, ss := range re.FindAllString(line, -1) {
			n, _ := strconv.ParseInt(ss, 10, 32)
			numbers = append(numbers, int(n))
		}
		if strings.Contains(line, "|") {
			orders = append(orders, numbers)
		} else {
			if len(numbers) == 0 { continue } // skip empty line
			if slices.IsSortedFunc(numbers, cmp) {
				sums[0] += numbers[len(numbers) / 2]
			} else {
				slices.SortFunc(numbers, cmp)
				sums[1] += numbers[len(numbers) / 2]
			}
		}
	}
	fmt.Println(sums)
}
