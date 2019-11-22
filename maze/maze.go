package main

import (
	"fmt"
	"os"
)

func main() {
	maze := readMaze("maze/maze.in")
	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d", val)
		}
		fmt.Println()
	}

	//walk the map
	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}

	steps := walk(maze, start, end)

	//print steps
	for _, step := range steps {
		for _, val := range step {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

/**
  point struct
*/
type point struct {
	i, j int
}

//define 4 directions delta position value
var dirs = [4]point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

/**
  move the point
*/
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

/**
check point is out of range of map,
 and return the value at the point of map
*/
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.j < 0 || p.i >= len(grid) || p.j >= len(grid[0]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

/**
  walk the maze
  return steps
*/
func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}
	fmt.Println("steps", steps)
	q := []point{start}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		//check is last
		if cur == end {
			fmt.Println("cur == end")
			break
		}
		for _, dir := range dirs {
			next := cur.add(dir)
			fmt.Println("next:", next)

			//check is 1 or out of range of maze ,
			n, ok := next.at(maze)
			if !ok || n == 1 {
				continue
			}

			//check > 0 or out of range of steps map
			n, ok = next.at(steps)
			if !ok || n > 0 {
				continue
			}

			//check is start
			if next == start {
				continue
			}

			//save next step into step map
			curStep := steps[cur.i][cur.j]
			nextStep := curStep + 1
			steps[next.i][next.j] = nextStep

			//append next to queue
			q = append(q, next)
			fmt.Println("q", q, " next", next)

		}

	}
	return steps

}

/**
  read maze from file
*/
func readMaze(s string) [][]int {
	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}

	//get number of  column and row
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col)

	//define maze
	maze := make([][]int, row) // 定义行数
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze

}
