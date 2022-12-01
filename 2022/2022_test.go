package y2022

import (
	"testing"

	"github.com/dithmer/aoc-go/util"
)

func TestY2022_1_1(t *testing.T) {
	util.GetTest(2022, 1, 1, []util.TestCase{
		{Input: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`, Expected: "24000"},
	}, y2022_1_1)(t)
}

func TestY2022_1_2(t *testing.T) {
	util.GetTest(2022, 1, 2, []util.TestCase{
		{Input: `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`, Expected: "45000"},
	}, y2022_1_2)(t)
}

