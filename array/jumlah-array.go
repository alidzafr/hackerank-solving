package main

import "fmt"

func jawaban() int32 {
	var n int32
	fmt.Scan(&n)

	for i := range ar {
		fmt.Scan(&ar[i])
	}

	var sum int32
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}

	return sum
}

func sumArray() {
	arr := []int{2, 3, 5}
	tempArr := []int{}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	tempArr = append(tempArr, sum)
	fmt.Println(tempArr)
}

func simpleArraySum(ar []int32) int32 {
	for i := range ar {
		_, err := fmt.Scan(&ar[i])
		if err != nil {
			panic(err)
		}
	}

	var sum int32
	for i := 0; i < len(ar); i++ {
		sum += ar[i]
	}

	return sum, nil
}

func main() {
	var n int
	fmt.scan(&n)

	x, err := simpleArraySum(n)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", x)
}
