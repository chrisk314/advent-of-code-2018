package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// check checks if an error occurred and panics if so.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// loadData reads the file specified by path. Returns lines in file as
// array of strings and any error encountered.
func loadData(path string) (lines []string, err error) {
	data, err := ioutil.ReadFile("../03.txt")
	if err != nil {
		return
	}
	lines = strings.Split(strings.TrimSpace(string(data)), "\n")
	return
}

// regexpExtractDims extracts cloth dimensions from the AoC2018 day 3
// input. Returns N * 4 array of ints and any error encountered.
func regexpExtractDims(lines []string) (dims []([4]int), err error) {
	re := regexp.MustCompile(`#\d+\s@\s(\d+),(\d+):\s(\d+)x(\d+)`)
	dims = make([]([4]int), len(lines))
	for i, line := range lines {
		for j, dim := range re.FindStringSubmatch(line)[1:] {
			dims[i][j], err = strconv.Atoi(dim)
			if err != nil {
				return
			}
		}
	}
	return
}

// getMaxDims finds the maximum horizontal and vertical dimensions
// in an N * 4 array of ints.
func getMaxDims(dims []([4]int)) (max_h, max_v int) {
	max_h, max_v = dims[0][0], dims[0][0]
	for _, dim := range dims {
		h := dim[0] + dim[2]
		v := dim[1] + dim[3]
		if h > max_h {
			max_h = h
		}
		if v > max_v {
			max_v = v
		}
	}
	return
}

func getStartStop(tid, np, n int) (start, stop int) {
	npp, r := n/np, n%np
	start, stop = tid*npp, (tid+1)*npp
	if tid < r {
		start, stop = start+tid, stop+tid+1
	} else {
		start, stop = start+r, stop+r
	}
	return
}

func main() {
	// Load input data.
	lines, err := loadData("../03.txt")
	check(err)

	// Regexp extract cloth dimensions.
	dims, err := regexpExtractDims(lines)
	check(err)

	// Find square inches claimed more than once.
	max_h, max_v := getMaxDims(dims)

	np := runtime.NumCPU()

	claims := make([]int, max_h*max_v)

	wait := make(chan bool)
	for tid := 0; tid < np; tid++ {
		go func(tid int) {
			startN, stopN := getStartStop(tid, np, max_v)
			for _, dim := range dims {
				if !(dim[1] < stopN && dim[1]+dim[3] > startN) {
					continue
				}
				start, stop := startN, stopN
				if dim[1] > start {
					start = dim[1]
				}
				if dim[1]+dim[3] < stop {
					stop = dim[1] + dim[3]
				}
				for i := start; i < stop; i++ {
					for j := dim[0]; j < dim[0]+dim[2]; j++ {
						claims[i*max_h+j]++
					}
				}
			}
			wait <- true
		}(tid)
	}
	for tid := 0; tid < np; tid++ {
		<-wait
	}

	c := make(chan int)

	dblClaimCnt := 0
	for tid := 0; tid < np; tid++ {
		go func(tid int) {
			sum := 0
			startN, stopN := getStartStop(tid, np, max_v)
			for i := startN; i < stopN; i++ {
				for j := 0; j < max_h; j++ {
					if claims[i*max_h+j] > 1 {
						sum++
					}
				}
			}
			c <- sum
		}(tid)
	}
	for tid := 0; tid < np; tid++ {
		dblClaimCnt += <-c
	}
	close(c)

	// Output part 1 result
	fmt.Println(dblClaimCnt)

	c = make(chan int)

	go func() {
		stop := make(chan struct{})
		defer close(c)
		for id, dim := range dims {
			go func(id int, dim [4]int) {
			loop:
				for {
					select {
					case <-stop:
						return
					default:
						for i := dim[1]; i < dim[1]+dim[3]; i++ {
							for j := dim[0]; j < dim[0]+dim[2]; j++ {
								if claims[i*max_h+j] > 1 {
									break loop
								}
							}
						}
						c <- id + 1
						close(stop)
						return
					}
				}
			}(id, dim)
		}
	}()

	// Output part 2 result
	fmt.Printf("#%d\n", <-c)
}
