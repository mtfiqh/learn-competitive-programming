// https://www.hackerrank.com/challenges/sparse-arrays/problem?isFullScreen=true&h_r=next-challenge&h_v=zen#!
package main

import "fmt"

func main() {
	arr := make(map[string]int, 0)
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if _, ok := arr[s]; !ok {
			arr[s] = 1
			continue
		}
		arr[s]++
	}

	var m int
	fmt.Scan(&m)

	for i := 0; i < m; i++ {
		var q string
		fmt.Scan(&q)
		fmt.Println(arr[q])

	}
}
