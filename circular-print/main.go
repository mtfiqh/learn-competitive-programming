package main

import "fmt"

func doPrint(arr [100][100]int, width, height int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(arr[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}
func Print(width, height int) {
	type v struct {
		i int
		j int
	}

	visited := make(map[v]bool, 0)
	i := int(0)
	j := int(0)
	c := int(1)
	arr := [100][100]int{}
	directions := []string{
		"RIGHT",
		"BOTTOM",
		"LEFT",
		"TOP",
	}

	currDir := int(0)
	direction := directions[currDir]

	borderI := height - 1
	borderJ := width - 1

	for true {
		x := c
		_, ok := visited[v{i: i, j: j}]
		if ok {
			switch direction {
			case "TOP":
				i++
			case "BOTTOM":
				i--
			case "LEFT":
				j++
			case "RIGHT":
				j--
			}
			currDir++
			c++
		} else if direction == "RIGHT" && j == borderJ {
			currDir++
			c++
		} else if direction == "LEFT" && j == 0 {
			currDir++
			c++
		} else if direction == "TOP" && i == 0 {
			currDir++
			c++
		} else if direction == "BOTTOM" && i == borderI {
			currDir++
			c++
		}

		if currDir > 3 {
			currDir = 0
		}
		visited[v{i: i, j: j}] = true
		arr[i][j] = x

		direction = directions[currDir]
		switch direction {
		case "RIGHT":
			j++
		case "LEFT":
			j--
		case "TOP":
			i--
		case "BOTTOM":
			i++
		}
		doPrint(arr, width, height)

		if len(visited) == width*height {
			break
		}
	}
}

func main() {
	Print(3, 4)
}
