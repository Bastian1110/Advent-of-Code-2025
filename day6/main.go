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
func parseInput(lines []string) [][]int {
	var operations [][]int

	for _, v := range lines {
		content := strings.Fields(v)

		for i := 0; i < len(content); i++ {
			if i >= len(operations) {
				operations = append(operations, []int{})
			}

			switch content[i] {
			case "*":
				operations[i] = append(operations[i], 0)
			case "+":
				operations[i] = append(operations[i], 1)
			default:
				num, _ := strconv.Atoi(content[i])
				operations[i] = append(operations[i], num)
			}
		}
	}

	return operations
}
func one(operations [][]int) {
	total := 0
	for _, v := range operations {
		temp := 0
		operator := v[len(v)-1]
		numns := v[:len(v)-1]
		if operator == 0 {
			temp = 1
		}
		for i := 0; i < len(numns); i++ {
			if operator == 0 {
				temp *= numns[i]
			}
			if operator == 1 {
				temp += numns[i]
			}
		}
		total += temp
	}
	fmt.Println(total)
}

func digits(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}

	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

func MaxIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	maxIdx := 0
	maxVal := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxIdx = i
		}
	}
	return maxIdx
}

func parseInputTwo(lines []string) [][]int {
	if len(lines) < 2 {
		return nil
	}

	numLines := lines[:len(lines)-1]
	opLine := lines[len(lines)-1]

	width := 0
	for _, s := range lines {
		if len(s) > width {
			width = len(s)
		}
	}

	numBufs := make([][]byte, len(numLines))
	for i, s := range numLines {
		b := make([]byte, width)
		copy(b, []byte(s))
		for j := len(s); j < width; j++ {
			b[j] = ' '
		}
		numBufs[i] = b
	}

	opBuf := make([]byte, width)
	copy(opBuf, []byte(opLine))
	for j := len(opLine); j < width; j++ {
		opBuf[j] = ' '
	}

	var result [][]int

	inCol := false
	var cur []int
	opCode := -1

	for col := width - 1; col >= 0; col-- {
		var digits []byte
		for r := 0; r < len(numBufs); r++ {
			c := numBufs[r][col]
			if c != ' ' {
				digits = append(digits, c)
			}
		}

		cOp := opBuf[col]

		if len(digits) > 0 {
			if !inCol {
				inCol = true
				cur = nil
				opCode = -1
			}

			if opCode == -1 {
				if cOp == '*' {
					opCode = 0
				} else if cOp == '+' {
					opCode = 1
				}
			}

			n, err := strconv.Atoi(string(digits))
			if err != nil {
				panic(err)
			}
			cur = append(cur, n)
		} else {
			if inCol {
				if opCode == -1 {
					opCode = 0
				}
				cur = append(cur, opCode)
				result = append(result, cur)
				inCol = false
			}
		}
	}

	if inCol {
		if opCode == -1 {
			opCode = 0
		}
		cur = append(cur, opCode)
		result = append(result, cur)
	}

	return result
}

func main() {
	lines := ReadFile("./test.txt")
	//ops := parseInput(lines)
	//one(ops)
	ops := parseInputTwo(lines)
	one(ops)

}
