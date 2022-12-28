package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InsertionSort(arr []int) int {
	var counter int
	for i := 1; i < len(arr); i++ {
		k := i
		for k > 0 && arr[k] < arr[k-1] {
			arr[k-1], arr[k] = arr[k], arr[k-1]
			k = k - 1
			counter++
		}
	}
	return counter
}

func main() {
	in, _ := os.Open("insertion_sort/rosalind_ins.txt")
	defer in.Close()

	file := bufio.NewScanner(in)
	file.Scan()
	file.Scan()
	array := SplitStringToArray(file.Text())

	b := InsertionSort(array)
	fmt.Println(b)
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
