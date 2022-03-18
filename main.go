package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("difficulty", "easy", "a csv file in the format of 'question, answer'") // get difficulty from the user as flag. Default is easy.
	timeLimit := flag.Int("time", 30, "The time limit for quiz game in seconds")                       // get time limit from the user. Default is 30 seconds.
	flag.Parse()

	fmt.Println("Opening", *csvFilename, "questions")
	fmt.Println("You have", *timeLimit, "seconds to complete the quiz")

	file, err := os.Open("./questions/" + *csvFilename + ".csv")
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file: %s\n", *csvFilename))
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse csv file")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer) // Scanf trims whitespace. So if user puts something like "  2  " it becomes "2"
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("You scored", correct, "out of", len(problems))
			return
		case answer := <-answerCh:
			if answer == problem.a {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Println("Incorrect!")
			}
		}
	}
	fmt.Println("You scored", correct, "out of", len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // If there are empty spaces in the answer trim it.
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
