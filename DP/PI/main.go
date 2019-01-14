package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func nextInt() int {
	scanner.Scan()
	r := 0
	f := 1
	for _, c := range scanner.Bytes() {
		if c == '-' {
			f = -1
			continue
		}
		r *= 10
		r += int(c - '0')
	}
	return r * f
}

func nextString() string {
	scanner.Scan()

	s := make([]byte, 0, 10)

	for _, c := range scanner.Bytes() {
		s = append(s, c)
	}

	return string(s[:])
}

func main() {
	scanner.Split(bufio.ScanWords)

	caseNum := nextInt()
	for i := 0; i < caseNum; i++ {
		numString := nextString()
		fmt.Fprintln(out, GetTotalScore(numString))
	}

	out.Flush()
}

func GetTotalScore(num string) int {
	cache := make([]int, len(num))
	for i := 0; i < len(num); i++ {
		cache[i] = math.MaxInt32
	}

	return getScore(num, -1, cache)
}

func getScore(num string, lastIdx int, cache []int) int {
	beginIdx := lastIdx + 1
	if beginIdx == len(num) {
		return 0
	}

	if cache[beginIdx] != math.MaxInt32 {
		return cache[beginIdx]
	}

	for i := 3; i <= 5; i++ {
		if beginIdx + i <= len(num) {
			cache[beginIdx] = min(cache[beginIdx], score(num, beginIdx, beginIdx + i) + getScore(num, lastIdx + i, cache))
		}
	}

	return cache[beginIdx]
}

func score(num string, startIdx, endIdx int) int {
	piece := num[startIdx : endIdx]

	var isArithmeticSeq bool
	var diff int

	if checkAllSame(piece) {
		return 1
	}

	isArithmeticSeq, diff = checkArithmeticSeq(piece)
	diff = abs(diff)
	if isArithmeticSeq && diff == 1 {
		return 2
	}

	if checkInterval(piece) {
		return 4
	}

	if isArithmeticSeq {
		return 5
	}

	return 10
}

func checkAllSame(piece string) bool {
	start := piece[0]
	for i := 1; i < len(piece); i++ {
		if start != piece[i] {
			return false
		}
	}

	return true
}

func checkInterval(piece string) bool {
	for i := 0; i < len(piece); i++ {
		if piece[i] != piece[i % 2] {
			return false
		}
	}

	return true
}

func checkArithmeticSeq(piece string) (bool, int) {
	for i := 0; i < len(piece) - 1; i++ {
		if (piece[i + 1] - piece[i]) != (piece[1] - piece[0]) {
			return false, 0
		}
	}

	return true, int(piece[1] - piece[0])
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	} else {
		return a
	}
}