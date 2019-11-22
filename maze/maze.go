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

	//走迷宫
	start := point{0, 0}
	end := point{len(maze) - 1, len(maze[0]) - 1}

	steps := walk(maze, start, end)

	//打印steps
	for _, step := range steps {
		for _, val := range step {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

/**
  点的结构体
*/
type point struct {
	i, j int
}

//定义四个方向的节点
var dirs = [4]point{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

/**
节点的添加
*/
func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

/**
检查节点是否越界，
    如果不越界，在一个二维图的值是多少
*/
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.j < 0 || p.i >= len(grid) || p.j >= len(grid[0]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

/**point
走迷宫
*/
func walk(maze [][]int, start, end point) [][]int {

	// 路径数组,等行等列的，全部设置成0的二维数组
	steps := make([][]int, len(maze))
	for i := range maze {
		steps[i] = make([]int, len(maze[0]))
	}

	//将起点加入到队列中去
	q := []point{start}

	//循环，当队列不为空时一直执行
	for len(q) > 0 {
		//将当前位置设置为队列的第一个节点
		cur := q[0]
		q = q[1:]
		//发现当前节点是终点，执行退出
		if cur == end {
			fmt.Println("end!")
			break
		}
		//探索当前节点的：上，下，左，右 四个方向的 相邻节点
		for _, dir := range dirs {
			//设置下一个节点是方向偏移量加上当前节点的
			next := cur.add(dir)

			//在maze中越界或者不能走的 , 跳过，处理下一个节点
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			//在steps中越界或者已经走过的， 跳过，处理下一个
			val, ok = next.at(steps)
			if !ok || val > 0 {
				continue
			}

			//当前点等于原点的，跳过
			if next == start {
				continue
			}

			//将step填好访问到的点，下一步节点的数量等于当前节点的数量+1
			curStep, _ := cur.at(steps)
			fmt.Println(curStep + 1)
			steps[next.i][next.j] = curStep + 1

			//将下一个节点追加到队列中
			q = append(q, next)
		}
	}

	return steps

}

/**
  从文件读取迷宫
*/
func readMaze(s string) [][]int {
	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}

	//获得第行数和列数
	var row, col int
	_, _ = fmt.Fscanf(file, "%d %d", &row, &col)

	//定义二维数组存储数据
	maze := make([][]int, row) // 定义行数
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze

}
