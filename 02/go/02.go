package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Load data from file
	data, err := ioutil.ReadFile("../02.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1] // drop empty line at end

	// Get chksum for list of ids
	counts := []int{0, 0}

	for _, s := range lines {
		m := make(map[rune]int)
		for _, c := range s {
			m[c]++
		}
		var twos, threes int
		for _, v := range m {
			switch v {
			case 2:
				twos = 1
			case 3:
				threes = 1
			}
			if twos == 1 && threes == 1 {
				break
			}
		}
		counts[0] += twos
		counts[1] += threes
	}

	chksum := counts[0] * counts[1]

	// Output part 1 result
	fmt.Println(chksum)

	// Find strings with smallest Levenshtein distance
	min_dist := len(lines[0])
	min_i, min_j := len(lines), len(lines)
	for i, s1 := range lines {
		for j, s2 := range lines[:i] {
			dist := 0
			for k := 0; k < len(s1); k++ {
				if s1[k] != s2[k] {
					dist++
					if dist == min_dist {
						break
					}
				}
			}
			if dist < min_dist {
				min_dist, min_i, min_j = dist, i, j
			}
		}
	}

	similar := make([]byte, 0, len(lines[0]))
	s1, s2 := lines[min_i], lines[min_j]
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			similar = append(similar, s1[i])
		}
	}

	// Output part 2 result
	fmt.Println(string(similar))
}
