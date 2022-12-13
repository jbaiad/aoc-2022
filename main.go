package main

import (
	"embed"
	"flag"
	"fmt"

	"github.com/jbaiad/aoc-2022/days"
)

//go:embed input/*
var input embed.FS

func main() {
	dayPtr := flag.Int("day", 1, "which day is being run")
	flag.Parse()
	fmt.Println("executing solution to problem from day", *dayPtr)

	switch *dayPtr {
	case 1:
		days.SolveDay1(input)
	case 2:
		fmt.Println("solving part 1")
		days.SolveDay2Part1(input)
		fmt.Println("solving part 2")
		days.SolveDay2Part2(input)
	case 3:
		fmt.Println("solving part 1")
		days.SolveDay3Part1(input)
		fmt.Println("solving part 2")
		days.SolveDay3Part2(input)
	case 4:
		fmt.Println("solving part 1")
		days.SolveDay4Part1(input)
		fmt.Println("solving part 2")
		days.SolveDay4Part2(input)
	default:
		panic("invalid day provided")
	}
}
