package main

import (
	"container/heap"
	"fmt"
)

type Heuristics func([][]int, [][]int) int

type Puzzle struct {
	seen    map[string]bool
	start   Node
	target  [][]int
	handler Heuristics
}

func NewPuzzle(input, target [][]int, action Heuristics) Puzzle {
	startNode := Node{
		parent: nil,
		state:  clone(input),
		depth:  0,
		score:  -1,
	}

	seen := make(map[string]bool)
	seen[fmt.Sprint(startNode.state)] = true

	return Puzzle{
		seen:    seen,
		start:   startNode,
		target:  clone(target),
		handler: action,
	}
}

func (p *Puzzle) Solve() {
	p.seen[fmt.Sprint(p.start.state)] = true
	queue := PriorityQueue{p.start}
	heap.Init(&queue)

	for {
		if queue.Len() == 0 {
			fmt.Println("NO SOLUTION")
			return
		}

		curr := heap.Pop(&queue).(Node)
		for _, v := range curr.Eval() {
			_, ok := p.seen[fmt.Sprint(v.state)]
			if ok {
				continue
			}

			score := p.handler(v.state, p.target)
			if score == 0 {
				fmt.Println(fmt.Sprintf("SOLVED (in %d steps)", v.parent.depth+1))
				return
			}

			v.score = v.depth + score

			heap.Push(&queue, v)
			p.seen[fmt.Sprint(v.state)] = true
		}
	}
}

type Node struct {
	parent *Node
	state  [][]int
	depth  int
	score  int
}

func (n *Node) Eval() []Node {
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
		if v[0] < 0 || v[0] >= 3 || v[1] < 0 || v[1] >= 3 {
			continue
		}

		newState := clone(n.state)

		tmp := newState[row][col]
		newState[row][col] = newState[v[0]][v[1]]
		newState[v[0]][v[1]] = tmp

		children = append(children, Node{
			parent: n,
			state:  newState,
			depth:  n.depth + 1,
			score:  0,
		})
	}

	return children
}
