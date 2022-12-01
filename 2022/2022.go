package y2022

import (
	"sort"
	"strconv"
	"strings"
)

func y2022_1_1(input string) string {
	split := strings.Split(input, "\n\n")

	var sum int
	var max int

	max = 0
	for _, s := range split {
		sum = 0
		for _, s2 := range strings.Split(s, "\n") {
			s2s, err := strconv.Atoi(s2)
			if err != nil {
				continue
			}
			sum += s2s
		}
		if sum > max {
			max = sum
		}
	}

	return strconv.Itoa(max)
}

func y2022_1_2(input string) string {
	split := strings.Split(input, "\n\n")

	var sum int
	var sums []int

	for _, s := range split {
		sum = 0
		for _, s2 := range strings.Split(s, "\n") {
			s2s, err := strconv.Atoi(s2)
			if err != nil {
				continue
			}
			sum += s2s
		}
		sums = append(sums, sum)
	}

	sort.Ints(sums)

	return strconv.Itoa(sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3])
}

