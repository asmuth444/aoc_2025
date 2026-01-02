package app

import (
	"fmt"
	"log"
	"strings"
)

type Day7 struct {
	Log *log.Logger
}

func NewDay7(log *log.Logger) Day7 {
	return Day7{
		Log: log,
	}
}

func (d Day7) Part1(content string) error {
	lines := strings.Split(content, "\n")
	r := len(lines)
	c := len(lines[0])
	m := [][]string{}
	result := 0
	for i := 0; i < r; i++ {
		m = append(m, []string{})
		for j := 0; j < c; j++ {
			m[i] = append(m[i], string(lines[i][j]))
		}
	}

	for i := 1; i < r; i++ {
		for j := 0; j < c; j++ {
			switch m[i-1][j] {
			case "S":
				m[i][j] = "|"
			case "|":
				switch m[i][j] {
				case "^":
					result++
					if j-1 >= 0 && m[i][j-1] == "." {
						m[i][j-1] = "|"
					}
					if j+1 <= c-1 && m[i][j+1] == "." {
						m[i][j+1] = "|"
					}
				case ".":
					m[i][j] = "|"
				}
			}
		}
	}

	for i, l := range m {
		d.Log.Println(i, l)
	}

	d.Log.Println("Result:", result)
	return nil
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

// https://www.reddit.com/r/adventofcode/comments/1pgi0sm/2025_day_07_part_2_python_efficient_algorithm_on/?share_id=0VeDuri7wB1144lNehXRy&utm_content=2&utm_medium=android_app&utm_name=androidcss&utm_source=share&utm_term=1
func (d Day7) Part2(content string) error {
	lines := strings.Split(content, "\n")
	r := len(lines)
	c := len(lines[0])
	m := [][]string{}
	result := 0
	g := make(map[string]int)
	for i := 0; i < r; i++ {
		m = append(m, []string{})
		for j := 0; j < c; j++ {
			switch lines[i][j] {
			case 'S':
				g[getKey(i+1, j)] = 1
			case '^':
				if pv, ok := g[getKey(i-1, j)]; ok {
					if j-1 >= 0 {
						g[getKey(i, j-1)] += pv
					}
					if j+1 < c {
						g[getKey(i, j+1)] += pv
					}
					delete(g, getKey(i-1, j))
				}
			default:
				if pv, ok := g[getKey(i-1, j)]; ok {
					g[getKey(i, j)] += pv
					delete(g, getKey(i-1, j))
				}
			}

		}
	}

	for k, v := range g {
		d.Log.Println(k, "->", v)
		result += v
	}

	d.Log.Println("Result:", result)
	return nil
}
