package main

import (
	"fmt"
)

func main() {
	var prev, current int
	var sequenceType string

	fmt.Scan(&prev)

	sequenceType = "CONSTANT"

	for {
		fmt.Scan(&current)

		if current == -2000000000 {
			break
		}

		if sequenceType == "CONSTANT" {
			if current > prev {
				sequenceType = "ASCENDING"
			} else if current < prev {
				sequenceType = "DESCENDING"
			}
		} else if sequenceType == "ASCENDING" {
			if current < prev {
				sequenceType = "RANDOM"
			} else if current == prev {
				sequenceType = "WEAKLY ASCENDING"
			}
		} else if sequenceType == "WEAKLY ASCENDING" {
			if current < prev {
				sequenceType = "RANDOM"
			}
		} else if sequenceType == "DESCENDING" {
			if current > prev {
				sequenceType = "RANDOM"
			} else if current == prev {
				sequenceType = "WEAKLY DESCENDING"
			}
		} else if sequenceType == "WEAKLY DESCENDING" {
			if current > prev {
				sequenceType = "RANDOM"
			}
		}

		prev = current
	}

	fmt.Println(sequenceType)
}
