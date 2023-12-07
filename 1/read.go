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

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)
	re := regexp.MustCompile("[0-9]")
	totalSum := 0

	for scanner.Scan() {
		s := scanner.Text()
		digits := re.FindAllString(s, -1)

		var concat string
		if len(digits) >= 2 {
			concat = digits[0] + digits[len(digits)-1]
		} else if len(digits) == 1 {
			concat = digits[0] + digits[0]
		}

		if concat != "" {
			concatInt, err := strconv.Atoi(concat)
			if err != nil {
				log.Printf("Error: %v\n", err)
				continue
			}
			totalSum += concatInt
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total sum: %d\n", totalSum)
}
