package days

import (
	"bufio"
	"embed"
	"fmt"
	"strconv"
	"strings"
)

func SolveDay4Part1(input embed.FS) (int, error) {
	f, err := input.Open("input/day-04.txt")
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	pairs_fully_contained := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		ranges := strings.Split(line, ",")
		first_range, second_range := ranges[0], ranges[1]

		first_bounds := strings.Split(first_range, "-")
		l0_str, l1_str := first_bounds[0], first_bounds[1]
		l0, err := strconv.Atoi(l0_str)
		if err != nil {
			return -1, err
		}
		l1, err := strconv.Atoi(l1_str)
		if err != nil {
			return -1, err
		}

		second_bounds := strings.Split(second_range, "-")
		r0_str, r1_str := second_bounds[0], second_bounds[1]
		r0, err := strconv.Atoi(r0_str)
		if err != nil {
			return -1, err
		}
		r1, err := strconv.Atoi(r1_str)
		if err != nil {
			return -1, err
		}

		if r0 <= l0 && l1 <= r1 {
			// L < R: first range is a subset of second
			fmt.Printf("1 | %s contained by %s: %d <= %d <= %d <= %d\n", first_range, second_range, r0, l0, l1, r1)
			pairs_fully_contained += 1
		} else if l0 <= r0 && r1 <= l1 {
			// L > R: second range is subset of first
			fmt.Printf("2 | %s contained by %s: %d <= %d <= %d <= %d\n", second_range, first_range, l0, r0, r1, l1)
			pairs_fully_contained += 1
		}
	}

	fmt.Printf("pairs where one range is fully contained by the other = %d\n", pairs_fully_contained)
	return pairs_fully_contained, nil
}

func SolveDay4Part2(input embed.FS) (int, error) {
	f, err := input.Open("input/day-04.txt")
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	pairs_partially_contained := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		ranges := strings.Split(line, ",")
		first_range, second_range := ranges[0], ranges[1]

		first_bounds := strings.Split(first_range, "-")
		l0_str, l1_str := first_bounds[0], first_bounds[1]
		l0, err := strconv.Atoi(l0_str)
		if err != nil {
			return -1, err
		}
		l1, err := strconv.Atoi(l1_str)
		if err != nil {
			return -1, err
		}

		second_bounds := strings.Split(second_range, "-")
		r0_str, r1_str := second_bounds[0], second_bounds[1]
		r0, err := strconv.Atoi(r0_str)
		if err != nil {
			return -1, err
		}
		r1, err := strconv.Atoi(r1_str)
		if err != nil {
			return -1, err
		}

		if r0 <= l1 && l1 <= r1 {
			fmt.Printf("%s overlaps with %s: %d <= %d <= %d\n", first_range, second_range, r0, l1, r1)
			pairs_partially_contained += 1
		} else if l0 <= r1 && r1 <= l1 {
			fmt.Printf("%s overlaps with %s: %d <= %d <= %d\n", first_range, second_range, l0, r1, l1)
			pairs_partially_contained += 1
		}
	}

	fmt.Printf("pairs where one range is partially contained by the other = %d\n", pairs_partially_contained)
	return pairs_partially_contained, nil
}
