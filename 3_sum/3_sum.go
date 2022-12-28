package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

//ThreeSum from bytebybyte
func ThreeSum(arr []int) (res []int) {
	index := make(map[int]int, len(arr))
	for idx, val := range arr {
		index[val] = idx
	}
	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		if !(i == 0 || arr[i] > arr[i-1]) {
			continue
		}
		start := i + 1
		end := len(arr) - 1
		for start < end {
			if arr[i]+arr[start]+arr[end] == 0 {
				return append(res, index[arr[i]]+1, index[arr[start]]+1, index[arr[end]]+1)
			}
			if arr[i]+arr[start]+arr[end] < 0 {
				curStart := start
				for arr[start] == arr[curStart] {
					start++
				}
			} else {
				curEnd := end
				for arr[end] == arr[curEnd] {
					end--
				}
			}
		}
	}
	return append(res, -1)
}

func main() {
	var resArray []string

	in, _ := os.Open("3_sum/rosalind_3sum.txt")
	defer in.Close()
	out, _ := os.Create("3_sum/answer.txt")
	defer out.Close()

	file := bufio.NewScanner(in)
	file.Scan()
	args := SplitStringToArray(file.Text())
	for i := 0; i < args[0]; i++ {
		file.Scan()
		arr := SplitStringToArray(file.Text())
		sum := ThreeSum(arr)
		sort.Ints(sum)
		s1 := strconv.Itoa(sum[0])
		if len(sum) > 1 {
			s2 := strconv.Itoa(sum[1])
			s3 := strconv.Itoa(sum[2])
			resArray = append(resArray, s1+" "+s2+" "+s3)
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
