package main

import (
	"fmt"
)

func pagerBerpola() {
	n := 6
	string := ""

	for i := 1; i <= n; i++ {
		for j := 0; j < n-i; j++ {
			string += "*"
		}
		for k := 0; k < i; k++ {
			string += "1"
		}
		string += "\n"
	}
	fmt.Println(string)
}

func plusminus() {
	in := []int32{-4, 3, -9, 0, 4, 1}

	n := len(in)

	neg_slc := []int32{}
	zero_slc := []int32{}
	pos_slc := []int32{}
	for i := 0; i < n; i++ {
		if in[i] > 0 {
			neg_slc = append(neg_slc, in[i])
		}
		if in[i] == 0 {
			zero_slc = append(zero_slc, in[i])
		}
		if in[i] < 0 {
			pos_slc = append(pos_slc, in[i])
		}
	}

	p := float64(len(neg_slc)) / float64(n)
	q := float64(len(zero_slc)) / float64(n)
	r := float64(len(pos_slc)) / float64(n)

	s := fmt.Sprintf("%.6f \n%.6f \n%.6f ", p, r, q)
	fmt.Println(s)
}

func bigSum(ar []int64) int64 {
	var n int64
	fmt.Scan(&n)
	fmt.Scan(&ar[n])

	var sum int64
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}
	return sum

}

func main() {
	// ar := []int64{1000, 1001, 1002, 1003, 1004}
	// bigSum()
	matrix := [][]int{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	priDi := []int{}
	secDi := []int{}

	for i := range matrix {
		for j := range matrix {
			if i == j {
				priDi = append(priDi, matrix[i][j])
			}
			if i+j == len(matrix)-1 {
				secDi = append(secDi, matrix[i][j])
			}
		}
	}

	var a int
	var b int

	for i := 0; i < len(priDi); i++ {
		a += priDi[i]
	}
	for i := 0; i < len(secDi); i++ {
		b += secDi[i]
	}
	diff := a - b
	if diff <= 0 {
		diff *= -1
	}
	fmt.Println(a, b, diff)

}
