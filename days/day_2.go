package days

import (
	"bufio"
	"embed"
	"errors"
	"fmt"
	"strings"
)

func SolveDay2Part1(input embed.FS) (int, error) {
	f, err := input.Open("input/day-02.txt")
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	totalScore := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		directives := strings.Fields(line)
		if len(directives) != 2 {
			return -1, errors.New("bad input")
		}

		opponentToken, yourToken := directives[0], directives[1]
		var opponentState, yourState string
		switch opponentToken {
		case "A":
			opponentState = "Rock"
		case "B":
			opponentState = "Paper"
		case "C":
			opponentState = "Scissor"
		}

		score := 0
		switch yourToken {
		case "X":
			yourState = "Rock"
			score += 1
		case "Y":
			yourState = "Paper"
			score += 2
		case "Z":
			yourState = "Scissor"
			score += 3
		}

		if yourState == opponentState {
			score += 3
		} else if yourState == "Rock" && opponentState == "Scissor" {
			score += 6
		} else if yourState == "Paper" && opponentState == "Rock" {
			score += 6
		} else if yourState == "Scissor" && opponentState == "Paper" {
			score += 6
		}

		totalScore += score
	}

	fmt.Println("total score:", totalScore)
	return totalScore, nil
}

func SolveDay2Part2(input embed.FS) (int, error) {
	f, err := input.Open("input/day-02.txt")
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	totalScore := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		directives := strings.Fields(line)
		if len(directives) != 2 {
			return -1, errors.New("bad input")
		}

		opponentToken, yourToken := directives[0], directives[1]
		var opponentState, yourState string
		switch opponentToken {
		case "A":
			opponentState = "Rock"
		case "B":
			opponentState = "Paper"
		case "C":
			opponentState = "Scissor"
		}

		score := 0
		switch yourToken {
		case "X":
			if opponentState == "Rock" {
				score += 3
				yourState = "Scissor"
			} else if opponentState == "Paper" {
				score += 1
				yourState = "Rock"
			} else if opponentState == "Scissor" {
				score += 2
				yourState = "Paper"
			}
		case "Y":
			score += 3
			yourState = opponentState
			if yourState == "Rock" {
				score += 1
			} else if yourState == "Paper" {
				score += 2
			} else if yourState == "Scissor" {
				score += 3
			}
		case "Z":
			score += 6
			if opponentState == "Rock" {
				score += 2
				yourState = "Paper"
			} else if opponentState == "Paper" {
				score += 3
				yourState = "Scissor"
			} else if opponentState == "Scissor" {
				score += 1
				yourState = "Rock"
			}
		}
		totalScore += score
	}

	fmt.Println("total score:", totalScore)
	return totalScore, nil
}
