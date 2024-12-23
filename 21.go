package main

import "bufio"
import "os"
import "fmt"

type pos struct {row, col int}
type dir struct {dr, dc int}
var numPad = map[rune]pos{'7': {0, 0}, '8': {0, 1}, '9': {0, 2}, '4': {1, 0}, '5': {1, 1}, '6': {1, 2}, '1': {2, 0}, '2': {2, 1}, '3': {2, 2}, '0': {3, 1}, 'A': {3, 2}}
var dirPad = map[rune]pos{'^': {0, 1}, 'A': {0, 2}, '<': {1, 0}, 'v': {1, 1}, '>': {1, 2}}
var dirChar = map[dir]rune{{0, 1}: '>', {1, 0}: 'v', {0, -1}: '<', {-1, 0}: '^'}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	var codes []string
	for scanner.Scan() { codes = append(codes, scanner.Text() )}
	var results [2]int
	for _, code := range codes {
		var prefix = int(code[0] - '0') * 100 + int(code[1] - '0') * 10 + int(code[2] - '0')
		movesMap = make(map[moveKey][][]rune)
		cache = make(map[cacheKey]int)
		results[0] += seqLen([]rune(code), 0, 2) * prefix
		movesMap = make(map[moveKey][][]rune)
		cache = make(map[cacheKey]int)
		results[1] += seqLen([]rune(code), 0, 25) * prefix
	}
	fmt.Println(results)
}

type moveKey struct {start, end pos}
type cacheKey struct {code string; depth int}
var movesMap = make(map[moveKey][][]rune)
var cache = make(map[cacheKey]int)

func seqLen(code []rune, indir int, maxindir int) (_len int) {
	var key = cacheKey{string(code), indir}
	if l, found := cache[key]; found { return l }
	var from = pos{3, 2}
	var pad = numPad
	if indir > 0 { from = pos{0, 2};  pad = dirPad }
	for _, c := range code {
		var to = pad[c]
		if indir > 0 {
			var found bool
			var key = moveKey{from, to}
			if sequences, found = movesMap[key]; !found {
				sequences = make([][]rune, 0)
				sequence = make([]rune, 0)
				gen(from, to, pad)
				movesMap[key] = sequences
			}
		} else {
			sequences = make([][]rune, 0)
			sequence = make([]rune, 0)
			gen(from, to, pad)
		}
		var m = 10000000000000
		if indir == maxindir {
			for _, seq := range sequences { if len(seq) < m { m = len(seq) }}
		} else {
			for _, seq := range sequences {
				var l = seqLen(seq, indir + 1, maxindir)
				if l < m { m = l }
			}
		}
		_len += m
		from = to
	}
	cache[key] = _len
	return
}

var sequence []rune
var sequences [][]rune

func gen(from pos, to pos, pad map[rune]pos) {
	var seq = make([]rune, len(sequence))
	copy(seq, sequence)
	if to == from {
		seq = append(seq, 'A')
		sequences = append(sequences, seq)
		return
	}
	var dist = abs(from.row - to.row) + abs(from.col - to.col)
	for _, d := range []dir{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
		var next = pos{from.row + d.dr, from.col + d.dc}
		if abs(next.row - to.row) + abs(next.col - to.col) < dist {
			var found = false
			for _, v := range pad {
				if v == next {
					found = true
					break
				}
			}
			if !found { continue }
			sequence = append(sequence, dirChar[d])
			gen(next, to, pad)
			sequence = sequence[0:len(sequence) - 1]
		}
	}
}

func abs(val int) int {
	if val < 0 { return -val }
	return val
}
