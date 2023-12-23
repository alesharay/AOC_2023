package day_one

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/alesharay/aoc_2023/config"
)

func Part_2() {
	fmt.Printf("\nDay 1 Part 2:\n\n")

	input_filename := "config/input/day1/part1.txt"
	lines := config.Read_from_File(input_filename)

	num_reg := regexp.MustCompile(`[123456789]`)
	sum := 0

	for _, line := range lines {
		num_string := ""
// ------------------------------------------------------------------------------------------
		first_num := num_reg.FindAllStringIndex(line, -1)
		first_num_index := 999

		if len(first_num) > 0 {
			first_num_index = first_num[0][0]
		}

		first_word_index, first_word := config.Word_Values(line, "first")

<<<<<<< HEAD
=======
		// fmt.Printf(" ==== first #: %v ; first # ind: %v\n", string(line[first_num_index]), first_num_index)
		// fmt.Printf("\n ==== first word: %v ; first word ind: %v\n", first_word, first_word_index)

>>>>>>> 223f4eebf3efcf7a04dacbd5ca4ef9b7b8bd767d
		if first_num_index < first_word_index {
			num_string += string(line[first_num_index])
		} else {
			num_string += strconv.Itoa(config.Convert_Word_to_Num(first_word))
		}
// ------------------------------------------------------------------------------------------
		last_num_slice := num_reg.FindAllStringIndex(line, -1)
		last_num_index := -1
		last_num := ""

		if len(last_num_slice) > 0 {
			last_num_index = last_num_slice[len(last_num_slice)-1][0]
			last_num = strings.Split(line, "")[last_num_index]
		}

		last_word_index, last_word := config.Word_Values(line, "last")

<<<<<<< HEAD
=======
		fmt.Printf(" ==== last #: %v ; last # ind: %v\n", string(last_num), last_num_index)
		fmt.Printf("\n ==== last word: %v ; last word ind: %v\n", last_word, last_word_index)

>>>>>>> 223f4eebf3efcf7a04dacbd5ca4ef9b7b8bd767d
		if last_num_index > last_word_index {
			num_string += string(last_num)
		}else {
			num_string += strconv.Itoa(config.Convert_Word_to_Num(last_word))
		}

<<<<<<< HEAD


		if string(num_string[0]) == "0" || string(num_string[1]) == "0" {
			if string(num_string[0]) == "0" {
				num_string = strings.Repeat(string(num_string[1]),2)
			} else {
				num_string = strings.Repeat(string(num_string[0]),2)
			}
		}

		config.Write_to_file("config/output/day1/part1.txt", num_string)

		fmt.Printf("\nNumber: %v\n", num_string)
		fmt.Println("len of file:", len(lines))
=======
		if string(num_string[0]) == "0" {
			num_string = strings.Repeat(string(num_string[1]),2)
		}
		if string(num_string[1]) == "0" {
			num_string = strings.Repeat(string(num_string[0]),2)
		}

		io = "output"
		full_filename = config.Get_Full_Filename(io, day, filename)

		config.Write_to_file(full_filename, num_string)

		fmt.Printf("\nNumber: %v\n\n", num_string)
>>>>>>> 223f4eebf3efcf7a04dacbd5ca4ef9b7b8bd767d

		num, _ := strconv.Atoi(num_string)
		sum += num;
	}
// ------------------------------------------------------------------------------------------
	fmt.Printf("---- Sum: %v\n\n", sum)
<<<<<<< HEAD
}
=======
}
>>>>>>> 223f4eebf3efcf7a04dacbd5ca4ef9b7b8bd767d
