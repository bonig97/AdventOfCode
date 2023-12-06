package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile("[0-9]")

	for scanner.Scan() {
		s := scanner.Text()
		digits := re.FindAllString(s, -1)

		if len(digits) >= 2 {
			firstDigit, lastDigit := digits[0], digits[len(digits)-1]

			firstInt, err := strconv.Atoi(firstDigit)
			if err != nil {
				log.Printf("Error: %v\n", err)
				continue
			}
			lastInt, err := strconv.Atoi(lastDigit)
			if err != nil {
				log.Printf("Error: %v\n", err)
				continue
			}

			sum := firstInt + lastInt
			fmt.Printf("First: %d, Last: %d, Sum: %d\n", firstInt, lastInt, sum)
		} else if len(digits) == 1 {
			firstInt, err := strconv.Atoi(digits[0])
			if err != nil {
				log.Printf("Error: %v\n", err)
				continue
			}
			fmt.Printf("First: %d, Sum: %d\n", firstInt, firstInt*2)
		} else {
			fmt.Println("No digits found")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
