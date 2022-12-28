package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Merge(arr1, arr2 []int) []int {
	arr3 := make([]int, len(arr1)+len(arr2))
	var (
		i1 int
		i2 int
	)
	for (i1 + i2) < len(arr3) {
		switch {
		case i1 < len(arr1) && i2 < len(arr2):
			if arr1[i1] < arr2[i2] {
				arr3[i1+i2] = arr1[i1]
				i1++
			} else {
				arr3[i1+i2] = arr2[i2]
				i2++
			}

		case i1 == len(arr1):
			arr3[i1+i2] = arr2[i2]
			i2++

		case i2 == len(arr2):
			arr3[i1+i2] = arr1[i1]
			i1++
		}
	}
	return arr3
}

func main() {
	var resArray []string
	in, _ := os.Open("merge/rosalind_mer.txt")
	defer in.Close()
	out, _ := os.Create("merge/answer.txt")
	defer out.Close()

	file := bufio.NewScanner(in)
	file.Scan()
	file.Scan()
	arr1 := SplitStringToArray(file.Text())
	file.Scan()
	file.Scan()
	arr2 := SplitStringToArray(file.Text())

	mergedArray := Merge(arr1, arr2)

	for _, val := range mergedArray {
		resArray = append(resArray, strconv.Itoa(val))
	}

	result := strings.Join(resArray, " ")
	out.WriteString(result)
}

func SplitStringToArray(s string) []int {
	var res []int
	stringArray := strings.Split(s, " ")

	for _, elem := range stringArray {
		i, _ := strconv.Atoi(elem)
		res = append(res, i)
	}
	return res
}
