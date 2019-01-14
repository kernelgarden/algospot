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
		cache[i] = -1
	}

	return getScore(num, -1, cache)
}

func getScore(num string, lastIdx int, cache []int) int {
	beginIdx := lastIdx + 1
	if beginIdx >= len(num) {
		return 0
	}

	if cache[beginIdx] != -1 {
		return cache[beginIdx]
	}

	low := math.MaxInt32
	for i := 3; i <= 5; i++ {
		//if beginIdx + i < len(num) {
			low = min(low, score(num, beginIdx, beginIdx + i) + getScore(num, lastIdx + i, cache))
			/*
			ret := score(num, beginIdx, beginIdx + i) + getScore(num, lastIdx + i, cache)
			//fmt.Printf("begin:%v, i:%v, ret:%v\n", beginIdx, i, ret)
			if ret < low {
				low = ret
				//fmt.Printf("i: %v, beginIdx: %v, low: %v\n", i, lastIdx + 1, low)
			}
			*/
		//}
	}

	cache[beginIdx] = low
	return low
}

func score(num string, startIdx, endIdx int) int {
	if endIdx > len(num) {
		return math.MaxInt32
	}

	piece := num[startIdx : endIdx]

	if checkAllSame(piece) {
		return 1
	} else if flag, diff := checkArithmeticSeq(piece); flag {
		if diff == -1 || diff == 1 {
			return 2
		} else {
			return 5
		}
	} else if checkInterval(piece) {
		return 4
	} else {
		return 10
	}
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
	first, second := piece[0], piece[1]

	for i := 2; i < len(piece); i++ {
		if i % 2 == 0 {
			if piece[i] != first {
				return false
			}
		} else {
			if piece[i] != second {
				return false
			}
		}
	}

	return true
}

func checkArithmeticSeq(piece string) (bool, int) {
	first, second := piece[0], piece[1]
	diff := second - first

	prev := second
	for i := 2; i < len(piece); i++ {
		if prev + diff != piece[i] {
			return false, 0
		} else {
			prev = piece[i]
		}
	}

	return true, int(diff)
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
