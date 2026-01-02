package app

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Polygon []Point

type Day9 struct {
	Log *log.Logger
	pc  map[Point]bool
	lc  map[Point]map[Point]bool
}

func NewDay9(log *log.Logger) Day9 {
	return Day9{
		Log: log,
		pc:  make(map[Point]bool),
		lc:  make(map[Point]map[Point]bool),
	}
}

func (p1 Point) area(p2 Point) int {
	return int((math.Abs(float64(p1.x-p2.x)) + 1) * (math.Abs(float64(p1.y-p2.y)) + 1))
}

func (d Day9) Part1(content string) error {
	result := 0
	polygon := Polygon{}
	for _, line := range strings.Split(content, "\n") {
		point := make([]int, 2)
		for i, ns := range strings.Split(line, ",") {
			if num, err := strconv.Atoi(ns); err != nil {
				return err
			} else {
				point[i] = num
			}
		}
		polygon = append(polygon, Point{point[0], point[1]})
	}

	for i := 0; i < len(polygon); i++ {
		for j := i + 1; j < len(polygon); j++ {
			area := polygon[i].area(polygon[j])
			if area > result {
				d.Log.Println(polygon[i], polygon[j], area, result)
				result = area
			}
		}
	}

	d.Log.Printf("Result:%d", result)
	return nil
}

func (d Day9) compressPoints(points []Point) (map[int]int, map[int]int) {
	xPoints := []int{}
	yPoints := []int{}
	for _, point := range points {
		xPoints = append(xPoints, point.x)
		yPoints = append(yPoints, point.y)
	}
	slices.Sort(xPoints)
	xPoints = slices.Compact(xPoints)
	slices.Sort(yPoints)
	yPoints = slices.Compact(yPoints)

	xMap := make(map[int]int)
	yMap := make(map[int]int)

	for i, xPoint := range xPoints {
		xMap[xPoint] = i
	}
	for i, yPoint := range yPoints {
		yMap[yPoint] = i
	}

	return xMap, yMap
}

// https://github.com/sleekmountaincat/aoc2025/blob/main/src/day9/q2.ts
func (d Day9) Part2(content string) error {
	result := 0
	oPolygon := Polygon{}
	for _, line := range strings.Split(content, "\n") {
		point := make([]int, 2)
		for i, ns := range strings.Split(line, ",") {
			if num, err := strconv.Atoi(ns); err != nil {
				return err
			} else {
				point[i] = num
			}
		}
		oPolygon = append(oPolygon, Point{point[0], point[1]})
	}

	xMap, yMap := d.compressPoints(oPolygon)

	d.Log.Println(len(xMap))
	d.Log.Println(len(yMap))

	grid := [][]int{}

	for i := range len(xMap) {
		grid = append(grid, make([]int, len(yMap)))
		for j := range len(yMap) {
			grid[i][j] = 0
		}
	}

	zPolygon := Polygon{}
	for _, point := range oPolygon {
		grid[xMap[point.x]][yMap[point.y]] = 1
		zPolygon = append(zPolygon, Point{xMap[point.x], yMap[point.y]})
	}

	for i := range len(zPolygon) {
		p1 := zPolygon[i]
		p2 := zPolygon[(i+1)%len(oPolygon)]
		if p1.x == p2.x {
			for i := min(p1.y, p2.y); i <= max(p1.y, p2.y); i++ {
				grid[p1.x][i] = 1
			}
		} else {
			for i := min(p1.x, p2.x); i <= max(p1.x, p2.x); i++ {
				grid[i][p1.y] = 1
			}
		}
	}

	insidePoint := Point{-1, -1}
	for x := range len(grid) {
		for y := range len(grid[0]) {
			if grid[x][y] != 0 {
				continue
			}

			hitsLeft := 0
			prev := 0

			for i := x; i >= 0; i-- {
				cur := grid[i][y]
				if cur != prev {
					hitsLeft++
				}
				prev = cur
			}

			if hitsLeft%2 == 1 {
				insidePoint.x = x
				insidePoint.y = y
				break
			}
		}
		if insidePoint.x != -1 && insidePoint.y != -1 {
			break
		}
	}

	for _, row := range grid {
		d.Log.Println(row)
	}

	d.Log.Println("Flood fill start", insidePoint)
	np := []Point{insidePoint}
	for len(np) > 0 {
		cur := np[0]
		np = np[1:]
		if grid[cur.x][cur.y] == 0 {
			grid[cur.x][cur.y] = 1
			for _, dir := range DIRS {
				x := cur.x + dir[0]
				y := cur.y + dir[1]
				if (x >= 0 && x < len(grid)) && (y >= 0 && y < len(grid[0])) {
					np = append(np, Point{x, y})
				}
			}
		}
	}
	d.Log.Println("Flood fill done")

	for _, row := range grid {
		d.Log.Println(row)
	}

	for i := range len(oPolygon) {
		for j := i + 1; j < len(oPolygon); j++ {
			p1 := Point{xMap[oPolygon[i].x], yMap[oPolygon[i].y]}
			p2 := Point{xMap[oPolygon[j].x], yMap[oPolygon[j].y]}

			check := true
			for x := min(p1.x, p2.x); x <= max(p1.x, p2.x); x++ {
				for y := min(p1.y, p2.y); y <= max(p1.y, p2.y); y++ {
					if grid[x][y] == 0 {
						d.Log.Println("Rejected:", oPolygon[i], oPolygon[j])
						check = false
						break
					}
				}
				if !check {
					break
				}
			}

			if check {
				result = max(result, oPolygon[i].area(oPolygon[j]))
				d.Log.Println(result, oPolygon[i], oPolygon[j])
			}
		}
	}

	d.Log.Printf("Result:%d", result)
	return nil
}
