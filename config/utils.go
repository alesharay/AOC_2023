package config

import (
	"fmt"
	"os"
	"slices"
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

	Keys = []string {
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
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