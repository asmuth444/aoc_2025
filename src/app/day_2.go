package app

import (
	"log"
	"strconv"
	"strings"
)

type Day2 struct {
	Log *log.Logger
}

func NewDay2(log *log.Logger) Day2 {
	return Day2{
		Log: log,
	}
}

func (d Day2) Part1(content string) error {
	result := 0
	for _, productRange := range strings.Split(content, ",") {
		indexes := strings.Split(productRange, "-")
		lIndex, err := strconv.Atoi(indexes[0])
		if err != nil {
			return err
		}
		rIndex, err := strconv.Atoi(indexes[1])
		if err != nil {
			return err
		}
		d.Log.Println(lIndex, rIndex)
		currentId := lIndex
		for currentId <= rIndex {
			d.Log.Println(currentId)
			sId := strconv.Itoa(currentId)
			sIdLen := len(sId)
			if sIdLen%2 == 0 {
				midPoint := sIdLen / 2
				d.Log.Println(sIdLen, currentId, midPoint, sId[:midPoint], sId[midPoint:])
				if sId[:midPoint] == sId[midPoint:] {
					result += currentId
				}
			}
			currentId++
		}
	}
	d.Log.Println("Result: ", result)
	return nil
}

func checkFrequency(value string, size int) bool {
	index := 0
	m := make(map[string]int)
	l := len(value)
	for index < l {
		ss := value[index:(index + size)]
		v, ok := m[ss]
		if !ok {
			m[ss] = 1
		} else {
			m[ss] = v + 1
		}
		index += size
	}
	return len(m) == 1
}

func (d Day2) Part2(content string) (err error) {
	result := 0
	for _, productRange := range strings.Split(content, ",") {
		indexes := strings.Split(productRange, "-")
		lIndex, err := strconv.Atoi(indexes[0])
		if err != nil {
			return err
		}
		rIndex, err := strconv.Atoi(indexes[1])
		if err != nil {
			return err
		}
		d.Log.Println(lIndex, rIndex)
		currentId := lIndex
		for currentId <= rIndex {
			sId := strconv.Itoa(currentId)
			l := len(sId)
			midPoint := l / 2
			idx := 1
			for idx <= midPoint {
				if l%idx == 0 && checkFrequency(sId, idx) {
					d.Log.Println("\tAdding", currentId, sId[:idx], idx)
					result += currentId
					break
				}
				idx++
			}
			currentId++
		}
	}
	d.Log.Println("Result: ", result)
	return nil
}
