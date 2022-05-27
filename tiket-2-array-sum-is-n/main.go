package tiket_2_array_sum_is_n

func Find(arr []int, sum int) bool {
	arrMap := make(map[int]int, 0)

	for _, v := range arr {
		if _, ok := arrMap[v]; !ok {
			arrMap[v] = 0
		}
		arrMap[v] += 1
	}

	// sum = 10
	// 2-> 2
	// 1 -> 1

	for v, _ := range arrMap {
		find := sum - v

		if find == v {
			if c := arrMap[v]; c > 1 {
				return true
			}
		} else {
			if _, ok := arrMap[find]; ok {
				return true
			}
		}
	}

	return false
}
