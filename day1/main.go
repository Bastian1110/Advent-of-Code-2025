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

func parseInput(lines []string) ([]string, []int) {
	var letters []string
	var turns []int
	for _, line := range lines {
		letter := string(line[0])
		n, _ := strconv.Atoi(line[1:])

		letters = append(letters, letter)
		turns = append(turns, n)
	}
	return letters, turns
}

func AbsInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}



func useLockOne(actualPosition int, direction string, steps int)int{
	tempPos := actualPosition
	for range steps {
		if(direction == "L"){
			tempPos--
		}else{
			tempPos++
		}
		if(tempPos == 100){
			tempPos = 0
			continue
		}
		if(tempPos == -1){
			tempPos = 99
		}
	}
	return tempPos
}

func one(lines []string) {
	letters, turns := parseInput(lines)
	lockposition := 50;
	answer := 0
	for i, _ := range letters {
		lockposition = useLockOne(lockposition, letters[i], turns[i])
		fmt.Println(letters[i], turns[i], " - ",lockposition)
		if(lockposition == 0){
			answer++
		}
	}
	fmt.Println(answer)

}

func useLockTwo(actualPosition int, direction string, steps int)(int, int){
	tempPos := actualPosition
	passedByZero := 0
	for range steps {
		if(direction == "L"){
			tempPos--
		}else{
			tempPos++
		}
		if(tempPos == 100){
			tempPos = 0
			passedByZero++
			continue
		}
		if(tempPos == -1){
			passedByZero++
			tempPos = 99
		}
	}
	return tempPos, passedByZero
}

func two(lines []string) {
	letters, turns := parseInput(lines)
	lockposition := 50;
	passedByZero := 0 
	answer := 0
	for i, _ := range letters {
		passedByZeroStep := 0
		lockposition, passedByZeroStep = useLockTwo(lockposition, letters[i], turns[i])
		passedByZero += passedByZeroStep
		fmt.Println(letters[i], turns[i], " - ",lockposition)
		if(lockposition == 0){
			answer++
		}
	}
	fmt.Println(answer, passedByZero)

}


func main() {
	lines := ReadFile("./test.txt")
	one(lines)
	two(lines)

}
