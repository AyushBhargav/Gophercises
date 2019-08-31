package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	filename := flag.String("f", "problem.csv", "Specify questionnaire")
	timeLimit := flag.Int("tl", 30, "Time limit for test")

	flag.Parse()

	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		panic("Boom! Couln't open file")
	}

	fileContent := string(data)

	questions, err := csv.NewReader(strings.NewReader(fileContent)).ReadAll()
	if err != nil {
		panic("Can't parse string")
	}

	fmt.Println("Press enter to continue...")
	fmt.Scanln()

	count := 0

	go func() {
		time.Sleep(time.Duration(*timeLimit) * time.Second)
		fmt.Printf("Time's up! Be proud of your stupidity. You answered %d of %d questions", count, len(questions))
		os.Exit(0)
	}()

	for i, question := range questions {
		fmt.Printf("Q.%d %s = ", i, question[0])
		var ans string
		fmt.Scanln(&ans)
		if ans == question[1] {
			count++
		}
	}
}
