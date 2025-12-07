package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(path string) []string {
	var res []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Cant open the file 😬")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	return res
}

func parseInput(lines []string) [][]rune {
	var result [][]rune
	for _, v := range lines {
		var temp []rune
		for _, r := range v {
			temp = append(temp, r)
		}
		result = append(result, temp)
	}
	return result
}
func one(lines [][]rune) {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '@' {
				temp := 0
				if j-1 >= 0 {
					if lines[i][j-1] == '@' {
						temp += 1
					}
				}
				if j+1 < len(lines[i]) {
					if lines[i][j+1] == '@' {
						temp += 1
					}
				}
				if i-1 >= 0 {
					if lines[i-1][j] == '@' {
						temp += 1
					}
				}
				if i+1 < len(lines) {
					if lines[i+1][j] == '@' {
						temp += 1
					}
				}

				if i-1 >= 0 && j-1 >= 0 {
					if lines[i-1][j-1] == '@' {
						temp += 1
					}
				}
				if i-1 >= 0 && j+1 < len(lines[i]) {
					if lines[i-1][j+1] == '@' {
						temp += 1
					}
				}

				if i+1 < len(lines) && j-1 >= 0 {
					if lines[i+1][j-1] == '@' {
						temp += 1
					}
				}

				if i+1 < len(lines) && j+1 < len(lines[i]) {
					if lines[i+1][j+1] == '@' {
						temp += 1
					}
				}
				fmt.Println("Temp", temp)
				if temp < 4 {
					total += 1
				}
			}
		}
	}
	fmt.Println(total)
}

func getHowManyCanBeRemoved(lines [][]rune) int {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '@' {
				temp := 0
				if j-1 >= 0 {
					if lines[i][j-1] == '@' {
						temp += 1
					}
				}
				if j+1 < len(lines[i]) {
					if lines[i][j+1] == '@' {
						temp += 1
					}
				}
				if i-1 >= 0 {
					if lines[i-1][j] == '@' {
						temp += 1
					}
				}
				if i+1 < len(lines) {
					if lines[i+1][j] == '@' {
						temp += 1
					}
				}

				if i-1 >= 0 && j-1 >= 0 {
					if lines[i-1][j-1] == '@' {
						temp += 1
					}
				}
				if i-1 >= 0 && j+1 < len(lines[i]) {
					if lines[i-1][j+1] == '@' {
						temp += 1
					}
				}

				if i+1 < len(lines) && j-1 >= 0 {
					if lines[i+1][j-1] == '@' {
						temp += 1
					}
				}

				if i+1 < len(lines) && j+1 < len(lines[i]) {
					if lines[i+1][j+1] == '@' {
						temp += 1
					}
				}
				if temp < 4 {
					total += 1
				}
			}
		}
	}
	return total
}

func two(lines [][]rune) {
	total := 0

	for {
		toRemove := getHowManyCanBeRemoved(lines)
		fmt.Println("Step:")
		printLines(lines)
		fmt.Println("To Remove: ", toRemove)

		if toRemove == 0 {
			break
		}

		next := make([][]rune, len(lines))
		for i := range lines {
			next[i] = make([]rune, len(lines[i]))
			copy(next[i], lines[i])
		}

		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[i]); j++ {
				if lines[i][j] != '@' {
					continue
				}

				temp := 0
				if j-1 >= 0 && lines[i][j-1] == '@' {
					temp++
				}
				if j+1 < len(lines[i]) && lines[i][j+1] == '@' {
					temp++
				}
				if i-1 >= 0 && lines[i-1][j] == '@' {
					temp++
				}
				if i+1 < len(lines) && lines[i+1][j] == '@' {
					temp++
				}
				if i-1 >= 0 && j-1 >= 0 && lines[i-1][j-1] == '@' {
					temp++
				}
				if i-1 >= 0 && j+1 < len(lines[i]) && lines[i-1][j+1] == '@' {
					temp++
				}
				if i+1 < len(lines) && j-1 >= 0 && lines[i+1][j-1] == '@' {
					temp++
				}
				if i+1 < len(lines) && j+1 < len(lines[i]) && lines[i+1][j+1] == '@' {
					temp++
				}
				if temp < 4 {
					next[i][j] = '.'
				} else {
					next[i][j] = '@'
				}
			}
		}

		total += toRemove
		lines = next
	}

	fmt.Println("Total:", total)
}

func printLines(lines [][]rune) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			fmt.Print(string(lines[i][j]))
		}
		fmt.Print("\n")
	}
}

func main() {
	lines := ReadFile("./test.txt")
	paperRolls := parseInput(lines)
	one(paperRolls)
	two(paperRolls)

}
