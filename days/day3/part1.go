package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	utils "github.com/alesharay/aoc_2023/config"
)


func Part_1(lines []string) {
	day := "day2"
	filename := "part1"

	full_input_filename := utils.Get_Full_Filename("input", day, filename)
	if len(lines) == 0 {
		lines = utils.Read_from_File(full_input_filename)
	}

	full_output_filename := utils.Get_Full_Filename("output", day, filename)
	out_e := os.Remove(full_output_filename)
	utils.Check(out_e)

	
}