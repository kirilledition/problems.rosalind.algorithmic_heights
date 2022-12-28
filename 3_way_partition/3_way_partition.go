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
	less, i, j, greater := 0, 0, len(arr)-1, len(arr)-1
	for i <= j {
		for i <= j && arr[i] <= pivot {
			if arr[i] < pivot {
				arr[less], arr[i] = arr[i], arr[less]
				less++
			}
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
		for i <= j && arr[j] >= pivot {
			if arr[j] > pivot {
				arr[greater], arr[j] = arr[j], arr[greater]
				greater--
			}
			j--
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	return i + 1
}

func main() {
	in, _ := os.Open("3_way_partition/rosalind_par3.txt")
	//in, _ := os.Open("3_way_partition/test")
	defer in.Close()
	out, _ := os.Create("3_way_partition/answer.txt")
	defer out.Close()

	file := bufio.NewScanner(in)
	file.Buffer(make([]byte, 0, 10000000), 10000000)
	file.Scan()
	file.Scan()
	array := SplitStringToArray(file.Text())
	fmt.Println(array[0])
	Partition(array)
	//fmt.Println(array)
	result := ArrayToString(array)
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

func ArrayToString(array []int) string {
	var resArray []string
	for _, val := range array {
		resArray = append(resArray, strconv.Itoa(val))
	}
	return strings.Join(resArray, " ")
}
