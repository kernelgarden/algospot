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

var totalProb float64

func main() {
	scanner.Split(bufio.ScanWords)

	caseNum := nextInt()
	for c := 0; c < caseNum; c++ {
		depth, days := nextInt(), nextInt()
		cache := make([][]float64, days + 1)
		for i := 0; i <= days; i++ {
			// 최고로 올라가 라 수 있는 거리를 고려한다.
			cache[i] = make([]float64, depth * 2 + 1)
			for j := 0; j < depth * 2 + 1; j++ {
				cache[i][j] = -1
			}
		}

		totalProb = float64(pow(2, days))
		prob := snail(0, days, depth, cache)
		fmt.Fprintf(out, "%-12.10f\n", prob)
	}

	out.Flush()
}

func snail(depth, days, target int, cache [][]float64) float64 {
	// 깊이가 0이 되지 않았는데 날이 다 지난 경우는 빼야함
	if days == 0  {
		if depth >= target {
			//fmt.Printf("[DEBUG] => SUCCSS depth: %v, days %v\n", depth, days)
			return 1
		}

		//fmt.Printf("[DEBUG] => FAIL depth: %v, days %v\n", depth, days)
		return 0
	}

	if cache[days][depth] != -1 {
		return cache[days][depth]
	}

	var prob float64
	prob = 0.75 * snail(depth + 2, days - 1, target, cache) + 0.25 * snail(depth + 1, days - 1, target, cache)
	//fmt.Printf("prob: %v, depth: %v, days %v\n", prob, depth, days)
	cache[days][depth] = prob
	return cache[days][depth]
}

func pow(base, num int) int {
	ret := base
	for i := 1; i < num; i++ {
		ret *= base
	}

	return ret
}
