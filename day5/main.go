package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func parseInput(lines []string) ([][]int, []int) {
	var ranges [][]int
	var ingredients []int
	parsingIds := false
	for _, v := range lines {
		if strings.TrimSpace(v) == "" {
			parsingIds = true
			continue
		}
		if parsingIds {
			num, _ := strconv.Atoi(v)
			ingredients = append(ingredients, num)
		} else {
			parts := strings.Split(v, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, []int{start, end})
		}
	}
	return ranges, ingredients
}

func isInRange(r []int, id int) bool {
	fmt.Println("Is in range:", r, id)
	is := (r[0] <= id) && (id <= r[1])
	fmt.Println("Returning :", is)
	return is
}

func contains(nums []int, target int) bool {
	for _, n := range nums {
		if n == target {
			return true
		}
	}
	return false
}

func one(ranges [][]int, ids []int) {
	total := 0
	var freshIds []int
	for _, v := range ids {
		for i := 0; i < len(ranges); i++ {
			fmt.Println("Checking range:", ranges[i])
			if isInRange(ranges[i], v) && !contains(freshIds, v) {
				total += 1
				freshIds = append(freshIds, v)
			}
		}
	}
	fmt.Println(total)
}

func two(ranges [][]int) {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})

	curStart := ranges[0][0]
	curEnd := ranges[0][1]
	total := 0

	for i := 1; i < len(ranges); i++ {
		L := ranges[i][0]
		R := ranges[i][1]

		if L > curEnd+1 {
			total += curEnd - curStart + 1
			curStart = L
			curEnd = R
		} else {
			if R > curEnd {
				curEnd = R
			}
		}
	}

	total += curEnd - curStart + 1
	fmt.Println("Total:", total)

}

func main() {
	lines := ReadFile("./test.txt")
	ranges, _ := parseInput(lines)
	//one(ranges, ids)
	two(ranges)

}
