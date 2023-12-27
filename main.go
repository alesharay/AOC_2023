package main

import (
	"fmt"
	"strings"

	"github.com/alesharay/aoc_2023/days/day1"
	"github.com/alesharay/aoc_2023/days/day2"
)


func main() {
	fmt.Println()

	fmt.Printf("Day 1:\n\n")

	fmt.Printf("  Part 1:\n")
	day1.Part_1(strings.Split("", ""))
	fmt.Println()

	fmt.Printf("  Part 2:\n")
	day1.Part_2()
	fmt.Println("===========================================\n\n")

	fmt.Printf("Day 2:\n\n")


	fmt.Printf("  Part 1:\n")
	day2.Part_1(strings.Split("", ""))
	fmt.Println()

	fmt.Printf("  Part 2:\n")
	day2.Part_2(strings.Split("", ""))
	fmt.Println("===========================================\n\n")
}