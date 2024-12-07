package main

import "bufio"
import "os"
import "regexp"
import "strconv"
import "fmt"

func match(numbers []int64, i int, c int64) bool {
	if i == len(numbers) - 1 {
		return numbers[0] == c + numbers[i] || numbers[0] == c * numbers[i]
	}
	return match(numbers, i + 1, c + numbers[i]) || match(numbers, i + 1, c * numbers[i])
}

func match3(numbers []int64, i int, c int64) bool {
	var concat = strconv.FormatInt(c, 10) + strconv.FormatInt(numbers[i], 10)
	var nc, _ = strconv.ParseInt(concat, 10, 64)
	if i == len(numbers) - 1 {
		return numbers[0] == c + numbers[i] || numbers[0] == c * numbers[i] || numbers[0] == nc
	}
	return match3(numbers, i + 1, c + numbers[i]) || match3(numbers, i + 1, c * numbers[i]) || match3(numbers, i + 1, nc)
}

func main() {
	var re = regexp.MustCompile("\\d+")
	var calibrations [2]int64
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var numbers []int64
		for _, s := range re.FindAllString(scanner.Text(), -1) {
			n, _ := strconv.ParseInt(s, 10, 64)
			numbers = append(numbers, n)
		}
		if match(numbers, 0, 0) {
			calibrations[0] += numbers[0]
		}
		if match3(numbers, 0, 0) {
			calibrations[1] += numbers[0]
		}
	}
	fmt.Println(calibrations)
}
