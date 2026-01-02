package app

import (
	"log"
	"strconv"
	"strings"
)

type Day1 struct {
	Log *log.Logger
}

func NewDay1(log *log.Logger) Day1 {
	return Day1{
		Log: log,
	}
}

func (d Day1) Part1(content string) error {
	position := 50
	result := 0
	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			break
		}
		direction := line[0]
		count, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}
		d.Log.Printf("Direction %s => Position: %d Step: %d", string(direction), position, count)
		if direction == 76 {
			position -= (count % 100)
			if position < 0 {
				position += 100
			}
		} else {
			position += count
			position %= 100
		}

		if position == 0 {
			result++
		}
		d.Log.Printf("\t\tNew Position: %d Count: %d", position, result)
	}
	d.Log.Println("Result:", result)
	return nil
}

func (d Day1) Part2(content string) (err error) {
	position := 50
	result := 0
	for _, line := range strings.Split(content, "\n") {
		if len(line) == 0 {
			break
		}
		direction := line[0]
		step, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}
		directionSymbol := "+"
		directionFactor := 1
		if direction == 76 {
			directionSymbol = "-"
			directionFactor = -1
		}
		d.Log.Printf("Position: %d Step: %s%d", position, directionSymbol, step)
		result += step / 100
		newPosition := position + ((step % 100) * directionFactor)
		if newPosition <= 0 {
			if newPosition < 0 {
				newPosition += 100
			}
			if position != 0 {
				result++
			}
		}
		if newPosition >= 100 {
			newPosition %= 100
			result++
		}
		d.Log.Printf("Position: %d Count: %d", newPosition, result)
		d.Log.Printf("")
		position = newPosition
	}
	d.Log.Println("Result: ", result)
	return nil
}
