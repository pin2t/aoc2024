package main

import "bufio"
import "fmt"
import "slices"
import "os"
import "strings"
import "regexp"
import "strconv"

func main() {
	type key struct {left, right int}
	var less = make(map[key]int)
	var compare = func (left, right int) int {
		if _, found := less[key{left, right}]; found { return -1 }
		if _, found := less[key{right, left}]; found { return 1 }
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
			less[key{numbers[0], numbers[1]}] = -1
		} else {
			if len(numbers) == 0 { continue } // skip empty line
			if slices.IsSortedFunc(numbers, compare) {
				sums[0] += numbers[len(numbers) / 2]
			} else {
				slices.SortFunc(numbers, compare)
				sums[1] += numbers[len(numbers) / 2]
			}
		}
	}
	fmt.Println(sums)
}
