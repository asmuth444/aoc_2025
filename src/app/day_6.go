package app

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Day6 struct {
	Log *log.Logger
}

func NewDay6(log *log.Logger) Day6 {
	return Day6{
		Log: log,
	}
}

func (d Day6) Part1(content string) error {
	lines := strings.Split(content, "\n")
	ops := strings.Split(regexp.MustCompile(`\s+`).ReplaceAllString(lines[len(lines)-1], ","), ",")
	lines = lines[:len(lines)-1]
	l := len(lines)
	nums := [][]int{}
	result := 0
	for idx, line := range lines {
		if len(line) == 0 {
			break
		}
		if idx < l {
			nums = append(nums, []int{})
		}
		line = strings.TrimSpace(line)
		line = regexp.MustCompile(`\s+`).ReplaceAllString(line, ",")
		d.Log.Println(line)
		for _, el := range strings.Split(line, ",") {
			if num, err := strconv.Atoi(el); err == nil {
				nums[idx] = append(nums[idx], num)
			} else {
				return err
			}
		}
	}
	d.Log.Println(nums)
	d.Log.Println(ops)

	for i := 0; i < len(nums[0]); i++ {
		subResult := nums[0][i]
		op := ops[i]
		for j := 1; j < len(nums); j++ {
			if op == "+" {
				subResult += nums[j][i]
			} else {
				subResult *= nums[j][i]
			}
		}
		d.Log.Println("\t", subResult)
		result += subResult
	}

	d.Log.Println("Result:", result)
	return nil
}

func (d Day6) Part2(content string) error {
	lines := strings.Split(content, "\n")
	d.Log.Println(lines)
	c := len(lines) - 1
	r := len(lines[0]) - 1
	lc := len(lines[c])
	for i := lc; i <= r; i++ {
		lines[c] += " "
	}
	for _, line := range lines {
		d.Log.Println(len(line), line)
	}
	result := 0
	d.Log.Println(c, r)

	nsa := [][]string{{""}}
	operation := ""
	nIdx := 0
	for j := r; j >= 0; j-- {
		for i := c; i >= 0; i-- {
			if i == c {
				if len(operation) == 0 && (lines[i][j] == '+' || lines[i][j] == '*') {
					operation = string(lines[i][j])
				} else if len(operation) > 0 && lines[i][j] == ' ' {
					nsa[len(nsa)-1][nIdx] = operation
					nsa = append(nsa, []string{})
					operation = ""
					nIdx = -1
				} else {
					operation = ""
				}
			} else {
				if lines[i][j] != ' ' {
					nsa[len(nsa)-1][nIdx] = string(lines[i][j]) + nsa[len(nsa)-1][nIdx]
				}
				if i == 0 {
					nIdx++
					nsa[len(nsa)-1] = append(nsa[len(nsa)-1], "")
				}
			}

			if i == 0 && j == 0 {
				nsa[len(nsa) - 1][nIdx] = operation
			}
			// d.Log.Println(nsa)
		}
	}

	for _, nss := range nsa {
		nums := []int{}
		for _, ns := range nss[:len(nss) - 1] {
			if len(ns) == 0 {
				continue
			}
			num, err := strconv.Atoi(strings.TrimSpace(ns))
			if err != nil {
				return err
			}
			nums = append(nums, num)
		}

		subResult := nums[0]
		operation :=  nss[len(nss) - 1]
		for _, num := range nums[1:] {
			// d.Log.Println(num)
			if operation == "+" {
				subResult += num
			} else {
				subResult *= num
			}
		}
		// d.Log.Println(subResult)
		result += subResult
	}
	d.Log.Println("Result:", result)
	return nil
}
