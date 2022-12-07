package days

import (
	"bufio"
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SolveDay1(input embed.FS) ([]int, error) {
	f, err := input.Open("input/day-01.txt")
	if err != nil {
		fmt.Println(err)
		return []int{}, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	caloriesCarried := []int{}
	currentCaloriesCarried := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			caloriesCarried = append(caloriesCarried, currentCaloriesCarried)
			currentCaloriesCarried = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			return []int{}, err
		}
		currentCaloriesCarried += calories
	}
	caloriesCarried = append(caloriesCarried, currentCaloriesCarried)

	sort.Ints(caloriesCarried)
	fmt.Println("top 3 calories carried", caloriesCarried[len(caloriesCarried)-3:])
	return caloriesCarried, nil
}
