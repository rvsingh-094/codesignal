package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isLucky(n int) bool {
	strNum := strconv.Itoa(n)
	strLst := strings.Split(strNum, "")
	if len(strLst)%2 != 0 {
		return false
	}
	firstHalf := strLst[:len(strLst)/2]
	secondHalf := strLst[len(strLst)/2:]
	firstSum := 0
	secondSum := 0
	for _, val := range firstHalf {
		num, _ := strconv.Atoi(val)
		firstSum += num
	}
	for _, val := range secondHalf {
		num, _ := strconv.Atoi(val)
		secondSum += num
	}
	if firstSum == secondSum {
		return true
	}
	fmt.Println("AfterSplit: ", firstHalf, secondHalf)
	return false
}

func main() {
	isLucky(123456)
}

