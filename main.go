package main

import (
	"fmt"
	"io"
	"math"
	"sort"
)

type Info struct {
	mean float64
	med  float64
	mode int
	sd   float64
}

func CalcMean(nums *[]int) float64 {
	total := 0.0

	for _, v := range *nums {
		total += float64(v)
	}
	return total / float64(len(*nums))
}

func CalcMedian(sortNums []int) float64 {
	x := len(sortNums) / 2

	if len(sortNums)%2 == 0 {
		return float64((sortNums[x-1] + sortNums[x]) / 2)
	} else {
		return float64(x)
	}
}

func CalcMode(nums *[]int) int {
	m := make(map[int]int)
	var n int
	var count int

	for _, v := range *nums {
		m[v]++
	}
	for key, v := range m {
		if count == v {
			if key < n {
				n = key
			}
		} else if count < v {
			count = v
			n = key
		}
	}
	return n
}

func CalcSd(nums []int, info Info) float64 {
	var tmp float64

	for v := range nums {
		tmp += (float64(v) - info.mean) * (float64(v) - info.mean)
	}
	return math.Sqrt(tmp / float64(len(nums)))
}

func ChooseInfo(x int, info Info) {
	switch x {
	case 1:
		fmt.Printf("%.2f\n", info.mean)
	case 2:
		fmt.Printf("%.2f\n", info.med)
	case 3:
		fmt.Printf("%d\n", info.mode)
	case 4:
		fmt.Printf("%.2f\n", info.sd)
	case 5:
		fmt.Printf("%.2f\n", info.mean)
		fmt.Printf("%.2f\n", info.med)
		fmt.Printf("%d\n", info.mode)
		fmt.Printf("%.2f\n", info.med)
	default:
		fmt.Println("Choose parameters:\npress 1 for Mean\npress 2 for Median\npress 3 for Mode\npress 4 for SD\n" +
			"press 5 for all\n------------------------------------")
	}
}

func main() {
	var tmp int
	var err error
	var info Info
	nums := make([]int, 1)

	fmt.Println("Please input numbers:")
	for _, err = fmt.Scan(&tmp); err == nil; _, err = fmt.Scan(&tmp) {
		if tmp > 2147483647 && tmp < -2147483648 {
			fmt.Println("Error with argument", tmp)
			return
		}
		nums = append(nums, tmp)
	}
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
	var sNums = nums[1:]
	sortNums := make([]int, len(sNums))
	copy(sortNums, sNums)
	sort.Ints(sortNums)
	info.mean = CalcMean(&sNums)
	info.med = CalcMedian(sortNums)
	info.mode = CalcMode(&sNums)
	info.sd = CalcSd(sNums, info)
	fmt.Println("Choose parameters:\npress 1 for Mean\npress 2 for Median\npress 3 for Mode\npress 4 for SD\n" +
		"press 5 for all\n------------------------------------")
	for _, err = fmt.Scan(&tmp); err == nil; _, err = fmt.Scan(&tmp) {
		ChooseInfo(tmp, info)
	}
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	}
}
