package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func parseInput(line string) [][]int {
	var result [][]int
	stringRanges := strings.Split(line, ",")
	for _, v := range stringRanges {
		parts := strings.Split(v, "-")
		begin, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		result = append(result, []int{begin, end})
	}
	return result
}

func isInsideBox(l int, r int, v int) bool {
	return (l <= v) && (v <= r)
}

func z(number string) []int {
	Z := make([]int, len(number))
	L := 0
	R := 0
	for i := range len(number) - 1 {
		i = i + 1
		if !isInsideBox(L, R, i) {
			counter := 0
			tempI := i
			lStart := -1
			rStart := -1
			for _, v := range number {
				if string(v) != string(number[tempI]) {
					break
				}
				if lStart == -1 {
					lStart = tempI
				}
				rStart = tempI
				tempI++
				counter++
				if tempI > len(number)-1 {
					break
				}
			}
			Z[i] = counter
			if lStart != -1 && rStart != -1 {
				L = lStart
				R = rStart
			}
		} else {
			k := i - L
			bound := R - i + 1
			if Z[k] < bound {
				Z[i] = Z[k]
			} else {
				current := bound
				tempL := -1
				temrR := 1
				for {
					if i+current > len(number)-1 || current > len(number)-1 {
						break
					}
					if number[current] != number[i+current] {
						break
					}
					if tempL == -1 {
						tempL = i
					}
					temrR = i + current
					current += 1
				}
				if temrR != -1 && tempL != -1 {
					L = tempL
					R = temrR
				}
				Z[i] = current
			}
		}
	}
	return Z
}

func checkForInvalidIdsOne(r []int) int {
	total := 0
	for i := r[0]; i <= r[1]; i++ {
		id := strconv.Itoa(i)
		if len(id)%2 == 0 {
			if !isRepeated(id) {
				fmt.Println("Its repated!", id)
				total += i
			}
		}
	}
	return total
}

func isRepeated(number string) bool {
	center := len(number) / 2
	left := number[:center]
	right := number[center:]
	return left != right
}

func one(lines string) {
	ranges := parseInput(lines)
	total := 0

	for _, v := range ranges {
		total += checkForInvalidIdsOne(v)
	}
	fmt.Println(total)

}

func IsPeriodicFromZ(z []int) bool {
	n := len(z)
	if n == 0 {
		return false
	}

	for i := 1; i < n; i++ {
		if n%i == 0 && z[i] >= n-i {
			return true
		}
	}
	return false
}

func checkForInvalidIdTwo(r []int) int {
	total := 0
	for i := r[0]; i <= r[1]; i++ {
		id := strconv.Itoa(i)
		zArray := z(id)
		if IsPeriodicFromZ(zArray) {
			fmt.Println("Its repated!", id)
			total += i
		}
	}
	return total
}

func two(lines string) {
	ranges := parseInput(lines)
	total := 0

	for _, v := range ranges {
		total += checkForInvalidIdTwo(v)
	}
	fmt.Println(total)

}

func main() {
	lines := ReadFile("./test.txt")
	one(lines[0])
	two(lines[0])

}
