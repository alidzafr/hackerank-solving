package main

import "fmt"

func main() {
	a := []int{1, 3, 6}
	b := []int{2, 4, 5}

	tempArr := []int{}

	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			tempArr = append(tempArr, a[i], b[i])
		}
		if a[i] > b[i] {
			tempArr = append(tempArr, b[i], a[i])
		}
	}
	fmt.Println(tempArr)
}
