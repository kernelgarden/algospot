package main

import (
	"bufio"
	"fmt"
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

func nextWord() string {
	scanner.Scan()

	s := make([]byte, 0, 40)

	for _, c := range scanner.Bytes() {
		s = append(s, c)
	}

	return string(s[:])
}

func main() {
	scanner.Split(bufio.ScanWords)

	caseNum := nextInt()
	for c := 0; c < caseNum; c++ {
		num, maxSize := nextInt(), nextInt()
		stuffs := make([]Stuff, num, num)
		for n := 0; n < num; n++ {
			name, size, needs := nextWord(), nextInt(), nextInt()
			stuffs[n] = NewStuff(name, size, needs)
		}

		totalNeeds, packedList :=  getMaxScore(stuffs, maxSize)
		fmt.Fprintf(out, "%v %v\n", totalNeeds, len(packedList))
		for _, s := range packedList {
			fmt.Fprintln(out, s.name)
		}
	}

	out.Flush()
}

type Stuff struct {
	name	string
	size 	int
	needs	int
}

func NewStuff(name string, size, needs int) Stuff {
	return Stuff{name: name, size: size, needs: needs}
}

func getMaxScore(stuffs []Stuff, maxSize int) (totalNeeds int, packList []Stuff) {
	maxScore := 0

	cache := make([][]int, maxSize + 1)
	for i := 0; i < maxSize + 1; i++ {
		cache[i] = make([]int, len(stuffs))
		for j := 0; j < len(stuffs); j++ {
			cache[i][j] = -1
		}
	}

	maxStartIdx := -1
	for i := 0; i < len(stuffs); i++ {
		if nextScore := choice(stuffs, cache, maxSize, i); nextScore > maxScore {
			maxStartIdx = i
			maxScore = nextScore
		}
	}

	packList = getPackList(stuffs, cache, maxStartIdx, maxSize, maxScore)
	totalNeeds = maxScore
	return
}

func choice(stuffs []Stuff, cache [][]int, leftSize, choiceIdx int) int {
	if leftSize <= 0 {
		return 0
	}

	if stuffs[choiceIdx].size > leftSize {
		return 0
	}

	if cache[leftSize][choiceIdx] != -1 {
		return cache[leftSize][choiceIdx]
	}

	size := stuffs[choiceIdx].size

	maxNeeds := 0

	for i := choiceIdx + 1; i < len(stuffs); i++ {
		if nextNeeds := choice(stuffs, cache, leftSize - size, i); nextNeeds > maxNeeds {
			maxNeeds = nextNeeds
		}
	}

	cache[leftSize][choiceIdx] = stuffs[choiceIdx].needs + maxNeeds
	return cache[leftSize][choiceIdx]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func getPackList(stuffs []Stuff, cache [][]int, startIdx int, startSize int, acc int) []Stuff {
	packList := make([]Stuff, 0, 10)

	if startIdx == -1 {
		return make([]Stuff, 0, 0)
	}

	packList = append(packList, stuffs[startIdx])
	acc -= stuffs[startIdx].needs
	leftSize := startSize - stuffs[startIdx].size

	for i := startIdx + 1; i < len(stuffs); i++ {
		// 현재 누적된 절박도의 총 합과 같다면 선택 가능한 선택지다.
		if cache[leftSize][i] == acc {
			acc -= stuffs[i].needs
			leftSize -= stuffs[i].size
			packList = append(packList, stuffs[i])
		} else {
			// not choice
		}
	}

	return packList
}