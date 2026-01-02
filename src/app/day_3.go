package app

import (
	"log"
	"strconv"
	"strings"
)

type Day3 struct {
	Log *log.Logger
}

func NewDay3(log *log.Logger) Day3 {
	return Day3{
		Log: log,
	}
}

func (d Day3) Part1(content string) error {
	result := 0
	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			continue
		}
		d.Log.Println(line)
		fNum := -1
		sNum := -1
		sHNum := -1
		l := len(line) - 1
		for l >= 0 {
			pNum, err := strconv.Atoi(string(line[l]))
			if err != nil {
				return err
			}
			if pNum >= fNum {
				sNum = fNum
				fNum = pNum
			}
			if pNum != fNum && pNum > sHNum {
				sHNum = pNum
			}
			l--
			d.Log.Println(pNum, fNum, sNum, sHNum)
		}
		if sNum == -1 {
			sNum = fNum
			fNum = sHNum
		}
		d.Log.Printf("Adding %d%d", fNum, sNum)
		result += (fNum * 10) + sNum
	}
	d.Log.Println("Result:", result)
	return nil
}

func (d Day3) Part2(content string) error {
	result := 0
	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			continue
		}

		nums := make([]int, 0, 12)
		start := 0
		size := 12

		for size > 0 {
			end := len(line) - size + 1
			max := 0
			mIdx := 0

			for i := start; i < end; i++ {
				pNum, err := strconv.Atoi(string(line[i]));
				if err != nil {
					return err
				}

				if pNum > max {
					max = pNum
					mIdx =  i - start
				}
			}

			nums = append(nums, max)
			start += mIdx + 1
			size--
		}
		d.Log.Println(line)
		subResult := 0
		for _, num := range nums {
			subResult = subResult * 10 + num
		}
		d.Log.Println(subResult)
		result += subResult
	}
	d.Log.Println("Result: ", result)
	return nil
}