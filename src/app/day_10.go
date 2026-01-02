package app

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Day10 struct {
	Log *log.Logger
}

func NewDay10(log *log.Logger) Day10 {
	return Day10{
		Log: log,
	}
}

type Machine struct {
	Lights  []int
	Buttons [][]int
	Joltage []int
}

type Machines []Machine

func key(a []int) (result string) {
	return fmt.Sprint(a)
}

func (d Day10) Part1(content string) error {
	machines := Machines{}
	for line := range strings.SplitSeq(content, "\n") {
		machine := Machine{}
		for token := range strings.SplitSeq(line, " ") {
			switch token[0] {
			case '[':
				for _, e := range token[1 : len(token)-1] {
					if e == '#' {
						machine.Lights = append(machine.Lights, 1)
					} else {
						machine.Lights = append(machine.Lights, 0)
					}
				}
			case '(':
				button := []int{}
				for _, bN := range strings.Split(token[1:len(token)-1], ",") {
					bNum, err := strconv.Atoi(bN)
					if err != nil {
						return err
					}
					button = append(button, bNum)
				}
				machine.Buttons = append(machine.Buttons, button)
			}
		}
		// d.Log.Println(line)
		// d.Log.Println(machine)
		machines = append(machines, machine)
	}

	result := 0
	for _, machine := range machines {
		initial := make([]int, len(machine.Lights))
		q := [][]int{initial}
		visited := make(map[string]bool)
		visited[key(initial)] = true
		depth := 0
		for len(q) > 0 {
			k := len(q)
			stop := false
			for k > 0 {
				// d.Log.Println(q, visited, machine.Lights)
				cur := q[0]
				q = q[1:]
				if key(cur) == key(machine.Lights) {
					result += depth
					stop = true
					break
				}
				for _, button := range machine.Buttons {
					nn := make([]int, len(machine.Lights))
					copy(nn, cur)
					for _, b := range button {
						nn[b] = (nn[b] + 1) % 2
					}
					nnk := key(nn)
					if !visited[nnk] {
						visited[nnk] = true
						q = append(q, nn)
					}
					// d.Log.Println(cur, button, nn)
				}
				k--
			}
			if stop {
				break
			}
			depth++
		}
	}

	d.Log.Println("Result:", result)
	return nil
}

// https://topaz.github.io/paste/#XQAAAQDhCwAAAAAAAAA4GEiZzRd1JAgz+whYRQxSFI7XvmlfhtGDinguAj8sFyd/fs4fcP3B6Imi/tk3ZYXXaBodH8sU6wtG61aizJg/9+wHoSAFAM2FgKYOpsEgP8i6lRkBVMbZTjHdi2+8LEzUU19XNEtLdnikOgk2lWCVdivPPAlORRZ5HnWhq6K44DPKcZnpDbH71Up8dIikUHQmc05U6jptkJ+R6fSi792mw5AZ7h3/g8DMJeSqSDoG9cK/+DK12/jpL3Ya9YnCoUrXtRFLYX7C/tYgpBHhjeosWUDPmM8jLfj7yjFz8kqMLiZBAi7oPq5NqJlS4al8i/HOhyx2Bm/jjAXobXP5ES8c8A2XOVpY7zIwlf88P86dA5qd5DSzmDTsZOKfL55NXUe1uD9uKtDqO+Wpwx01s4o53kJ8DQmx5Mbbnm2PliFif54r67CmrpyiHJvjRrG4+F8wFP9WB/wvjTCTNa4jqQllnQ1x/aofuY/NOH5viOqpJAbDj3xpyJXoPO6lEYZU5yGrIkvXjfElF6Y5OxblluTT6BWWbDbVVaZcp8WsSF+cCfLpqcPuNCiRPXbo87rig2hUPb8xQiEBCQxnDEIVzz9aTs7+nZWBrLIhU1wIuS8kiZTwzxkHiAh7KvyDA+lrwYwoUPHAdMsRcHPfpGORuLKh7TW3yQioTQfAHxakHmWuU9/ZIxwpFNHK+2QsvAaxaLb8WlGY7m+8U4xLrlYYwfdZhg2RqxHUWLLhuV2iYcQLZPviLk3aqB1Eej5g8tJLq0DBnnd6U4t/CAy/X8X61asuxEzBfz5/Y7d9VDWJzH2hm+0Pq/A9DEtI/VlCk/zz4in4F2s3VHlUBL/SgS91P3Jhsxkx7ChiNUyW8Ct9Ghk9QkkbWIW2J4Gjaqxr2MCuEwWrw1jQviN3TesIxDoNghnjzrvbUGKyr0t74LFgTqqtyQdzwazjOnce9XRNiUAuYv9WOPIpnJ1n2Z9U8vWenvNkJU4yrhCTz1lei7899El4QDv1MKT/e0HNXDnCg2TDQXjSeAtRwbH+oF2tL0nEohl6fXY/GlaSC+NRlUcMU8q3R2vnuCvF5F5k6EZXMOt9N6bj5RzFwoPauydkwUvgnh3ohC8OJZI5T3ZfqCe5MkssesmuZAt8jlxbBOO5/iJJDa5g4IyDDTi478nlUssZ7Ny+glOU61do6Rw9EOI0VtcEfgc2KUqCE0b57j5GQpIWAZOaB5+1TgAErmkD5ZBhG99wjHXUouHMsNwHAq9vbiMfsTBTIFoNmVx+pMY7efDSQW8D2uAuNVHjPrHStMD+kMw/pXF4ew/T5kyC0AMugc0lq7K5kyQGO63ckpRBT0TEysLsrhkQ1NdZDs3KBMQHcGOtJBx0LxB2WRBJ65Nlasy6NyumGVmJHzNa3mTFv3u93TNbycst6jPqv8LFEwltgzpmyoMeLRTKW1KUqqRf/4KyLYlpov0EcAsv/Gp9COGFG0tz+miEsWslc1UsHfWZlt5USFX9DVeBmM3CMn40YjAZXZWwO0ebWqkHa8ptQD6yPxOukZ+tYsPlgVZY0GuAhzlvzr5rRbfEx76wrauYQbu4KqCmJwjcCHU1aWgzk3ZxCTPVPj8BLifm5+oR0RM1MJ9+P71M61GFVDOU6xqMAv9G0S+LL56csTyan/fULfPo0grAe/9w0HYA
func (d Day10) minPresses(buttons [][]int, target []int, memo map[string]int) int {
	// d.Log.Println(buttons, target, memo)
	allZero := true
	for _, v := range target {
		if v != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		return 0
	}
	k := key(target)
	if val, ok := memo[k]; ok {
		return val
	}
	n, m := len(buttons), len(target)
	limit := 1<<n
	best := -1
	// d.Log.Println(limit)
	for mask := range limit {
		remainder := make([]int, m)
		copy(remainder, target)
		cp1, poss := 0, true
		for b := range n {
			// d.Log.Println(mask, b, 1<<b, mask & (1<<b))
			if (mask & (1<<b)) != 0 {
				cp1++
				for i := range m {
					remainder[i] -= buttons[b][i];
				}
			}
		}
		// d.Log.Println(remainder)
		for i := range m {
			if (remainder[i] < 0) || (remainder[i] % 2 != 0) {
				poss = false
				break
			}
		}
		// d.Log.Println(poss)
		if poss {
			nextTarget := make([]int, m)
			for i := range m {
				nextTarget[i] = remainder[i] / 2
			}
			res := d.minPresses(buttons, nextTarget, memo)
			if res != -1 {
				totalCost := cp1 + (2 * res)
				if best == -1 || totalCost < best {
					best = totalCost
				}
			}
		}
	}
	memo[k] = best
	return best
}

func (d Day10) Part2(content string) error {
	machines := Machines{}
	result := 0
	for line := range strings.SplitSeq(content, "\n") {
		machine := Machine{}
		for token := range strings.SplitSeq(line, " ") {
			switch token[0] {
			case '[':
				for _, e := range token[1 : len(token)-1] {
					if e == '#' {
						machine.Lights = append(machine.Lights, 1)
					} else {
						machine.Lights = append(machine.Lights, 0)
					}
				}
			case '(':
				m := len(machine.Lights)
				button := make([]int, m)
				for _, bN := range strings.Split(token[1:len(token)-1], ",") {
					bNum, err := strconv.Atoi(bN)
					if err != nil {
						return err
					}
					if bNum < m {
						button[bNum] = 1
					}
				}
				machine.Buttons = append(machine.Buttons, button)
			case '{':
				for _, jN := range strings.Split(token[1:len(token)-1], ",") {
					jNum, err := strconv.Atoi(jN)
					if err != nil {
						return err
					}
					machine.Joltage = append(machine.Joltage, jNum)
				}
			}
		}
		// d.Log.Println(line)
		// d.Log.Println(machine)
		machines = append(machines, machine)
	}

	for i, machine := range machines {
		memo := make(map[string]int)
		subResult := d.minPresses(machine.Buttons, machine.Joltage, memo)
		result += subResult
		d.Log.Println("Machine", i, ":", subResult)
	}

	d.Log.Println("Result:", result)
	return nil
}