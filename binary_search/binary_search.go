package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func BinarySearch(keys []int, target []int) (res []string) {
	for _, key := range keys {
		val := bSearch(target, key, 0, len(target)-1)
		if val > 0 {
			val++
		}
		s := strconv.Itoa(val)
		res = append(res, s)
	}
	return res
}

func bSearch(arr []int, key, left, right int) int {
	if right >= left {

		mid := left + (right-left)/2
		switch {
		case arr[mid] == key:
			return mid
		case arr[mid] < key:
			return bSearch(arr, key, mid+1, right)
		case arr[mid] > key:
			return bSearch(arr, key, left, mid-1)
		}
	}
	return -1
}

func main() {
	var arrs [2][]int
	in, _ := os.Open("binary_search/rosalind_bins.txt")
	defer in.Close()
	out, _ := os.Create("binary_search/answer.txt")
	defer out.Close()

	file := bufio.NewScanner(in)
	for i := 0; i < 4; i++ {
		file.Scan()
		if i < 2 {
			continue
		} else {
			s := file.Text()
			arr := SplitStringToArray(s)
			arrs[i-2] = arr
		}
	}

	resArray := BinarySearch(arrs[1], arrs[0])
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
