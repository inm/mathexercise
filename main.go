// command mathexcercise for practicing math
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	total := 0
	tok := 0

	for {
		max := int64(0)
		min := int64(1<<63 - 1)
		numbers := make([]int64, 5)
		mode := r.Intn(2)

		numbers[0] = r.Int63n(101)
		for i := 1; i < len(numbers); i++ {
			numbers[i] = r.Int63n(101) + numbers[i-1]
			d := numbers[i] - numbers[i-1]
			if d > max {
				max = d
			}
			if d < min {
				min = d
			}
		}
		if total == 0 {
			fmt.Print(" ") // space instead of +/-
		}
		switch mode {
		case 0:
			fmt.Print("max ", numbers, " ")
		case 1:
			fmt.Print("min ", numbers, " ")
		}

		var s string
		ok := false

		fmt.Scanf("%s", &s)
		res, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return
		}
		switch mode {
		case 0:
			ok = res == max
		case 1:
			ok = res == min
		}
		if ok {
			fmt.Print("+")
			tok++
		} else {
			fmt.Print("-")
		}
		total++
	}
}
