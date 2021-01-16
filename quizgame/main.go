package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type puzzle struct {
	question string
	answer   string
}

func main() {
	filename := flag.String("filename", "problems.csv", "Indicate a filename")
	timeout := flag.Int64("deadline", 30, "Deadline in seconds")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(file)

	puzzles := []puzzle{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		puzzles = append(puzzles, puzzle{record[0], record[1]})
	}

	fmt.Println("Please, hit enter to start the quiz!")
	var started string
	fmt.Scanln(&started)

	var correct int
	quit := make(chan bool)

	go func() {
		time.Sleep(time.Duration(*timeout) * time.Second)
		fmt.Println("Sorry, your time is up! The quiz is over")
		quit <- true
	}()

	go func() {
		for _, p := range puzzles {
			fmt.Println(p.question)
			var answer string
			fmt.Scanln(&answer)
			answer = strings.TrimSpace(answer)

			if answer == p.answer {
				correct++
			}
		}
		quit <- true
	}()

	<-quit

	fmt.Printf("Total questions: %d, Correctly answered: %d\n", len(puzzles), correct)
}
