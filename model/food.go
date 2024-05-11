package model

import (
	"math/rand"
)

type Point struct {
	X int
	Y int
}

func GenerateFood(snake *Snake, maxX int, maxY int) Point {
	validPositions := make([]Point, 0, maxX*maxY-len(snake.Body.Nodes))

	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			valid := true
			for _, bp := range snake.Body.Nodes {
				if bp.Value.X == x && bp.Value.Y == y {
					valid = false
					break
				}
			}
			if valid {
				validPositions = append(validPositions, Point{x, y})
			}
		}
	}

	return validPositions[rand.Intn(len(validPositions))]
}
