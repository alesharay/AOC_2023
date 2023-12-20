package day_one

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/alesharay/aoc_2023/config"

)


func Part_1() {
	fmt.Printf("\nDay 1 Part 1:\n\n")

	filename := "part1"
	day := "day1"
	io := "input"
	full_filename := config.Get_Full_Filename(io, day, filename)

	fmt.Printf("\n===== Full Filename: %v\n", full_filename)

	input, err := os.ReadFile(full_filename)
	config.Check(err)
	lines := strings.Split(string(input), "\n")

	alpha := regexp.MustCompile(`[a-zA-z]`)
	sum := 0

	for _, line := range lines {
		num := 0
		only_nums := alpha.ReplaceAllString(line, "")

		if len(only_nums) == 1 {
			num_string := strings.Repeat(only_nums,2)
			num, _ = strconv.Atoi(num_string)
		} else {
			num_string := string(only_nums[0]) + string(only_nums[len(only_nums)-1])
			num, _ = strconv.Atoi(num_string)
		}

		sum += num;
	}

	fmt.Printf("---- Sum: %v\n\n", sum)
}