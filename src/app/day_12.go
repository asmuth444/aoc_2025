package app

import (
	"log"
	"strconv"
	"strings"
)

type Day12 struct {
	Log *log.Logger
}

func NewDay12(log *log.Logger) Day12 {
	return Day12{
		Log: log,
	}
}

type Shape struct {
	s        [][]int
	size int
}

type Query struct {
	grid   []int
	sCount []int
}

func (d Day12) parseContent(content string) (shapes []Shape, queries []Query, err error) {
	sc := false
	for _, line := range strings.Split(content, "\n") {
		d.Log.Println(line)
		if len(line) == 0 {
			sc = false
			continue
		}
		if strings.Contains(line, ":") && !strings.Contains(line, "x") {
			sc = true
			shapes = append(shapes, Shape{[][]int{}, 0})
			continue
		}
		if sc {
			row := make([]int, len(line))
			count := 0
			for i, e := range line {
				if e == '#' {
					row[i] = 1
					count++
				}
			}
			shapes[len(shapes)-1].s = append(shapes[len(shapes)-1].s, row)
			shapes[len(shapes)-1].size += count
		}
		if strings.Contains(line, "x") {
			query := Query{make([]int, 2), []int{}}
			sl := strings.Split(line, ": ")
			for i, sm := range strings.Split(sl[0], "x") {
				num, err := strconv.Atoi(sm)
				if err != nil {
					return shapes, queries, err
				}
				query.grid[i] = num
			}
			for _, sm := range strings.Split(sl[1], " ") {
				num, err := strconv.Atoi(sm)
				if err != nil {
					return shapes, queries, err
				}
				query.sCount = append(query.sCount, num)
			}
			queries = append(queries, query)
		}
	}

	return shapes, queries, nil
}

func (d Day12) Part1(content string) error {
	shapes, queries, err := d.parseContent(content)
	if err != nil {
		return err
	}
	result := 0

	for _, shape := range shapes {
		for _, row := range shape.s {
			d.Log.Println(row)
		}
		d.Log.Println(shape.size)
		d.Log.Println()
	}

	for _, query := range queries {
		d.Log.Println(query)
	}

	for _, query := range queries {
		size := query.grid[0] * query.grid[1]
		osize := 0
		for si, sc := range query.sCount {
			osize += (sc * shapes[si].size)
		}
		d.Log.Println(osize, size)
		if osize <= size {
			result++
		}
	}

	d.Log.Println("Result:", result)
	return nil
}

func (d Day12) Part2(content string) error {
	panic("unimplemented")
}
