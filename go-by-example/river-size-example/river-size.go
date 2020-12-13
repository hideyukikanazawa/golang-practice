package main

import (
	"fmt"
)

type coord struct {
	y int
	x int
}

func getRiverSizes(matrix [][]int) []int {
	riverSizes := make([]int, 0)
	marked := make(map[coord]struct{}, 0)
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if _, ok := marked[coord{row, col}]; matrix[row][col] == 1 && !ok {
				marked[coord{row, col}] = struct{}{}
				riverSize := 1
				nodes := []coord{coord{row, col}}
				for len(nodes) > 0 {
					current := nodes[0]
					nodes = append(nodes[1:])
					neighbours := getNeighbours(current, matrix)
					for _, nb := range neighbours {
						if _, ok := marked[nb]; matrix[nb.y][nb.x] == 1 && !ok {
							nodes = append(nodes, nb)
							marked[nb] = struct{}{}
							riverSize++
						}
					}
				}
				riverSizes = append(riverSizes, riverSize)
			}
		}
	}
	return riverSizes
}

func getNeighbours(c coord, m [][]int) []coord{
	ns := []coord{}
	if c.y < len(m) - 1 {
		ns = append(ns, coord{c.y + 1, c.x})
	}
	if c.y > 0 {
		ns = append(ns, coord{c.y - 1, c.x})
	}
	if c.x < len(m[c.y]) - 1 {
		ns = append(ns, coord{c.y, c.x + 1})
	}
	if c.x > 0 {
		ns = append(ns, coord{c.y, c.x - 1})
	}
	return ns
}

func main() {
	matrix := [][]int{
		[]int{1, 0, 0, 0, 0},
		[]int{1, 0, 1, 0, 1},
		[]int{1, 0, 1, 0, 1},
		[]int{1, 0, 1, 0, 0},
		[]int{0, 1, 0, 1, 0},
	}
	riverSizes := getRiverSizes(matrix)
	fmt.Printf("The river sizes in matrix are %v", riverSizes)
}