package day_one

import (
	"fmt"
	"strings"

	utils "github.com/alesharay/aoc_2023/config"
)

func Part_2() {
	day := "day1"
	filename := "sample2"

	old_lines := utils.Read_from_File(utils.Get_Full_Filename("input", day, filename))
	lines := []string {}
	
	start_old := 0
	end_old := 0 
	shares_letter := false
	should_break := false

	for _, line := range old_lines {
		new_line := line
		for _, word := range utils.Keys {
			if strings.Contains(new_line, word) {


				shares_letter = utils.Shares_Letter(
					new_line, word, start_old, end_old)

				start_old = strings.Index(new_line, word)
				end_old = start_old + len(word)



				switch shares_letter {
					case true:
						fmt.Printf("\nThere are numbers in %v that share a letter\n\n", line)
						should_break = true

					case false:
						// replace the word with it's number value
						new_line = strings.ReplaceAll(new_line, word, utils.Word_to_Num[word])
						lines = append(lines, new_line)
				}

				if should_break {
					break
				}
			} else {
				lines = append(lines, word)
			}
		}

		fmt.Printf("old line: %v\n", line)
		fmt.Printf("new line: %v\n\n", new_line)
	}

	Part_1(lines)
}