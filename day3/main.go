package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	var result [][]int
	for _, v := range lines {
		var temp []int
		for _, n := range v {
			num, _ := strconv.Atoi(string(n))
			temp = append(temp, num)
		}
		result = append(result, temp)
	}
	return result
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

func getMaxBaterries(batteries []int) int {
	fmt.Println("Bateries:", batteries)
	firstIndex := MaxIndex(batteries)
	first := strconv.Itoa(batteries[firstIndex])
	var second string
	if firstIndex+2 > len(batteries) {
		firstIndex = MaxIndex(batteries[:len(batteries)-1])
		first = strconv.Itoa(batteries[firstIndex])
		second = strconv.Itoa(batteries[len(batteries)-1])
	} else {
		secondArray := batteries[firstIndex+1:]
		fmt.Println(secondArray)
		secondIndex := MaxIndex(secondArray)
		second = strconv.Itoa(secondArray[secondIndex])
	}

	final, _ := strconv.Atoi(first + second)
	fmt.Println("voltage:", final)
	return final
}

func one(lines []string) {
	banks := parseInput(lines)
	total := 0
	for _, bank := range banks {
		total += getMaxBaterries(bank)
	}
	fmt.Println(total)
}

func getMaxRightContained(numbers []int, missing int) int {
	limit := len(numbers) - missing

	maxIdx := 0
	maxVal := numbers[0]

	for i := 1; i <= limit; i++ {
		if numbers[i] > maxVal {
			maxVal = numbers[i]
			maxIdx = i
		}
	}
	return maxIdx
}

func getMaxBaterriesTwo(batteries []int) int {
	fmt.Println(batteries)
	var result string
	tempBatteries := batteries
	for i := 0; i < 12; i++ {
		max := getMaxRightContained(tempBatteries, 12-i)
		fmt.Println("Max Index", max, " Value", tempBatteries[max], "Temp:", tempBatteries)
		result += strconv.Itoa(tempBatteries[max])
		tempBatteries = tempBatteries[max+1:]
	}
	final, _ := strconv.Atoi(result)
	fmt.Println("Final:", final)
	return final

}

func two(lines []string) {
	banks := parseInput(lines)
	total := 0
	for _, v := range banks {
		total += getMaxBaterriesTwo(v)
	}
	fmt.Println("Total:", total)
}

func main() {
	lines := ReadFile("./input.txt")
	//one(lines)
	two(lines)
}
