package main

import "bufio"
import "os"
import "regexp"
import "strconv"
import "fmt"

type op func (calibration, item int64) int64

var plus op = func (calibration, item int64) int64 { return calibration + item }
var mul op = func (calibration, item int64) int64 { return calibration * item }
var concat op = func (calibration, item int64) int64 {
	var result, _ = strconv.ParseInt(strconv.FormatInt(calibration, 10) + strconv.FormatInt(item, 10), 10, 64)
	return result
}

func match(ns []int64, i int, c int64, ops []op) bool {
	if i == len(ns) - 1 {
		return ns[0] == ops[0](c, ns[i]) ||
			   ns[0] == ops[1](c, ns[i]) ||
			   (len(ops) == 3 && ns[0] == ops[2](c, ns[i]))
	}
	return match(ns, i + 1, ops[0](c, ns[i]), ops) ||
		   match(ns, i + 1, ops[1](c, ns[i]), ops) ||
		   (len(ops) == 3 && match(ns, i + 1, ops[2](c, ns[i]), ops))
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
		if match(numbers, 1, 0, []op{plus, mul}) {
			calibrations[0] += numbers[0]
		}
		if match(numbers, 1, 0, []op{plus, mul, concat}) {
			calibrations[1] += numbers[0]
		}
	}
	fmt.Println(calibrations)
}
