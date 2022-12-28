package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func MergeSort(arr []int) []int {
	return hiddenMergeSort(arr)
}

func hiddenMergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	i := len(arr) / 2
	arr1 := hiddenMergeSort(arr[:i])
	arr2 := hiddenMergeSort(arr[i:])
	return Merge(arr1, arr2)
}

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

	in, _ := os.Open("merge_sort/rosalind_ms.txt")
	defer in.Close()
	out, _ := os.Create("merge_sort/answer.txt")
	defer out.Close()

	file := bufio.NewScanner(in)
	file.Buffer(make([]byte, 0, 10000000), 10000000)
	file.Scan()
	file.Scan()
	arr := SplitStringToArray(file.Text())
	mergedArray := MergeSort(arr)
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
