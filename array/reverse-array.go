package main

import "fmt"

func main() {
	xd := []int{2, 3, 4}

	// for _, value := range xd {
	// 	fmt.Printf("%d", value)
	// }
	tempArr := []int{}
	for i := len(xd) - 1; i >= 0; i-- {
		tempArr = append(tempArr, xd[i])
	}
	fmt.Print(tempArr)

}

// s := []int{5, 4, 3, 2, 1}
// for i := len(s)-1; i >= 0; i-- {
//    fmt.Println(s[i])
// }

// tempArr = list()
//     for i in range(len(a)):
//         tempArr.append(a[len(a)-1-i])
//     print(tempArr)
//     return tempArr
