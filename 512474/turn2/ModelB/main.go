package main

import (
	"fmt"
	"sort"
)

type student struct {
	Name  string
	Score int
}

func (s student) String() string {
	return fmt.Sprintf("%s: %d", s.Name, s.Score)
}

type byScore []student

func (a byScore) Len() int           { return len(a) }
func (a byScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

func main() {
	students := []student{
		{"Alice", 85},
		{"Charlie", 90},
		{"Bob", 75},
	}

	sort.Sort(byScore(students))

	output := fmt.Sprintf("Sorted students by score: \n%s", students)

	fmt.Println(output)
}