package config

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	Word_to_Num = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	Keys = strings.Split("one two three four five six seven eight nine", " ")
)

type Word_Index struct {
	Word string
	Index int
	Exists bool
}

func Sort_Map(input map[int]string) ([]Word_Index) {
	keys := []int{}
	result := make([]Word_Index, len(input))

	for k := range input {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for i, key := range keys {
		result[i].Index = key
		result[i].Word = input[key]
		result[i].Exists = true
	}

	return result
}

func Word_Index_Func(
	line string,
	reverse bool,
) Word_Index {

	index_map := map[int]string{}
	contains := false

	for _, word := range Keys {

		if reverse {
			word = Reverse_String(word)
		}

		if strings.Contains(line, word) {
			contains = true
			index := strings.Index(line, word)
			index_map[index] = word
		}
	}

	if contains {
		result_slice := Sort_Map(index_map)

		return Word_Index{
			Word: result_slice[0].Word,
			Index: result_slice[0].Index,
			Exists: true,
		}
	} else {
		return Word_Index{Exists: false}
	}
}


func Reverse_String(str string) (result string) {
	// iterate over str and prepend to result
	for _ , v := range str {
		result = string(v) + result
	}
	return
}

func Debug(args ...string) {
	fmt.Println(args)
}

func Word_Values(
	line string,
	reverse bool,
) (int, string) {

	word_index := -1

	word := Word_Index_Func(line, reverse)
	if word.Exists {
		word_index = word.Index
	}

	return word_index, word.Word
}

func Check(e error) {
	if e != nil {
			panic(e)
	}
}

func Read_from_File(filename string) []string {
	input, err := os.ReadFile(filename)
	Check(err)

	return strings.Split(string(input), "\n")
}

func Write_to_file(filename string, output string) {
	out, err := os.OpenFile(filename,
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	defer out.Close()
	_, err = out.WriteString(output+"\n")
	Check(err)
}

func Get_Full_Filename(io, day, filename string) string {
	return fmt.Sprintf("days/%v/%v/%v.txt", day, io, filename)
}

func Is_Game_Possible(count, max int) bool {
	return count <= max
}

func Create_Grid(lines []string) (grid [][]string) {
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	return grid
}

func Is_Number(char string) bool {
	_, e := strconv.Atoi(char)

	if e != nil {
		return false
	}

	return true
}

func Is_End_Index (line []string, current_index int) bool {
	return (current_index == len(line)-1) || !Is_Number(line[current_index+1])
}

type Values_Map struct {
	Value int
	Row int
	Start int
	End int
}

func Get_Grid_Values(grid [][]string) (values_map []Values_Map) {
	values_map = []Values_Map{}

	for i, line := range grid {
		value := ""
		start_index := 0
		end_index := 0

		for j, val := range line {
			is_num := Is_Number(val)
			if value == "" && is_num {
				value = val
				start_index = j
			} else if is_num { value += val }

			if Is_End_Index(line, j) {
				end_index = j
				val_int, _ := strconv.Atoi(value)
				if val_int != 0 {
					values_map = append(values_map, Values_Map{
						Value: val_int,
						Row: i,
						Start: start_index,
						End: end_index,
					})
				}
				value = ""
			}
		}
	}

	return values_map
}

func Neighbor_Is_Symbol(grid [][]string, row, start, end int) bool {
	row_len := len(grid[row])
	col_len := len(grid)

	for i := row-1 ; i <= row+1 ; i++ {
		if i < 0 || i >= col_len { continue }
		for j := start-1 ; j <= end+1 ; j++ {
			if j < 0 || j >= row_len { continue }
			if i == row && (j >= start && j <= end) { continue }
			if grid[i][j] != "." { return true }
		}
	}

	return false
}

type Coordinates struct{
	Row int
	Col int
}

func Get_Star_Values(grid [][]string) (values_map []Coordinates) {
	values_map = []Coordinates{}

	for i, line := range grid {

		for j, val := range line {
			if val == "*" {
				values_map = append(values_map, Coordinates{
					Row: i,
					Col: j,
				})
			}
		}
	}

	return values_map
}

func Gear_Neighbor_Count(grid [][]string, coordinates []Coordinates) (result []int) {

	for i, coordinate := range coordinates {
		row_len := len(grid)
		col_len := len(grid[coordinate.Row])
		count := 0

		if i < 0 || i >= row_len { continue }
		for j := coordinate.Col-1 ; j <= coordinate.Col+1 ; j++ {
			if j < 0 || j >= col_len { continue }
			if (i == coordinate.Row) && (j == coordinate.Col) { continue }

			if Is_Number(grid[i][j]) { count++ }

			result = append(result, )
		}
		result = append(result, count)
	}

	return result
}

func Sum_of_Neighbord_Vals(full_output_filename string, values_map []Values_Map, grid [][]string) (sum int) {
	for _, value := range values_map {
		row := value.Row
		start := value.Start
		end := value.End

		if Neighbor_Is_Symbol(grid, row, start, end) {
			val := strconv.Itoa(value.Value)
			Write_to_file(full_output_filename, val)
			sum += value.Value
		}
	}
	return sum
}