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
	var tachyon [][]rune
	for _, l := range lines {
		tachyon = append(tachyon, []rune(l))
	}
	return tachyon
}

func indexOfRune(slice []rune, target rune) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func CopyRuneMatrix(src [][]rune) [][]rune {
	if src == nil {
		return nil
	}

	dst := make([][]rune, len(src))
	for i := range src {
		if src[i] != nil {
			row := make([]rune, len(src[i]))
			copy(row, src[i])
			dst[i] = row
		}
	}
	return dst
}

func prettyPrint(m [][]rune) {
	for _, row := range m {
		fmt.Println(string(row))
	}
}

func one(tachyon [][]rune) {
	total := 0
	for i := 0; i < len(tachyon); i++ {
		fmt.Println("Step:", i)
		prettyPrint(tachyon)
		temp := CopyRuneMatrix(tachyon)
		for y := 0; y < len(tachyon); y++ {
			for x := 0; x < len(tachyon[0]); x++ {
				if y+1 < len(tachyon) {

					if tachyon[y][x] == 'S' {
						temp[y+1][x] = '|'
					}
					if tachyon[y][x] == '^' && tachyon[y-1][x] == '|' {
						if tachyon[y][x-1] != '^' {
							temp[y][x-1] = '|'
						}
						if tachyon[y][x+1] != '^' {
							temp[y][x+1] = '|'
						}
					}
					if tachyon[y][x] == '|' {
						if tachyon[y+1][x] != '^' {
							temp[y+1][x] = '|'
						}
					}
				}
			}
		}
		tachyon = CopyRuneMatrix(temp)
	}
	for i := 0; i < len(tachyon); i++ {
		for j := 0; j < len(tachyon[0]); j++ {
			if tachyon[i][j] == '^' && tachyon[i-1][j] == '|' {
				total += 1

			}

		}

	}

	fmt.Println("Total:", total)

}

func two(grid [][]rune) {

	rows := len(grid)
	cols := len(grid[0])

	sr, sc := -1, -1
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 'S' {
				sr, sc = r, c
				break
			}
		}
		if sr != -1 {
			break
		}
	}

	dp := make([][]int, rows)
	for r := range dp {
		dp[r] = make([]int, cols)
	}

	dp[sr][sc] = 1

	for r := sr; r < rows-1; r++ {
		for c := 0; c < cols; c++ {
			ways := dp[r][c]
			if ways == 0 {
				continue
			}

			// Look at cell below
			below := grid[r+1][c]

			if below == '^' {
				// Split diagonally down-left and down-right (if in bounds)
				if c-1 >= 0 {
					dp[r+1][c-1] += ways
				}
				if c+1 < cols {
					dp[r+1][c+1] += ways
				}
			} else {
				// Anything that's not '^' is treated as normal pass-through
				dp[r+1][c] += ways
			}
		}
	}

	// Sum all paths that end in the bottom row.
	total := 0
	for c := 0; c < cols; c++ {
		total += dp[rows-1][c]
	}
	fmt.Println(total)
}

func main() {
	lines := ReadFile("./test.txt")
	tachyon := parseInput(lines)
	two(tachyon)
}
