package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func TwoSum(arr []int) (res []int) {
	var previous = make(map[int]int)
	for idx, val := range arr {
		need := -val
		if prev, ok := previous[need]; ok {
			return append(res, prev+1, idx+1)
		}
		previous[val] = idx
	}
	return append(res, -1)
}

func main() {
	var resArray []string

	in, _ := os.Open("2_sum/rosalind_2sum.txt")
	defer in.Close()
	out, _ := os.Create("2_sum/answer.txt")
	defer out.Close()

	file := bufio.NewScanner(in)
	file.Scan()
	args := SplitStringToArray(file.Text())

	for i := 0; i < args[0]; i++ {
		file.Scan()
		arr := SplitStringToArray(file.Text())
		sum := TwoSum(arr)
		s1 := strconv.Itoa(sum[0])
		if len(sum) > 1 {
			s2 := strconv.Itoa(sum[1])
			resArray = append(resArray, s1+" "+s2)
		} else {
			resArray = append(resArray, s1)
		}
	}

	result := strings.Join(resArray, "\n")
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
