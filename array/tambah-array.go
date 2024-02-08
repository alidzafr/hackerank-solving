package main

import (
	"fmt"
)

func read(n int) ([]int, error) {
	in := make([]int, n)
	for i := range in {
		_, err := fmt.Scan(&in[i])
		if err != nil {
			return in[:i], err
		}
	}

	return in, nil
}

func main() {
	var n int
	fmt.Scan(&n)

	x, err := read(n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", x)

	// var (
	// len int32
	// integ []int32
	// )
	// integ := make([]int, n)

	// fmt.Scan(&len)
	// fmt.Scanln(&integ)

	// fmt.Printf("lines of text: %d\n", len)
	// arr := integ[1:]
	// fmt.Println(integ)

	// fmt.Println("input text:")
	// var line1, line2, line3 int32
	// _, err := fmt.Scan(&line1, &line2, &line3)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("3 lines of text: %s %s %s\n", line1, line2, line3)
}
