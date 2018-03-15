package main

import (
	"os"
	"bufio"
	"strconv"
	"fmt"
	"strings"
)

func main() {
	file, err := os.Open("B-large-practice.in")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewReader(file)

	lines, _, err := scanner.ReadLine()
	ls, err := strconv.Atoi(string(lines))
	if err != nil {
		// handle error
		panic(err)
	}
	for x := 1; x <= ls; x++ {
		line, _, err := scanner.ReadLine();
		if err != nil {
			// handle error
			panic(err)
		}
		startingNum := string(line)
		num := splitInNums(strings.Split(startingNum, ""))
		if err != nil {
			// handle error
			panic(err)
		}

		//fmt.Printf("Case #%d: %s <-- %s\n", x, tidyNum(num), startingNum)
		fmt.Printf("Case #%d: %s\n", x, tidyNum(num))

	}
}
func splitInNums(numbers []string) []int {
	res := make([]int, len(numbers))
	var err error
	for i, v := range numbers {
		res[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	return res
}
func tidyNum(num []int) string {

	for x := len(num) - 1; x > 0; x-- {
		if num[x] < num[x-1] {
			num[x-1] = num[x-1] - 1
			for j := x; j < len(num); j++ {
				num[j] = 9
			}
		}
	}
	res := make([]string, len(num))
	for i, v := range num {
		res[i] = fmt.Sprintf("%d", v)
	}
	n, err := strconv.Atoi(strings.Join(res, ""))
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%d", n)
}

//func isTidy(num []int) bool {
//	for i:=len(num);i<le
//	for ok := true; ok; ok = num > 0 {
//		current := num % 10
//		if last < current {
//			return false
//		}
//		last = current
//		num = num / 10;
//	}
//	return true
//}
