package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// H1
// The number of misplaced tiles; the number of squares that are not in the right place.
// The space isnot a tile, so it cannot be out of place.
func h1(a, b [][]int) int {
	wrong := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] == 0 {
				continue
			}

			if a[i][j] != b[i][j] {
				wrong++
			}
		}
	}
	return wrong
}

// H2
// The Manhattan distance.
func h2(a, b [][]int) int {
	distance := 0
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if a[r][c] == 0 {
				continue
			}
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if a[r][c] == b[i][j] {
						distance += int(math.Abs(float64(r)-float64(i)) + math.Abs(float64(c)-float64(j)))
					}
				}
			}
		}
	}
	return distance
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	itr := 10

	target := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	for i := 0; i < itr; i++ {
		input := randomState()

		fmt.Println("=======================================")
		fmt.Println("input", input)
		fmt.Println("target", target)
		fmt.Println("--")

		if !isSolvable(input) {
			fmt.Println("NOT SOLVABLE")
		}

		p := NewPuzzle(input, target, h1)
		start := time.Now()
		p.Solve()
		duration := time.Since(start)

		fmt.Println(fmt.Sprintf("time: %f", duration.Seconds()))
		fmt.Println("=======================================")
		fmt.Println("")
	}

	for i := 0; i < itr; i++ {
		input := randomState()

		fmt.Println("=======================================")
		fmt.Println("input", input)
		fmt.Println("target", target)
		fmt.Println("--")

		if !isSolvable(input) {
			fmt.Println("NOT SOLVABLE")
		}

		p := NewPuzzle(input, target, h2)
		start := time.Now()
		p.Solve()
		duration := time.Since(start)

		fmt.Println(fmt.Sprintf("time: %f", duration.Seconds()))
		fmt.Println("=======================================")
		fmt.Println("")
	}
}

func randomState() [][]int {
	out := make([][]int, 3)
	seen := make(map[int]bool)
	for i := 0; i < 3; i++ {
		out[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			for {
				n := rand.Intn(9)
				if _, ok := seen[n]; ok {
					continue
				}
				out[i][j] = n
				seen[n] = true
				break
			}
		}
	}
	return out
}

// inversion: https://math.stackexchange.com/questions/293527/how-to-check-if-a-8-puzzle-is-solvable
func isSolvable(arr [][]int) bool {
	flattened := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 0 {
				continue
			}
			flattened = append(flattened, arr[i][j])
		}
	}

	invs := 0
	for i := 0; i < (len(flattened) - 1); i++ {
		for j := (i + 1); j < len(flattened); j++ {
			if flattened[i] > flattened[j] {
				invs++
			}
		}
	}

	return invs%2 == 0
}

func clone(arr [][]int) [][]int {
	out := make([][]int, len(arr))
	for i, v := range arr {
		row := make([]int, len(v))
		copy(row, v)
		out[i] = row
	}
	return out
}
