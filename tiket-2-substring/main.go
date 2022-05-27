package tiket_2_substring

import (
	"fmt"
	"strconv"
	"strings"
)

// 5257523
// 525
// 752
func MaxSubstring(s string) string {
	total := []int{
		0, 0, 0,
	}

	fill := false
	fillIdx := 1

	for i := 0; i < len(s)-3; i++ {
		d, _ := strconv.Atoi(fmt.Sprintf("%c", s[i]))
		if !fill {
			if total[0] < d {
				total[0] = d
				fill = true
			}
		} else {
			total[fillIdx] = d
			if fillIdx == 2 {
				fillIdx = 1
				fill = false
			} else {
				fillIdx++
			}
		}

	}

	return fmt.Sprintf("%d%d%d", total[0], total[1], total[2])
}

func MaxSubstringWithoutConvert(s string) string {
	n := 3

	beginMax := ""
	max := ""
	fillIdx := 1
	for i := 0; i < len(s); i++ {
		d := string(s[i])

		if beginMax < d {
			beginMax = d
			max = d
			fillIdx = 1
		} else {
			if fillIdx < n {
				max = max + d
				fillIdx++
			}
		}
	}

	return max
}

func MaxSub(s string) string {
	n := 3
	max := ""
	//5257523
	for i := 0; i < len(s)-n; i++ {
		threeS := s[i : n+i]
		c := strings.Compare(threeS, max)
		if c == 1 {
			max = threeS
		}
	}
	return max
}
