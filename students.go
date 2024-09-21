package main

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./students.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	studentsDict := make(map[string][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rawLine := scanner.Text()
		studentAndMark := strings.Split(strings.Trim(rawLine, " "), " ")
		if len(studentAndMark) != 2 {
			slog.Warn(fmt.Sprintf("More than 2 word! Line: \"%v\" - Skipping it...", rawLine))
			continue
		}

		mark, err := strconv.Atoi(studentAndMark[1])
		if err != nil {
			slog.Warn(fmt.Sprintf("Second word is not a mark! Line: \"%v\" - Skipping it...", rawLine))
			continue
		}

		_, exist := studentsDict[studentAndMark[0]]
		if !exist {
			studentsDict[studentAndMark[0]] = []int{mark}
		} else {
			buff := studentsDict[studentAndMark[0]]
			buff = append(buff, mark)
			studentsDict[studentAndMark[0]] = buff
		}
	}

	students := make([]string, 0)
	for s := range studentsDict {
		students = append(students, s)
	}
	sort.Strings(students)
	for _, s := range students {
		fmt.Println(s)
		fmt.Printf("Score: %v\n", arrString(studentsDict[s]))
		fmt.Printf("Average score: %.2f\n", avg(studentsDict[s]))
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func arrString(arr []int) string {
	res := strconv.Itoa(arr[0])
	for _, v := range arr {
		res += ", " + strconv.Itoa(v)
	}
	return res
}

func avg(arr []int) float32 {
	res := 0
	for _, v := range arr {
		res += v
	}
	return float32(res) / float32(len(arr))
}
