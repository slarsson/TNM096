package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"text/tabwriter"
	"time"
	"unicode"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	classes := []string{
		"MT101",
		"MT102",
		"MT103",
		"MT104",
		"MT105",
		"MT106",
		"MT107",

		"MT201",
		"MT202",
		"MT203",
		"MT204",
		"MT205",
		"MT206",

		"MT301",
		"MT302",
		"MT303",
		"MT304",

		"MT401",
		"MT402",
		"MT403",

		"MT501",
		"MT502",
	}

	classrooms := []string{
		"TP51",
		"SP34",
		"K3",
	}

	s, err := schedule(classes, classrooms)
	if err != nil {
		fmt.Println(err)
		return
	}

	s.Print()
}

var allowedHours []int = []int{9, 10, 11, 12, 1, 2, 3, 4}

type Schedule struct {
	rooms  []string
	matrix []string
}

func schedule(classes []string, classrooms []string) (*Schedule, error) {
	if len(classes) == 0 || len(classrooms) == 0 {
		return nil, fmt.Errorf("missing classes and/or classrooms")
	}

	size := len(classrooms) * len(allowedHours)

	if len(classes) > size {
		return nil, fmt.Errorf("to many classes")
	}

	s := Schedule{
		rooms:  []string{},
		matrix: make([]string, size),
	}

	for _, room := range classrooms {
		s.rooms = append(s.rooms, room)
	}

	for i := 0; i < len(classes); i++ {
		s.matrix[i] = classes[i]
	}

	// create random schedule
	for i := 0; i < (10 * size); i++ {
		a := rand.Intn(size)
		b := rand.Intn(size)
		s.matrix[a], s.matrix[b] = s.matrix[b], s.matrix[a]
	}

	err := s.solve()
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (s Schedule) solve() error {
	maxLimit := 100000
	for i := 0; i < maxLimit; i++ {
		conflicts := s.conflicts()
		n := len(conflicts)
		if n == 0 {
			fmt.Printf("Solved in %d steps \n", i)
			return nil
		}

		// get random conflicting item
		index := rand.Intn(n)

		s.swap(conflicts[index])
	}

	return fmt.Errorf("no solution found")
}

// find all conflicting items
func (s Schedule) conflicts() map[int]int {
	found := make(map[int]int)
	w := len(s.rooms)
	for i := 0; i < len(s.matrix); i += w {
		for j := i; j <= (i + w - 1); j++ {
			for k := j + 1; k < (i + w); k++ {
				if (s.matrix[j] == "MT501" || s.matrix[j] == "MT502") && (s.matrix[k] == "MT501" || s.matrix[k] == "MT502") {
					continue
				}

				d1, err := firstDigit(s.matrix[j])
				if err != nil {
					continue
				}

				d2, err := firstDigit(s.matrix[k])
				if err != nil {
					continue
				}

				if d1 == d2 {
					found[j] = j
				}
			}
		}
	}

	return found
}

func (s Schedule) swap(index int) {
	items := [][2]int{}

	// get the number of conflicts when swapping position with 'index'
	for i := 0; i < len(s.matrix); i++ {
		if i == index {
			continue
		}

		s.matrix[index], s.matrix[i] = s.matrix[i], s.matrix[index]
		items = append(items, [2]int{i, len(s.conflicts())})
		s.matrix[index], s.matrix[i] = s.matrix[i], s.matrix[index]
	}

	if len(items) == 0 {
		return
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i][1] < items[j][1]
	})

	// count the min-conflicts items
	size := 0
	for i := 0; i < len(items); i++ {
		size++
		if items[i][1] > items[0][1] {
			break
		}

	}

	// select random item, if there are more than one
	minIndex := items[0][0]
	if size > 1 {
		n := rand.Intn(size)
		minIndex = items[n][0]
	}

	s.matrix[index], s.matrix[minIndex] = s.matrix[minIndex], s.matrix[index]
}

func (s Schedule) Print() {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)

	h1 := ""
	h2 := ""
	for _, room := range s.rooms {
		h1 += "\t" + room
		h2 += "\t-----"
	}
	fmt.Fprintln(writer, h1)
	fmt.Fprintln(writer, h2)

	w := len(s.rooms)
	idx := 0
	for i := 0; i < len(s.matrix); i += w {
		out := fmt.Sprintf("%d", allowedHours[idx])
		idx++
		for j := i; j < (i + w); j++ {
			out += "\t" + s.matrix[j]
		}
		fmt.Fprintln(writer, out)
	}

	writer.Flush()
}

func firstDigit(str string) (int, error) {
	d := -1
	for _, v := range str {
		if unicode.IsDigit(v) {
			d = int(v - '0')
			break
		}
	}

	if d == -1 {
		return -1, fmt.Errorf("no digit found")
	}
	return d, nil
}
