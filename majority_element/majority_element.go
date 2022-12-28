package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func MajorityElement(arr []int) int {
	var maxNumber int
	var maxKey int
	m := map[int]int{}
	for _, val := range arr {
		m[val]++
	}
	for key, val := range m {
		if val > maxNumber {
			maxNumber = val
			maxKey = key
		}
	}
	if maxNumber > len(arr)/2 {
		return maxKey
	}
	return -1
}

func main() {
	var resArray []string

	in, _ := os.Open("majority_element/rosalind_maj.txt")
	defer in.Close()
	out, _ := os.Create("majority_element/answer.txt")
	defer out.Close()

	file := bufio.NewScanner(in)
	file.Scan()
	args := SplitStringToArray(file.Text())

	for i := 0; i < args[0]; i++ {
		file.Scan()
		arr := SplitStringToArray(file.Text())
		s := strconv.Itoa(MajorityElement(arr))
		resArray = append(resArray, s)
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
