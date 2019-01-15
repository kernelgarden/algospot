package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func main() {
	scanner.Split(bufio.ScanWords)

	caseNum := nextInt()
	for i := 0; i < caseNum; i++ {
		num, toUse := nextInt(), nextInt()
		nums := make([]int, num)
		for j := 0; j < num; j++ {
			nums[j] = nextInt()
		}

		fmt.Fprintln(out, quantization(nums, toUse))
	}

	out.Flush()
}

func quantization(nums []int, toUse int) int {
	sort.Ints(nums)

	cache := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		cache[i] = make([]int, toUse + 1)
		for j := 0; j <= toUse; j++ {
			cache[i][j] = -1
		}
	}

	ret := divide(nums, 0, toUse, cache)
	return ret
}

// nums 는 정렬된 상태로 가정한다.
func divide(nums[]int, startIdx, toUse int, cache [][]int) int {
	if startIdx == len(nums) {
		return 0
	}

	if toUse <= 0 {
		return math.MaxInt32
	}

	// 캐시의 키를 startIdx만 잡고 하려다 실패했다.. 이러면 안된다. 캐싱을 할때는 모든 매개 변수를 다 키로 사용해야한다.
	if cache[startIdx][toUse] != -1 {
		//fmt.Printf("Cache Hit! startIdx: %v, toUse: %v, cache: %v\n", startIdx, toUse, cache[startIdx][toUse])
		return cache[startIdx][toUse]
	}

	low := math.MaxInt32

	for i := startIdx + 1; i <= len(nums); i++ {
		err := getTotalErr(nums[startIdx : i])
		remain := divide(nums, i, toUse - 1, cache)
		low = min(low, err + remain)
		//low = min(low, getTotalErr(nums[startIdx : i]) + divide(nums, i, toUse - 1, cache))
	}

	cache[startIdx][toUse] = low
	return cache[startIdx][toUse]
}

func round(num float64) float64 {
	intN := int(num)
	if (num - float64(intN)) >= 0.5 {
		return num + 1
	} else {
		return num
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func pow2(num int) int {
	return num * num
}

func avg(nums []int) int {
	sum := float64(0)
	for _, num := range nums {
		sum += float64(num)
	}
	//avg := int(math.Round(sum / float64(len(nums))))
	avg := int(round(sum / float64(len(nums))))
	return avg
}

func getTotalErr(nums []int) int {
	sum := 0
	avg := avg(nums)
	for _, num := range nums {
		sum += pow2(avg - num)
	}
	//fmt.Printf("[DEBUG] sum: %v, avg: %v, len: %v, nums: %v\n", sum, avg, len(nums), nums)
	return sum
}
