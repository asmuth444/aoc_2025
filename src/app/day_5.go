package app

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type Day5 struct {
	Log *log.Logger
}

func NewDay5(log *log.Logger) Day5 {
	return Day5{
		Log: log,
	}
}

func (d Day5) Part1(content string) error {
	stop := false
	ranges := [][]int{}
	ids := []int{}
	result := 0
	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			if stop {
				break
			}
			stop = !stop
			continue
		}
		if !stop {
			r := strings.Split(line, "-")
			s, err := strconv.Atoi(r[0])
			if err != nil {
				return err
			}
			e, err := strconv.Atoi(r[1])
			if err != nil {
				return err
			}
			ranges = append(ranges, []int{s, e})
		} else {
			if id, err := strconv.Atoi(line); err != nil {
				return err
			} else {
				ids = append(ids, id)
			}
		}
	}

	slices.SortFunc(ranges, func(a, b []int) int {
		if a[0]-b[0] == 0 {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	for _, id := range ids {
		for _, mr := range ranges {
			if id >= mr[0] && mr[1] >= id {
				d.Log.Println(id)
				d.Log.Println(mr)
				result++
				break
			}
		}
	}

	d.Log.Println(ranges)
	d.Log.Println(len(ranges))
	d.Log.Println(ids)
	d.Log.Println(len(ids))
	d.Log.Println("Result:", result)
	return nil
}

func (d Day5) Part2(content string) error {
	ranges := [][]int{}
	result := 0
	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			break
		}

		r := strings.Split(line, "-")
		s, err := strconv.Atoi(r[0])
		if err != nil {
			return err
		}
		e, err := strconv.Atoi(r[1])
		if err != nil {
			return err
		}
		ranges = append(ranges, []int{s, e})
	}
	slices.SortFunc(ranges, func(a, b []int) int {
		if a[0] == b[0] {
			return a[0] - b[0]
		}
		return a[0] - b[0]
	})

	cm := 0
	for _, r := range ranges {
		if r[1] >= cm {
			result += r[1] - max(r[0], cm) + 1
			cm = r[1] + 1
		}
	}

	d.Log.Println("Results", result)

	return nil
}
