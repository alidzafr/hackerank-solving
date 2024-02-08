package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func enqueue(i []string) {
	result := []string{}

	if i[0] == "1" {
		result = append(result, i[1])
	} else if i[0] == "2" {
		result[0] = ""
	} else if i[0] == "3" {
		fmt.Println(result)
	}
}

func main() {
	// Total Queue
	scanner := bufio.NewScanner(os.Stdin)
	q := 0
	fmt.Scan(&q)

	tempArr := []int{}
	for i := 0; i < q; i++ {
		scanner.Scan()
		input := scanner.Text()

		// Split
		e := strings.Split(input, " ")

		// 1. Enque
		// 2. Deque front element [:1]
		// 3. Print front element
		if e[0] == "1" {
			v, _ := strconv.Atoi(e[1])
			tempArr = append(tempArr, v)
			fmt.Println(tempArr)
		} else if e[0] == "2" {
			if tempArr != nil {
				tempArr = append(tempArr[:0], tempArr[1:]...) //First in First Out
			} else {
				tempArr = []int{}
			}
			fmt.Println("after delete ", tempArr)
		} else if e[0] == "3" {
			fmt.Println(tempArr[0])
		}
	}
}
