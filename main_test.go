package main

import "testing"

var tidyNums = []int{
	123,
	122225,
	1255550,
	12345677,
	4563,
}

var tidyNumsRes = []bool{
	true,
	true,
	false,
	true,
	false,
}

func TestLastTidy(t *testing.T) {
	num := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}
	if tidyNum(num) != "99999999999999999" {
		t.FailNow()
	}
}
