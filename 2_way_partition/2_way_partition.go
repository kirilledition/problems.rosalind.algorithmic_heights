package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Partition(arr []int) int {
	pivot := arr[0]
	i, j := 1, len(arr)-1
	for i < j {
		for arr[i] <= pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[j] = arr[j], arr[i]

	arr[0], arr[j] = arr[j], arr[0]
	return i + 1
}

func main() {
	var resArray []string
	in, _ := os.Open("2_way_partition/rosalind_par.txt")
	defer in.Close()
	out, _ := os.Create("2_way_partition/answer.txt")
	defer out.Close()

	file := bufio.NewScanner(in)
	file.Buffer(make([]byte, 0, 10000000), 10000000)
	file.Scan()
	file.Scan()
	array := SplitStringToArray(file.Text())
	fmt.Println(array)
	i := Partition(array)
	fmt.Println(array, "i", i)
	for _, val := range array {
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
