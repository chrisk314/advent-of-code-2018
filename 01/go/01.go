package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Load data from file
	data, err := ioutil.ReadFile("../01.txt")
	check(err)
	lines := strings.Split(string(data), "\n")
	nums := make([]int, len(lines)-1)
	for i, line := range lines {
		if len(line) == 0 {
			break
		}
		num, err := strconv.Atoi(line)
		check(err)
		nums[i] = num
	}

	// Sum frequency changes
	result := 0
	for _, num := range nums {
		result += num
	}

	// Output part 1 result
	fmt.Println(result)

	// Loop forever until same frequency observed twice
	m := make(map[int]int)
	result = 0
	i := 0
	for {
		idx := i % len(nums)
		result += nums[idx]
		_, ok := m[result]
		if ok {
			break
		} else {
			m[result] = result
		}
		i++
	}

	// Output part 2 result
	fmt.Println(result)
}
