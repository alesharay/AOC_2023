package day1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/alesharay/aoc_2023/config"
)


func Part_1(lines []string) {
	day := "day1"
	filename := "part1"

	if len(lines) == 0 {
		lines = utils.Read_from_File(utils.Get_Full_Filename("input", day, filename))
	}

	alpha := regexp.MustCompile(`[a-zA-z]`)
	sum := 0

	for _, line := range lines {

		num := 0
		only_nums := alpha.ReplaceAllString(line, "")

		length := len(only_nums)
		switch length {
			case 0:
				continue
			case 1:
				num_string := strings.Repeat(only_nums,2)
				num, _ = strconv.Atoi(num_string)
			default:
				num_string := string(only_nums[0]) + string(only_nums[len(only_nums)-1])
				num, _ = strconv.Atoi(num_string)
		}

		sum += num;
	}

	fmt.Printf("---- Sum: %v\n\n", sum)
}