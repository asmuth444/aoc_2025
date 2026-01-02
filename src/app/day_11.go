package app

import (
	"log"
	"strings"
)

type Day11 struct {
	Log *log.Logger
}

func NewDay11(log *log.Logger) Day11 {
	return Day11{
		Log: log,
	}
}

func (d Day11) Part1(content string) error {
	g := make(map[string][]string)
	result := 0
	for _, line := range strings.Split(content, "\n") {
		sl := strings.Split(line, ": ")
		g[sl[0]] = strings.Split(sl[1], " ")
	}
	for u, vs := range g {
		d.Log.Println(u, vs)
	}

	q := []string{"you"}
	visited := make(map[string]bool)
	for len(q) > 0 {
		d.Log.Println(q)
		nq := len(q)
		for nq > 0 {
			cur := q[0]
			q = q[1:]
			nq--
			if cur == "out" {
				result++
				continue
			}
			if !visited[cur] {
				visited[cur] = true
			}
			for _, node := range g[cur] {
				if !visited[node] {
					q = append(q, node)
				}
			}
		}
	}

	d.Log.Println("result", result)
	return nil
}

type Node struct {
	name string
	pass int
}

func (d Day11) countPaths(g map[string][]string, start, end string, memo map[string]int) int {
	if start == end {
		return 1
	}
	if val, ok := memo[start]; ok {
		return val
	}
	total := 0

	for _, next := range g[start] {
		total += d.countPaths(g, next, end, memo)
	}

	memo[start] = total
	return total
}

func (d Day11) Part2(content string) error {
	g := make(map[string][]string)
	result := 0
	for _, line := range strings.Split(content, "\n") {
		sl := strings.Split(line, ": ")
		g[sl[0]] = strings.Split(sl[1], " ")
	}
	for u, vs := range g {
		d.Log.Println(u, vs)
	}

	result += d.countPaths(g, "svr", "dac", make(map[string]int)) * d.countPaths(g, "dac", "fft", make(map[string]int)) * d.countPaths(g, "fft", "out", make(map[string]int))
	result += d.countPaths(g, "svr", "fft", make(map[string]int)) * d.countPaths(g, "fft", "dac", make(map[string]int)) * d.countPaths(g, "dac", "out", make(map[string]int))

	d.Log.Println("result", result)
	return nil
}
