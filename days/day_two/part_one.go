package day_two

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "github.com/alesharay/aoc_2023/config"
)


func Part_1(lines []string) {
	day := "day_two"
	filename := "part1"

	full_input_filename := utils.Get_Full_Filename("input", day, filename)
	if len(lines) == 0 {
		lines = utils.Read_from_File(full_input_filename)
	}

	full_output_filename := utils.Get_Full_Filename("output", day, filename)
	out_e := os.Remove(full_output_filename)
	utils.Check(out_e)

	max_red, max_blue, max_green := 12, 14, 13
	game_is_possible := true
	sum := 0

	for _, line := range lines {
		if line == "" {
			utils.Write_to_file(full_output_filename, "")
			continue
		}

		games := strings.Split(line, ":")
		id, _ := strconv.Atoi(strings.Split(games[0], " ")[1])

		rounds := strings.Split(games[1], ";")

		for _, round := range rounds {
			colors_options := strings.Split(round, ",")

			for _, colors := range colors_options {
				count, _ := strconv.Atoi(strings.Split(colors, " ")[1])
				color := strings.Split(colors, " ")[2]

				switch color {
					case "blue":
						game_is_possible = utils.Is_Game_Possible(count, max_blue)
					case "red":
						game_is_possible = utils.Is_Game_Possible(count, max_red)
					case "green":
						game_is_possible = utils.Is_Game_Possible(count, max_green)
				}

				if !game_is_possible {
					break
				}
			}

			if !game_is_possible {
				break
			}
		}

		if game_is_possible {
			utils.Write_to_file(full_output_filename, strconv.Itoa(id))
			sum += id
		}
	}

	fmt.Printf("---- Sum: %v\n\n", sum)
}