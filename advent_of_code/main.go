package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	bestElf, bestCalories := findMostCalores("input.txt")
	fmt.Println("The elf with the most calories is ", bestElf, "with ", bestCalories, " calories")
}

func findMostCalores(inputPath string) (int, int){
	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	calorieList := strings.Split(string(data), "\r\n")
	var totalCalories int
	var bestCalories int
	elves := 0
	bestElf := 0
	for _, calories := range calorieList {
		if calories == "" {
			elves++
			if bestCalories < totalCalories{
				bestCalories = totalCalories
				bestElf = elves
			}
			totalCalories = 0
			continue
		}
		intCalories, err := strconv.Atoi(calories)
		if err != nil {
			log.Fatal(err)
		}
		totalCalories += intCalories

	}
	
	return bestElf, bestCalories
}
