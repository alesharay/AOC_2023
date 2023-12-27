package day_two

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "github.com/alesharay/aoc_2023/config"
)


func Part_2(lines []string) {
	day := "day2"
	filename := "part1"

	full_input_filename := utils.Get_Full_Filename("input", day, filename)
	if len(lines) == 0 {
		lines = utils.Read_from_File(full_input_filename)
	}

	full_output_filename := utils.Get_Full_Filename("output", day, filename)
	out_e := os.Remove(full_output_filename)
	utils.Check(out_e)

	sum := 0

	for _, line := range lines {
		if line == "" {
			utils.Write_to_file(full_output_filename, "")
			continue
		}

		max_red, max_blue, max_green := 0, 0, 0
		games := strings.Split(line, ":")
		rounds := strings.Split(games[1], ";")

		for _, round := range rounds {
			colors_options := strings.Split(round, ",")

			for _, colors := range colors_options {
				count, _ := strconv.Atoi(strings.Split(colors, " ")[1])
				color := strings.Split(colors, " ")[2]

				switch color {
					case "blue":
						if count > max_blue {
							max_blue = count
						}
					case "red":
						if count > max_red {
							max_red = count
						}
					case "green":
						if count > max_green {
							max_green = count
						}
				}
			}
		}

		power := max_blue * max_green * max_red
		sum += power
		utils.Write_to_file(full_output_filename, strconv.Itoa(power))
	}

	fmt.Printf("---- Sum: %v\n\n", sum)
}