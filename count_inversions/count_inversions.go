package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MergeSort(arr []int) int {
	tmp := make([]int, len(arr))
	return hiddenMergeSort(arr, tmp, 0, len(arr)-1)
}

func hiddenMergeSort(arr, tmp []int, left, right int) int {
	var invCount int

	if right > left {
		mid := (right + left) / 2
		invCount += hiddenMergeSort(arr, tmp, left, mid)
		invCount += hiddenMergeSort(arr, tmp, mid+1, right)
		invCount += merge(arr, tmp, left, mid, right)
	}
	return invCount
}

func merge(arr, tmp []int, left, mid, right int) int {
	var invCount int
	i, j, k := left, mid+1, left

	for (i <= mid) && (j <= right) {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			k++
			i++
		} else {
			tmp[k] = arr[j]
			invCount += (mid - i) + 1
			k++
			j++
		}
	}
	for i <= mid {
		tmp[k] = arr[i]
		k++
		i++
	}

	for j <= right {
		tmp[k] = arr[j]
		k++
		j++
	}

	for i := left; i <= right; i++ {
		arr[i] = tmp[i]
	}

	return invCount
}

func main() {

	in, _ := os.Open("count_inversions/rosalind_inv.txt")
	defer in.Close()

	file := bufio.NewScanner(in)
	file.Buffer(make([]byte, 0, 10000000), 10000000)
	file.Scan()
	file.Scan()
	arr := SplitStringToArray(file.Text())

	num := MergeSort(arr)
	fmt.Println(num)

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
