package main

import (
	"bufio"
	"fmt"
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
		n, m := nextInt(), nextInt()
		nums1 := make([]int, n)
		nums2 := make([]int, m)

		for j := 0; j < n; j++ {
			nums1[j] = nextInt()
		}

		for j := 0; j < m; j++ {
			nums2[j] = nextInt()
		}

		fmt.Fprintln(out, getJLIS(n,m, nums1, nums2))
	}

	out.Flush()
}

func getJLIS(n, m int, nums1, nums2 []int) int {
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, m)
		for j := 0; j < m; j++ {
			cache[i][j] = -1
		}
	}

	longest := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			longest = max(longest, jlis(nums1, nums2, i, j, cache))
		}
	}

	return longest
}

func jlis(nums1, nums2 []int, start1, start2 int, cache [][]int) int {
	if cache[start1][start2] != -1 {
		return cache[start1][start2]
	}

	longest := 2

	prev := nums1[start1]
	n := len(nums1)
	for i := start1 + 1; i < n; i++ {
		if nums1[i] > prev {
			longest = max(longest, jlis(nums1, nums2, i, start2, cache))
			prev = nums1[i]
		}
	}

	m := len(nums2)
	prev = nums2[start2]
	for i := start2 + 1; i < m; i++ {
		if nums2[i] > prev {
			longest = max(longest, jlis(nums1, nums2, start1, i, cache))
			prev = nums2[i]
		}
	}

	cache[start1][start2] = longest
	return cache[start1][start2]
}

func join(nums1, nums2 []int) []int {
	n := len(nums1)
	m := len(nums2)
	ret := make([]int, 0, n + m)

	ret = append(ret, nums1...)
	ret = append(ret, nums2...)

	sort.Ints(ret)

	return ret
}

// 정렬이 되있다고 가정한다.
func removeDuplicated(sortedNums []int) []int {
	ret := make([]int, 0, len(sortedNums))

	if len(sortedNums) > 0 {
		last := sortedNums[0]
		ret = append(ret, last)
		for i := 1; i < len(sortedNums); i++ {
			if sortedNums[i] > last {
				ret = append(ret, sortedNums[i])
				last = sortedNums[i]
			}
		}
	}

	return ret
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}