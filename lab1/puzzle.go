package main

import (
	"fmt"
	"sort"
)

type Puzzle struct {
	seen   map[string]bool
	start  Node
	target [][]int
}

func NewPuzzle(input, target [][]int, size int) Puzzle {
	startNode := Node{
		size:   size,
		depth:  0,
		parent: nil,
		state:  clone(input),
		score:  0, // TODO
	}

	seen := make(map[string]bool)
	seen[fmt.Sprint(startNode.state)] = true

	return Puzzle{
		seen:   seen,
		start:  startNode,
		target: clone(target),
	}
}

// TODO: add manhattan distans
// func (p *Puzzle) Diff(item *Node) int {
// 	wrong := 0
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			if item.state[i][j] == 0 {
// 				continue
// 			}

// 			if p.target[i][j] != item.state[i][j] {
// 				wrong++
// 			}
// 		}
// 	}
// 	return wrong
// }

func (p *Puzzle) Diff(item *Node) int {
	wrong := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if item.state[i][j] == 0 {
				continue
			}

			if p.target[i][j] != item.state[i][j] {
				wrong++
			}
		}
	}
	return wrong
}

func (p *Puzzle) Solve() {
	queue := NewQueue()

	queue.Push(p.start)
	p.seen[fmt.Sprint(p.start.state)] = true

	for {
		curr := queue.Pop()
		if curr == nil {
			fmt.Println("NO SOLUTION FOUND")
			break
		}

		for _, v := range curr.move() {
			_, ok := p.seen[fmt.Sprint(v.state)]
			if ok {
				continue
			}

			score := p.Diff(&v)
			if score == 0 {
				fmt.Printf("SOLVED (in %d steps) \n", v.parent.depth+1)
				return
			}

			v.score = score + v.parent.depth + 1
			v.depth = v.parent.depth + 1
			queue.Push(v)
			p.seen[fmt.Sprint(v.state)] = true
		}

	}

}

type Queue struct {
	data []Node
}

func NewQueue() Queue {
	return Queue{
		data: []Node{},
	}
}

func (q *Queue) Push(item Node) {
	q.data = append(q.data, item)
	sort.SliceStable(q.data, func(i, j int) bool {
		return q.data[i].score < q.data[j].score
	})
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Pop() *Node {
	if len(q.data) == 0 {
		return nil
	}

	node := q.data[0]
	q.data = q.data[1:]
	return &node
}

type Node struct {
	size   int
	depth  int
	parent *Node
	state  [][]int
	score  int
}

func (n *Node) move() []Node {
	children := []Node{}
	row := -1
	col := -1

	for r, arr := range n.state {
		for c, val := range arr {
			if val == 0 {
				row = r
				col = c
			}
		}
	}

	if row == -1 || col == -1 {
		return children
	}

	neighbors := [][]int{
		{row - 1, col},
		{row + 1, col},
		{row, col - 1},
		{row, col + 1},
	}

	for _, v := range neighbors {
		if v[0] < 0 || v[0] >= n.size || v[1] < 0 || v[1] >= n.size {
			//fmt.Println("WTF2:", neighbors)
			continue
		}

		newState := clone(n.state)

		tmp := newState[row][col]
		newState[row][col] = newState[v[0]][v[1]]
		newState[v[0]][v[1]] = tmp

		children = append(children, Node{
			size:   n.size,
			parent: n,
			state:  newState,
		})
	}

	return children
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
