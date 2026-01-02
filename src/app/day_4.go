package app

import (
	"log"
	"strings"
)

type Day4 struct {
	Log *log.Logger
}

var DIRS = [][]int{
	{1, 0},
	{0, 1},
	{1, 1},
	{-1, 0},
	{0, -1},
	{-1, 1},
	{1, -1},
	{-1, -1},
}

func NewDay4(log *log.Logger) Day4 {
	return Day4{
		Log: log,
	}
}

func (d Day4) getGrid(content string) [][]int {
	grid := [][]int{}
	for i, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			break
		}
		grid = append(grid, make([]int, len(line)))
		for j, el := range line {
			if el == '@' {
				grid[i][j] = 1
			} else {
				grid[i][j] = 0
			}
		}
	}
	return grid
}

func (d Day4) displayGrid(grid [][]int) {
	for _, row := range grid {
		d.Log.Println(row)
	}
}

func (d Day4) Part1(content string) error {
	grid := d.getGrid(content)
	d.displayGrid(grid)
	result := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == 0 {
				continue
			}

			count := 0
			for _, dir := range DIRS {
				x := r + dir[0]
				y := c + dir[1]

				if (x < 0 || x > len(grid)-1) || (y < 0 || y > len(grid[r])-1) {
					continue
				}

				if grid[x][y] > 0 {
					count += 1
				}
			}
			if count < 4 {
				result++
				grid[r][c] = 2
			}
		}
	}

	d.Log.Println("Updated grid:")
	d.displayGrid(grid)
	d.Log.Println("Result", result)
	return nil
}

func (d Day4) Part2(content string) error {
	grid := d.getGrid(content)
	d.displayGrid(grid)
	result := 0

	run := true

	for run {
		run = false
		for r := 0; r < len(grid); r++ {
			for c := 0; c < len(grid[r]); c++ {
				if grid[r][c] == 0 {
					continue
				}

				count := 0
				for _, dir := range DIRS {
					x := r + dir[0]
					y := c + dir[1]

					if (x < 0 || x > len(grid)-1) || (y < 0 || y > len(grid[r])-1) {
						continue
					}

					if grid[x][y] > 0 {
						count += 1
					}
				}
				if count < 4 {
					result++
					grid[r][c] = 0
					run = true
				}
			}
		}
		d.Log.Println("Updated grid:")
		d.Log.Println(run)
		d.displayGrid(grid)
	}
	d.Log.Println("Result", result)
	return nil
}
