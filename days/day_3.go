package days

import (
	"bufio"
	"embed"
	"fmt"
	"math"
	"strings"
)

func calculatePriority(item rune) int {
	var priority int
	if 65 <= int(item) && int(item) <= 90 {
		priority = int(item) - 38
	} else {
		priority = int(item) - 96
	}

	if calculateCharacter(priority) != item {
		panic("incorrect inversion")
	}
	return priority
}

func calculateCharacter(priority int) rune {
	if 1 <= priority && priority <= 26 {
		return rune(priority + 96)
	} else {
		return rune(priority + 38)
	}
}

func SolveDay3Part1(input embed.FS) (int, error) {
	f, err := input.Open("input/day-03.txt")
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	/* Parse the file line by line:
	Each line represents the contents of a rucksack, which is divided into two
	compartments. You can identify the contents of the first compartment by
	looking at the left half of the string, and the second compartment
	by the right half of the string. An item within a compartment is a character
	matching the regex [a-zA-Z] within the substring.
	*/
	priority_sum := 0
	line_num := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		left := map[rune]bool{}
		right := map[rune]bool{}

		// Process each rucksack
		for i, item := range line {
			if i < len(line)/2 {
				left[item] = true
			} else {
				right[item] = true
			}

			if !left[item] || !right[item] {
				continue
			}

			priority_sum += calculatePriority(item)
			break
		}
		line_num += 1
	}

	fmt.Printf("priority sum = %d\n", priority_sum)
	return priority_sum, nil
}

func SolveDay3Part2(input embed.FS) (int, error) {
	f, err := input.Open("input/day-03.txt")
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	/* Parse the file line by line:
	Each line represents the contents of a rucksack, which is divided into two
	compartments. You can identify the contents of the first compartment by
	looking at the left half of the string, and the second compartment
	by the right half of the string. An item within a compartment is a character
	matching the regex [a-zA-Z] within the substring.
	*/
	priority_sum := 0
	line_num := 0
	group_num := 1
	for scanner.Scan() {
		fmt.Println("---")
		fmt.Printf("group: %d\n", group_num)
		fmt.Printf("beginning at line: %d\n", line_num)

		// Process each rucksack into an integer between 0 and 2^52 - 1
		// that uniquely encodes the rucksack's contents. A rucksack's
		// corresponding integer has a 1 in the p-th most significant bit
		// if it contains the item with priority p + 1.
		elf_1 := strings.TrimSpace(scanner.Text())
		elf_1_rucksack_num := uint(0)
		for _, item := range elf_1 {
			priority := calculatePriority(item)
			priorityEncoding := uint(1 << (priority - 1))
			fmt.Printf("\t...adding 2^(%d - 1) = %d to rucksack num since %s has priority %d\n", priority, priorityEncoding, string(item), priority)
			elf_1_rucksack_num += priorityEncoding
		}
		fmt.Printf("\tElf 1's rucksack string: %s\n", elf_1)
		fmt.Printf("\tElf 1's rucksack number (decimal): %d\n", elf_1_rucksack_num)

		if !scanner.Scan() {
			panic("BAD INPUT!")
		}
		elf_2 := strings.TrimSpace(scanner.Text())
		elf_2_rucksack_num := uint(0)
		for _, item := range elf_2 {
			priority := calculatePriority(item)
			priorityEncoding := uint(1 << (priority - 1))
			fmt.Printf("\t...adding 2^(%d - 1) = %d to rucksack num since %s has priority %d\n", priority, priorityEncoding, string(item), priority)
			elf_2_rucksack_num += priorityEncoding
		}
		fmt.Printf("\tElf 2's rucksack string: %s\n", elf_2)
		fmt.Printf("\tElf 2's rucksack number (decimal): %d\n", elf_2_rucksack_num)

		if !scanner.Scan() {
			panic("BAD INPUT!")
		}
		elf_3 := strings.TrimSpace(scanner.Text())
		elf_3_rucksack_num := uint(0)
		for _, item := range elf_3 {
			priority := calculatePriority(item)
			priorityEncoding := uint(1 << (priority - 1))
			fmt.Printf("\t...adding 2^(%d - 1) = %d to rucksack num since %s has priority %d\n", priority, priorityEncoding, string(item), priority)
			elf_3_rucksack_num += priorityEncoding
		}
		fmt.Printf("\tElf 3's rucksack string: %s\n", elf_3)
		fmt.Printf("\tElf 3's rucksack number (decimal): %d\n", elf_3_rucksack_num)

		intersection := elf_1_rucksack_num & elf_2_rucksack_num & elf_3_rucksack_num
		priority := int(math.Log2(float64(intersection))) + 1
		fmt.Printf("\tintersection = %d, common priority (decimal) = %d, common item (character) = %s\n", intersection, priority, string(calculateCharacter(priority)))

		group_num += 1
		line_num += 3
	}

	return priority_sum, nil
}
