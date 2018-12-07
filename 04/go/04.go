package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
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
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	lines = strings.Split(strings.TrimSpace(string(data)), "\n")
	return
}

// sortByTime sorts an array of strings by a Time object extracted
// using this regex `\[(.*)\].*` and layout `2006-01-02 15:04`
func sortByTime(lines []string) []string {
	re := regexp.MustCompile(`\[(.*)\].*`)
	timeLayout := `2006-01-02 15:04`
	sort.Slice(lines, func(i, j int) bool {
		ti, _ := time.Parse(timeLayout, re.FindStringSubmatch(lines[i])[1])
		tj, _ := time.Parse(timeLayout, re.FindStringSubmatch(lines[j])[1])
		return ti.Unix() < tj.Unix()
	})
	return lines
}

func iterLines(lines []string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, line := range lines {
			out <- line
		}
	}()
	return out
}

func getGuardId(s string) (int, bool) {
	re := regexp.MustCompile(`.*Guard\s#(\d+).*`)
	if re.MatchString(s) {
		gid, _ := strconv.Atoi(re.FindStringSubmatch(s)[1])
		return gid, true
	}
	return -1, false
}

func getMinute(s string) int {
	re := regexp.MustCompile(`\[(.*)\].*`)
	timeLayout := `2006-01-02 15:04`
	t, _ := time.Parse(timeLayout, re.FindStringSubmatch(s)[1])
	return t.Minute()
}

type Minutes struct {
	sync.RWMutex
	_v []int
}

type Guards struct {
	sync.RWMutex
	_v map[int]Minutes
}

type GuardsItem struct {
	key   int
	value Minutes
}

func (mins *Minutes) Add(start, stop int) {
	mins.Lock()
	defer mins.Unlock()
	for i := start; i < stop; i++ {
		mins._v[i]++
	}
}

func (g *Guards) Get(key int) (Minutes, bool) {
	g.Lock()
	defer g.Unlock()
	v, ok := g._v[key]
	return v, ok
}

func (g *Guards) Set(key int, value Minutes) {
	g.Lock()
	defer g.Unlock()
	g._v[key] = value
}

func (g *Guards) Iter() <-chan GuardsItem {
	c := make(chan GuardsItem)
	go func() {
		g.Lock()
		defer g.Unlock()
		for k, v := range g._v {
			c <- GuardsItem{key: k, value: v}
		}
		close(c)
	}()
	return c
}

func main() {
	// Load input data.
	lines, err := loadData("../04.txt")
	check(err)

	l := iterLines(sortByTime(lines))

	guards := Guards{_v: make(map[int]Minutes)}

	var gid int
	for {
		s, ok := <-l
		if !ok {
			break
		}
		if tmp, ok := getGuardId(s); ok {
			gid = tmp
			v, ok := guards.Get(gid)
			if !ok {
				v = Minutes{_v: make([]int, 60)}
				guards.Set(gid, v)
			}
		} else {
			go func(gid int, s1, s2 string) {
				v, _ := guards.Get(gid)
				asleep := getMinute(s1)
				awake := getMinute(s2)
				v.Add(asleep, awake)
			}(gid, s, <-l)
		}
	}

	g := make(chan [4]int)
	for gitem := range guards.Iter() {
		gid, mins := gitem.key, gitem.value
		go func(gid int, mins []int) {
			var sum, max_freq, max_i int
			for i, freq := range mins {
				sum += freq
				if freq > max_freq {
					max_freq = freq
					max_i = i
				}
			}
			g <- [4]int{gid, sum, max_freq, max_i}
		}(gid, mins._v)
	}

	var p1_gid, p1_minute, p2_gid, p2_minute int
	var p1_max, p2_max int
	for range guards.Iter() {
		tmp := <-g
		gid, sum, max_freq, max_i := tmp[0], tmp[1], tmp[2], tmp[3]
		if sum > p1_max {
			p1_max, p1_gid, p1_minute = sum, gid, max_i
		}
		if max_freq > p2_max {
			p2_max, p2_gid, p2_minute = max_freq, gid, max_i
		}
	}
	close(g)

	// Output part 1 result
	fmt.Println(p1_gid * p1_minute)

	// Output part 2 result
	fmt.Println(p2_gid * p2_minute)
}
