package main

import (
	"fmt"
	"strconv"
	"strings"
)

func prefixSum() {
	n := int32(10)
	queries := [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	} // Range 1 sampai 5 (Q[1,5]) ditambahkan 3

	tempAarr := make([]int64, n+1)
	for _, q := range queries {
		a := q[0]
		b := q[1]
		k := int64(q[2])

		// Rumus Prefix sum A[i,j] = A[j] - A[i-1]
		tempAarr[a-1] += k
		tempAarr[b] -= k
	}

	max := int64(0)
	current_sum := int64(0)
	for _, value := range tempAarr {
		current_sum += value
		// fmt.Println(current_sum)
		if current_sum > max {
			max = current_sum
		}
	}
	fmt.Println(max)

}

func sparceArray() {
	stringList := []string{"def", "de", "fgh", "de"}
	queries := []string{"de", "lmn", "fgh"}

	dummyQueries := []int32{}
	for q := 0; q < len(queries); q++ {
		dummyQueries = append(dummyQueries, int32(q))
		dummyQueries[q] = 0
	}

	for i := 0; i < len(stringList); i++ {
		for j := 0; j < len(queries); j++ {
			if stringList[i] == queries[j] {
				dummyQueries[j] += 1
			}
		}
	}
	fmt.Println(dummyQueries)

}

func dynamicArray() {
	queries := [][]int{
		{1, 0, 5},
		{1, 1, 7},
		{1, 0, 3},
		{2, 1, 0},
		{2, 1, 1},
	}

	tempArr := [2][]int{}
	returnSeq := []int{}
	lastAnswer := 0
	n := 2

	for i := 0; i < len(queries); i++ {
		q := queries[i]
		tipe := q[0]
		x := q[1]
		y := q[2]

		sequenceNum := (x ^ lastAnswer) % n

		if tipe == 1 {
			tempArr[sequenceNum] = append(tempArr[sequenceNum], y)
		} else if tipe == 2 {
			ind_x := y % len(tempArr[sequenceNum])
			lastAnswer = tempArr[sequenceNum][ind_x]
			returnSeq = append(returnSeq, lastAnswer)
		}
		fmt.Println(tempArr)
	}

	// for i := 0; i < len(queries); i++ {
	// 	q := queries[i]
	// 	data := (q[1] ^ lastAnswer) % n

	// 	if q[0] == 1 {
	// 		col = append(col, data)
	// 		fmt.Println(data)
	// 		col = append(col, q[2])
	// 		fmt.Println(col)
	// 	} else if q[0] == 2 {
	// 		x := q[2] % 1
	// 		lastAnswer = col[x]

	// 		res = append(res, lastAnswer)
	// 		fmt.Println(x)
	// 	}
	// }
	// fmt.Println(len(col), lastAnswer, res)

}
func leftRotation() {
	// n := 5
	d := 4
	arr := []int{1, 2, 3, 4, 5}

	for i := 0; i < d; i++ {
		tmp := arr[0]
		arr = append(arr[:0], arr[1:]...)
		arr = append(arr, tmp)
	}
	fmt.Println(arr)
}

func kangarooJump() {
	// Write your code here
	x1 := 0
	v1 := 3
	x2 := 4
	v2 := 2
	kangaroo1 := []int{x1}
	kangaroo2 := []int{x2}
	matchingLoc := []int{}
	answer := false

	for i := 1; i < 10; i++ {
		kangaroo1 = append(kangaroo1, x1+v1*i)
	}
	for i := 1; i < 10; i++ {
		kangaroo2 = append(kangaroo2, x2+v2*i)
	}
	for i := 0; i < len(kangaroo1); i++ {
		if kangaroo1[i] == kangaroo2[i] {
			matchingLoc = append(matchingLoc, kangaroo2[i])
			answer = true
		}
	}
	fmt.Println(kangaroo1, kangaroo2)
	fmt.Println(matchingLoc, answer)
}

func appleandOrange() {
	s := 7
	t := 10

	a := 4  //apple tree
	b := 12 //orange tree

	apples := []int{2, 3, -4}
	oranges := []int{3, -2, -4}
	redA := []int{}
	redO := []int{}

	for i := 0; i < len(apples); i++ {
		apples[i] += a
		if apples[i] >= s && apples[i] <= t {
			redA = append(redA, apples[i])
		}
	}
	for i := 0; i < len(oranges); i++ {
		oranges[i] += b
		if oranges[i] >= s && oranges[i] <= t {
			redO = append(redO, oranges[i])
		}
	}
	fmt.Printf("%d\n%d", len(redA), len(redO))

}

func dadu() {
	// Write your code here

	// n start posisi
	// m nomor dadu
	// k jumlah rol
	// p probabilitas

}

func timeConversion() {

	s := "12:01:00AM"
	hour := s[0:2]
	minute := s[3:5]
	second := s[6:8]

	if strings.Contains(s, "PM") {
		strHour, _ := strconv.Atoi(hour)
		if strHour != 12 {
			strHour = strHour + 12
		}
		hour = strconv.Itoa(strHour)
	} else if strings.Contains(s, "AM") && hour == "12" {
		hour = "00"
	}

	t := fmt.Sprintf("%s:%s:%s", hour, minute, second)
	fmt.Printf(t)
}

func birthdayCandles() {
	candles := []int32{3, 2, 1, 3}
	// n := len(candles)

	var max, v int32
	l := []int32{}
	for _, i := range candles {
		if int32(i) > max {
			max = int32(i)
		}
	}
	for i := 0; i < len(candles); i++ {
		if candles[i] == max {
			l = append(l, candles[i])
			v = int32(len(l))
		}
	}
	fmt.Println(v)

}

func minMaxSum() {
	arr := []int32{7, 69, 2, 221, 8974}

	tempArr := []int64{}
	var sum, min, max, frontSum, backSum int64

	for i := range arr {
		tempArr = append(tempArr, int64(arr[i]))
		sum += tempArr[i]
	}

	min = tempArr[0]
	max = tempArr[0]
	for _, v := range tempArr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	frontSum = sum - max
	backSum = sum - min

	fmt.Println(frontSum, backSum)
}
