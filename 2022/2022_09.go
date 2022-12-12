package y2022

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var ps map[int]map[int]bool

type Vector struct {
	X int
	Y int
}

type Step struct {
	Direction string
	Distance  int
}

func parse(input string) []Step {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	steps := make([]Step, len(lines))
	for i, line := range lines {
		distance, err := strconv.Atoi(line[2:])
		if err != nil {
			panic(err)
		}

		steps[i] = Step{line[0:1], distance}
	}

	return steps
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{v.X - v2.X, v.Y - v2.Y}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y}
}

func appendDistinctPosition(position Vector) {
	if ps[position.X] == nil {
		ps[position.X] = make(map[int]bool)
	}

	ps[position.X][position.Y] = true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func normalize(x int) int {
	// check not needed for our purposes because we only call this with abs(x) > 1
	// optimized for our use case
	//  if x == 0 {
	//  return 0
	//  }

	return x / abs(x)
}

func moveKnots(head Vector, tail Vector) Vector {
	diff := head.Sub(tail)

	if abs(diff.X) > 1 {
		tail.X += normalize(diff.X)
		if abs(diff.Y) > 0 {
			tail.Y += diff.Y / abs(diff.Y)
		}
		return tail
	}

	if abs(diff.Y) > 1 {
		tail.Y += normalize(diff.Y)
		if abs(diff.X) > 0 {
			tail.X += diff.X / abs(diff.X)
		}
		return tail
	}

	return tail
}

func moveHead(rope []Vector, steps int, direction string) {
	dirs := map[string]Vector{
		"U": {0, 1},
		"D": {0, -1},
		"L": {-1, 0},
		"R": {1, 0},
	}

	for s := 0; s < steps; s++ {
		rope[0] = rope[0].Add(dirs[direction])
		for i := 0; i < len(rope)-1; i++ {
			rope[i+1] = moveKnots(rope[i], rope[i+1])
		}
		appendDistinctPosition(rope[len(rope)-1])
	}
}

// ///////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////// VISUALIZATION ///////////////////////////////////
// ///////////////////////////////////////////////////////////////////////////////////
var drawInterval = 100 * time.Millisecond

func visualizeRope(rope []Vector) {
	offset := 5
	tail := rope[len(rope)-1]

	// clear screen
	fmt.Printf("\033[2J")

	// heading line
	fmt.Printf(" X ")
	for x := tail.X - (len(rope) + offset); x < tail.X+len(rope)+offset; x++ {
		fmt.Printf("\033[33m%3d\033[0m", x)
	}
	fmt.Printf("\n")

	// rows
	for y := tail.Y - (len(rope) + offset); y < tail.Y+len(rope)+offset; y++ {
		fmt.Printf("\033[33m%3d\033[0m ", y)
		for x := tail.X - (len(rope) + offset); x < tail.X+len(rope)+offset; x++ {
			fmt.Printf("%s", getRopeChar(rope, x, y))
		}
		fmt.Println()
	}

	// stats
	fmt.Println()
	fmt.Printf("rope length:\t\t %5d\n", len(rope))

	n := 0
	for _, ys := range ps {
		for range ys {
			n++
		}
	}
	fmt.Printf("distinct positions:\t %5d\n", n)
}

func getRopeChar(rope []Vector, x int, y int) string {
	for i, p := range rope {
		if p.X == x && p.Y == y {
			if i == 0 {
				return " X "
			}

			return fmt.Sprintf("%3d", len(rope)-i)
		}
	}

	for px, ys := range ps {
		for py := range ys {
			if px == x && py == y {
				return " \033[31m*\033[0m "
			}
		}
	}

	return " . "
}

// //////////////////////////////////////////////////////////////////////////
// ///////////////////////////////// MAIN ///////////////////////////////////
// //////////////////////////////////////////////////////////////////////////

//var visualize = false

func solution(ropeLength int, steps []Step) int {
	ps = make(map[int]map[int]bool)

	rope := make([]Vector, ropeLength)

	appendDistinctPosition(rope[0])
	for _, step := range steps {
		moveHead(rope, step.Distance, step.Direction)
	}

	n := 0
	for _, ys := range ps {
		for range ys {
			n++
		}
	}

	return n
}

func y2022_9_1(input string) string {
	steps := parse(input)

	return strconv.Itoa(solution(2, steps))
}

func y2022_9_2(input string) string {
	steps := parse(input)

	return strconv.Itoa(solution(10, steps))
}

func Y91(input string) string {
	return y2022_9_1(input)
}

func Y92(input string) string {
	return y2022_9_2(input)
}
