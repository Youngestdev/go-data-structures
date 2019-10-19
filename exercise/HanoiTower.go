package main

import "fmt"

type HanoiState struct {
	src, dst, aux string
	n int
}

func (t *HanoiState) MoveDisk(src, dst string) {
	fmt.Println("Moved disk from ", src, "to ", dst)
}

func (t *HanoiState) MoveTower(src string, dst string, aux string, n int) {
	if n == 1{
		t.MoveDisk(src, dst)
	} else {
		t.MoveTower(src, aux, dst, n-1)
		t.MoveDisk(src, dst)
		t.MoveTower(aux, dst, src, n-1)
	}
	fmt.Println("Move disk ", n, " from tower", src, " to tower", dst )
}

func main() {
	a := HanoiState{
		src: "A",
		dst: "B",
		aux: "C",
		n:   1,
	}
	a.MoveTower("A","B","C", 4)
}