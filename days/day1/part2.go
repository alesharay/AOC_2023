package day1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/alesharay/aoc_2023/config"
)

func Part_2() {
	sum := 0
	num_reg := regexp.MustCompile(`[123456789]`)
	day := "day1"
	filename := "part1"

	full_input_filename := utils.Get_Full_Filename("input", day, filename)
	lines := utils.Read_from_File(full_input_filename)

	full_output_filename := utils.Get_Full_Filename("output", day, filename)
	out_e := os.Remove(full_output_filename)
	utils.Check(out_e)

	for _, line := range lines {

		if line == "" {
			utils.Write_to_file(full_output_filename, "")
			continue
		}

		num_string := ""

// ------------------------------------------------------------------------------------------
// CALCULATE FIRST NUMBER
// ------------------------------------------------------------------------------------------
		first_num := num_reg.FindAllStringIndex(line, -1)
		first_num_index := 999

		if len(first_num) > 0 {
			first_num_index = first_num[0][0]
		}

		first_word_index, first_word := utils.Word_Values(line, false)

		if 	first_word_index == -1 ||
				first_num_index < first_word_index {
			num_string += string(line[first_num_index])
		} else {
			num_string += strconv.Itoa(utils.Word_to_Num[first_word])
		}


// ------------------------------------------------------------------------------------------
// CALCULATE LAST NUMBER
// ------------------------------------------------------------------------------------------
		reversed_line := utils.Reverse_String(line)

		last_num_slice := num_reg.FindAllStringIndex(line, -1)
		last_num_index := -1
		last_num := ""

		if len(last_num_slice) > 0 {
			last_num_index = last_num_slice[len(last_num_slice)-1][0]
			last_num = strings.Split(line, "")[last_num_index]
		}

		last_word_index, last_word := utils.Word_Values(reversed_line, true)

		if last_word_index != -1 {
			reverse_last_word := utils.Reverse_String(last_word)
			line_length := len(line)
			reverse_last_word_index := (line_length) - last_word_index - len(reverse_last_word)

			if last_num_index > reverse_last_word_index {
				num_string += string(last_num)
			} else {
				num_string += strconv.Itoa(utils.Word_to_Num[reverse_last_word])
			}
		} else {
			num_string += string(last_num)
		}


// ------------------------------------------------------------------------------------------
// CHECK IF THERE IS ONLY ONE NUMBER
// ------------------------------------------------------------------------------------------
		first := string(num_string[0])
		last := string(num_string[1])

		if first == "0" || last == "0" {
			if first == "0" {
				num_string = strings.Repeat(first, 2)
			} else {
				num_string = strings.Repeat(last, 2)
			}
		}

		utils.Write_to_file(full_output_filename, num_string)

		num, _ := strconv.Atoi(num_string)
		sum += num;
	}
// ------------------------------------------------------------------------------------------

	fmt.Printf("---- Sum: %v\n\n", sum)
}