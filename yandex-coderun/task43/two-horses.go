package main

import "fmt"

type Position struct {
	x, y int
}

var moves = []Position{
	{2, 1}, {2, -1}, {-2, 1}, {-2, -1},
	{1, 2}, {1, -2}, {-1, 2}, {-1, -2},
}

func parsePosition(pos string) Position {
	return Position{int(pos[0] - 'a'), int(pos[1] - '1')}
}

func isValid(x, y int) bool {
	return x >= 0 && x < 8 && y >= 0 && y < 8
}

func minMovesToMeet(red, green Position) int {
	if red == green {
		return 0
	}

	visited := make(map[string]bool)
	queue := []struct {
		red, green Position
		steps      int
	}{{red, green, 0}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, moveRed := range moves {
			newRed := Position{current.red.x + moveRed.x, current.red.y + moveRed.y}
			if !isValid(newRed.x, newRed.y) {
				continue
			}

			for _, moveGreen := range moves {
				newGreen := Position{current.green.x + moveGreen.x, current.green.y + moveGreen.y}
				if !isValid(newGreen.x, newGreen.y) {
					continue
				}

				if newRed == newGreen {
					return current.steps + 1
				}

				key := fmt.Sprintf("%d,%d,%d,%d", newRed.x, newRed.y, newGreen.x, newGreen.y)
				if !visited[key] {
					visited[key] = true
					queue = append(queue, struct {
						red, green Position
						steps      int
					}{newRed, newGreen, current.steps + 1})
				}
			}
		}
	}

	return -1
}

func main() {
	var redPos, greenPos string
	fmt.Scanf("%s %s", &redPos, &greenPos)

	red := parsePosition(redPos)
	green := parsePosition(greenPos)

	result := minMovesToMeet(red, green)
	fmt.Println(result)
}
