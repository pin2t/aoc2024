package main

import "bufio"
import "os"
import "regexp"
import "strconv"
import "fmt"

var reNumber = regexp.MustCompile("\\d+")

func safe(report []int) bool {
	var increasing = report[0] < report[len(report) - 1]
	for i := 1; i < len(report); i++ {
		if increasing && (report[i] <= report[i - 1] || report[i] - report[i - 1] > 3) ||
		   !increasing && (report[i] >= report[i - 1] || report[i - 1] - report[i] > 3)	{
			return false
		}
	}
	return true
}

func main() {
	var safes [2]int
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var report []int
		for _, s := range reNumber.FindAllString(scanner.Text(), -1) {
			n, _ := strconv.ParseInt(s, 10, 32)
			report = append(report, int(n))
		}
		if safe(report) {
			safes[0]++
			safes[1]++
		} else {
			for i, _ := range report {
				var r = make([]int, len(report))
				copy(r, report)
				if safe(append(r[0:i], r[i+1:]...)) {
					safes[1]++
					break
				}
			}
		}
	}
	fmt.Println(safes)
}
