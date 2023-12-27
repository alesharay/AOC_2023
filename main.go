package main

import (
	"fmt"
	"strings"

	"github.com/alesharay/aoc_2023/internal/day_one"
	"github.com/alesharay/aoc_2023/internal/day_two"
)


func main() {
	fmt.Println()

	fmt.Printf("Day 1:\n\n")

	fmt.Printf("  Part 1:\n")
	day_one.Part_1(strings.Split("", ""))
	fmt.Println()

	fmt.Printf("  Part 2:\n")
	day_one.Part_2()
	fmt.Println("===========================================\n\n")

	fmt.Printf("Day 2:\n\n")


	fmt.Printf("  Part 1:\n")
	day_two.Part_1(strings.Split("", ""))
	fmt.Println()

	fmt.Printf("  Part 2:\n")
	day_two.Part_2(strings.Split("", ""))
	fmt.Println("===========================================\n\n")
}