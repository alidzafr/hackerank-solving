package main

import "fmt"

func main() {
	slice_a := []int{1, 2, 3, 4}
	slice_b := []int{3, 2, 1, 2}
	var poin int

	score_a := []int{}
	for i := 0; i < len(slice_a); i++ {
		if slice_a[i] > slice_b[i] {
			poin = 1
		} else {
			poin = 0
		}
		score_a = append(score_a, poin)
	}

	score_b := []int{}
	for i := 0; i < len(slice_b); i++ {
		if slice_b[i] > slice_a[i] {
			poin = 1
		} else {
			poin = 0
		}
		score_b = append(score_b, poin)
	}

	var alice int
	for i := 0; i < len(score_a); i++ {
		alice += score_a[i]
	}

	var bob int
	for i := 0; i < len(score_b); i++ {
		bob += score_b[i]
	}
	score := []int{}
	score = append(score, alice, bob)

	fmt.Println(score)
}
