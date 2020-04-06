package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row)
	for i := range maze  {
		maze[i] = make([]int, col)
		for j := range maze[i]  {
			fmt.Fscanf(file,"%d", &maze[i][j])
		}
	}

	return maze
}


type point struct {
	i, j int
}

var dirs = [4]point {
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point)add(r point) point  {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	//fmt.Printf("====>>>%d,%d\n", p.i,p.j)
	return grid[p.i][p.j], true
}

func print(grid [][]int) {
	for _, row := range grid {
		for _, val := range row  {
			//if i == 0 && j == 0 {
			//	fmt.Printf("%2d ", 0)
			//	continue
			//}
			//if val == 0 {
			//	fmt.Printf("%2c ", ' ')
			//} else {
				fmt.Printf("%2d ", val)
			//}

		}
		fmt.Println()
	}
}

func walk(maze [][]int, start, end point) ([][]int, int, []point) {
	steps := make([][]int, len(maze))
	stepCount := -1
	for i := range maze  {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			stepCount, _ = cur.at(steps)
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)
		}
	}
	fmt.Println("---------------")
	fmt.Printf("总共要走%d步\n", stepCount)
	print(steps)
	fmt.Println("----以下为走法-----------")
	newsteps := make([][]int, len(maze))
	for i := range maze  {
		newsteps[i] = make([]int, len(maze[i]))
		print(newsteps)
	}

	for i, _ := range newsteps  {
		for j, _ := range newsteps[i]  {
			newsteps[i][j] = -1
		}
	}

	print(newsteps)

	stepStack := make([]point, 0)
	if stepCount != -1 {
		stepStack := append(stepStack, start)
		top := stepStack[len(stepStack) - 1]


		for top != end {

			topv, _ := top.at(steps)
			newsteps[top.i][top.j] = topv

			find := false;
			for _, dir := range dirs {
				next := top.add(dir)
				if next == start {
					continue
				}
				val, ok := next.at(steps)
				if !ok || val == 0 {
					continue
				}
				if val != topv + 1 {
					continue
				} else {
					stepStack = append(stepStack, next)
					find = true
				}
			}

			if !find {
				newsteps[top.i][top.j] = -1
				stepStack = stepStack[0:len(stepStack) - 1]
				top = stepStack[len(stepStack) - 1]
			} else {
				top = stepStack[len(stepStack) - 1]
				newsteps[top.i][top.j], _ = top.at(steps)

			}

			fmt.Println("---------------")
			fmt.Printf("第%d步\n", topv)
			print(newsteps)
		}
		

	}

	return steps, stepCount, stepStack
}

func main() {
	maze := readMaze("maze/maze.in")

	for _, row := range maze {
		for _, val := range row  {
			fmt.Printf("%2d ", val)
		}
		fmt.Println()
	}

	steps, stepCount, stpps:= walk(maze, point{0,0}, point{len(maze)-1, len(maze[0]) - 1})

	fmt.Println("---------------")
	fmt.Printf("总共要走%d步\n", stepCount)
	print(steps)

	fmt.Print("%v", stpps)

	//for p := range stpps {
	//	fmt.Printf("总共要走(%d,%d)步\n", p.i, p.j)
	//}
}
