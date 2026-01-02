package app

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day8 struct {
	Log *log.Logger
}

func NewDay8(log *log.Logger) Day8 {
	return Day8{
		Log: log,
	}
}

type Connection struct {
	start int
	end   int
	d     float64
}

func (d Day8) Part1(content string) error {
	result := 1
	positions := [][]int{}
	for i, line := range strings.Split(content, "\n") {
		positions = append(positions, make([]int, 3))
		for j, ns := range strings.Split(line, ",") {
			if num, err := strconv.Atoi(ns); err != nil {
				return err
			} else {
				positions[i][j] = num
			}
		}
	}

	ch := []Connection{}
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			distance := math.Sqrt(math.Pow(float64(positions[i][0]-positions[j][0]), 2) + math.Pow(float64(positions[i][1]-positions[j][1]), 2) + math.Pow(float64(positions[i][2]-positions[j][2]), 2))
			ch = append(ch, Connection{i, j, distance})
		}
	}

	slices.SortStableFunc(ch, func(i, j Connection) int {
		return int(i.d) - int(j.d)
	})

	circuits := [][]int{}
	visited := []int{}
	for _, c := range ch[:1000] {
		d.Log.Println(c)
		svCheck := slices.Contains(visited, c.start)
		evCheck := slices.Contains(visited, c.end)

		if !svCheck && !evCheck {
			circuits = append(circuits, []int{c.start, c.end})
			visited = append(visited, c.start, c.end)
		} else if svCheck && evCheck {
			sIdx := slices.IndexFunc(circuits, func(circuit []int) bool {
				return slices.Contains(circuit, c.start)
			})
			eIdx := slices.IndexFunc(circuits, func(circuit []int) bool {
				return slices.Contains(circuit, c.end)
			})

			if sIdx != eIdx {
				nCircuit := circuits[sIdx]
				nCircuit = append(nCircuit, circuits[eIdx]...)

				circuits = slices.DeleteFunc(circuits, func(circuit []int) bool {
					return slices.Contains(circuit, c.end) || slices.Contains(circuit, c.start)
				})

				circuits = append(circuits, nCircuit)
			}

		} else {
			pv := c.start
			av := c.end
			if !svCheck {
				pv = c.end
				av = c.start
			}
			visited = append(visited, av)
			vIdx := slices.IndexFunc(circuits, func(circuit []int) bool {
				return slices.Contains(circuit, pv)
			})
			circuits[vIdx] = append(circuits[vIdx], av)
		}
		d.Log.Println(circuits)
	}

	slices.SortStableFunc(circuits, func(a, b []int) int {
		return len(b) - len(a)
	})

	for _, circuit := range circuits[:3] {
		d.Log.Println(circuit)
		result *= len(circuit)
	}

	d.Log.Println("Result:", result)
	return nil
}

func (d Day8) Part2(content string) error {
	result := 1
	positions := [][]int{}
	for i, line := range strings.Split(content, "\n") {
		positions = append(positions, make([]int, 3))
		for j, ns := range strings.Split(line, ",") {
			if num, err := strconv.Atoi(ns); err != nil {
				return err
			} else {
				positions[i][j] = num
			}
		}
	}

	ch := []Connection{}
	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			distance := math.Sqrt(math.Pow(float64(positions[i][0]-positions[j][0]), 2) + math.Pow(float64(positions[i][1]-positions[j][1]), 2) + math.Pow(float64(positions[i][2]-positions[j][2]), 2))
			ch = append(ch, Connection{i, j, distance})
		}
	}

	slices.SortStableFunc(ch, func(i, j Connection) int {
		return int(i.d) - int(j.d)
	})

	circuits := [][]int{}
	visited := []int{}

	for _, c := range ch {
		d.Log.Println(c)
		svCheck := slices.Contains(visited, c.start)
		evCheck := slices.Contains(visited, c.end)

		if !svCheck && !evCheck {
			circuits = append(circuits, []int{c.start, c.end})
			visited = append(visited, c.start, c.end)
		} else if svCheck && evCheck {
			sIdx := slices.IndexFunc(circuits, func(circuit []int) bool {
				return slices.Contains(circuit, c.start)
			})
			eIdx := slices.IndexFunc(circuits, func(circuit []int) bool {
				return slices.Contains(circuit, c.end)
			})

			if sIdx != eIdx {
				nCircuit := circuits[sIdx]
				nCircuit = append(nCircuit, circuits[eIdx]...)

				circuits = slices.DeleteFunc(circuits, func(circuit []int) bool {
					return slices.Contains(circuit, c.end) || slices.Contains(circuit, c.start)
				})

				circuits = append(circuits, nCircuit)
			}

		} else {
			pv := c.start
			av := c.end
			if !svCheck {
				pv = c.end
				av = c.start
			}
			visited = append(visited, av)
			vIdx := slices.IndexFunc(circuits, func(circuit []int) bool {
				return slices.Contains(circuit, pv)
			})
			circuits[vIdx] = append(circuits[vIdx], av)
		}
		d.Log.Println(circuits)
		if len(visited) == len(positions) && len(circuits) == 1 {
			result = positions[c.start][0] * positions[c.end][0]
			break
		}
	}

	d.Log.Println("Result:", result)
	return nil
}
