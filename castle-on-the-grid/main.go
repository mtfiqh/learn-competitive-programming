package main

import "fmt"

const (
	RIGHT  = "right"
	LEFT   = "left"
	TOP    = "top"
	BOTTOM = "bottom"
)

const (
	NOT_VISITED = 0
	VISITED     = 1
	CHECKED     = 2
	BLOCKED     = 3
)

type Node struct {
	Next  map[string]*Node
	Prev  *Node
	Count int
}

type Data struct {
	X int
	Y int
	C int
}

type Runway struct {
	First *Node
	Last  *Node
}

func main() {
	var n int
	fmt.Scan(&n)
	grid := make([]string, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		grid[i] = s
	}
	var startX int
	fmt.Scan(&startX)
	var startY int
	fmt.Scan(&startY)
	var goalX int
	fmt.Scan(&goalX)
	var goalY int
	fmt.Scan(&goalY)

	run(grid, Data{
		X: startX,
		Y: startY,
	}, Data{
		X: goalX,
		Y: goalY,
	})
}

func insertHistory(history *Node, data Data, way string) {
	if history == nil {
		history.Next[way] = &Node{
			Next: nil,
			Prev: history,
		}
	}

	tail := findTail(history, way)
	tail.Next[way] = &Node{
		Next: nil,
		Prev: tail,
	}
}

func findTail(node *Node, way string) *Node {
	if node.Next[way] == nil {
		return node
	}

	return findTail(node.Next[way], way)
}

func run(grid []string, start, goal Data) {
	curr := start

	qs := []Data{
		{
			X: curr.X,
			Y: curr.Y,
		},
	}

	wayToGo := []string{
		RIGHT,
		TOP,
		BOTTOM,
		LEFT,
	}

	g := convertToSlice(grid)
	lenQs := len(qs)
	count := 0

	for k := 0; k < lenQs; k++ {
		count = qs[k].C
		for _, way := range wayToGo {
			qq := qs[k]

			for check(grid, qq) {
				if g[qq.X][qq.Y].V == NOT_VISITED {
					g[qq.X][qq.Y].V = VISITED
					qq.C = count + 1
					qs = append(qs, qq)
					lenQs++
				}

				if qq.X == goal.X && qq.Y == goal.Y {
					fmt.Println(qq.C)
					return
				}

				//fmt.Println("now: ", qq)
				//fmt.Println("queue: ", qs, " | in key: ", k, "| c: ", qq.C)
				//fmt.Println("GO TO", way)
				Print(g, len(grid))
				qq = goNext(qq, way)
				fmt.Println()

			}
		}
		count++

	}
}

func goNext(curr Data, way string) Data {
	if way == TOP {
		curr.X -= 1
	}
	if way == BOTTOM {
		curr.X += 1
	}
	if way == RIGHT {
		curr.Y += 1
	}
	if way == LEFT {
		curr.Y -= 1
	}

	return curr
}

type GridSlice struct {
	V int
}

func convertToSlice(grid []string) [100][100]GridSlice {
	var gridInSlice [100][100]GridSlice
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == 'X' {
				gridInSlice[i][j].V = BLOCKED
			} else {
				gridInSlice[i][j].V = NOT_VISITED
			}
		}
	}

	//	print

	return gridInSlice
}

func Print(gridInSlice [100][100]GridSlice, length int) {
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			fmt.Print(gridInSlice[i][j].V, "\t")
		}
		fmt.Println()
	}
}

func check(grid []string, point Data) bool {
	n := len(grid)
	if point.X < 0 || point.X >= n || point.Y < 0 || point.Y >= n {
		return false
	}
	if grid[point.X][point.Y] == 'X' {
		return false
	}

	return true
}

//func check(grid []string, currX, currY int, direction string, n int) bool {
//	if direction == TOP {
//		currX -= 1
//	}
//	if direction == BOTTOM {
//		currX += 1
//	}
//	if direction == RIGHT {
//		currY += 1
//	}
//	if direction == LEFT {
//		currY += 1
//	}
//
//	if currX < 0 || currX >= n || currY < 0 || currY >= n {
//		return false
//	}
//
//	return true
//}
