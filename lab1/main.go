package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// input := [][]int{
	// 	{1, 7, 3},
	// 	{4, 5, 6},
	// 	{2, 0, 8},
	// }

	// // target := [][]int{
	// // 	{1, 3, 6},
	// // 	{4, 5, 0},
	// // 	{2, 7, 8},
	// // }

	// target := [][]int{
	// 	{5, 1, 3},
	// 	{4, 0, 6},
	// 	{2, 7, 8},
	// }

	// target := [][]int{
	// 	{1, 2, 3},
	// 	{0, 4, 6},
	// 	{7, 5, 8},
	// }

	// p, err := NewPuzzle(input, target)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(p)

	// myNode := Node{
	// 	size: 3,
	// 	state: [][]int{
	// 		{5, 1, 3},
	// 		{4, 0, 6},
	// 		{2, 7, 8},
	// 	},
	// }

	//fmt.Println(myNode.move())

	//fmt.Println("test:", randomState())

	itr := 10
	for i := 0; i < itr; i++ {
		input := randomState()
		target := randomState()

		fmt.Println("=======================================")
		fmt.Println("input", input)
		fmt.Println("target", target)
		fmt.Println("--")

		if isSolvable(input) != isSolvable(target) {
			fmt.Println("NOT SOLVABLE")
			//continue
		}

		p := NewPuzzle(input, target, 3)
		start := time.Now()
		p.Solve()
		duration := time.Since(start)

		fmt.Println(fmt.Sprintf("time: %f", duration.Seconds()))
		fmt.Println("=======================================")
		fmt.Println("")
	}

	//fmt.Println(p.seen)

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
	flattend := []int{}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] == 0 {
				continue
			}
			flattend = append(flattend, arr[i][j])
		}
	}

	invs := 0
	for i := 0; i < (len(flattend) - 1); i++ {
		for j := (i + 1); j < len(flattend); j++ {
			if flattend[i] > flattend[j] {
				invs++
			}
		}
	}

	return invs%2 == 0
}
