package config

import (
	"fmt"
	"os"
	"strings"
)

var (
	word_to_num = map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}
)


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
	return fmt.Sprintf("config/%v/%v/%v.txt", io, day, filename)
}