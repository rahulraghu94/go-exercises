package main

import (
    "bufio"
    "encoding/csv"
	"fmt"
	"io"
    "log"
	"os"
	"strings"
)

type Quiz struct {
	question string
	answer string
}

func readCSV() []Quiz {
	csv_file, _ := os.Open("problems.csv")
	reader := csv.NewReader(bufio.NewReader(csv_file))
	var quiz []Quiz

	for {
		line, error := reader.Read()

		if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
		}
		quiz = append(quiz, Quiz{
			question: line[0],
			answer: line[1],
		})
	}
	return quiz
}

func run(){
	full_quiz := readCSV()
	correct := 0
	wrong := 0
	for _, question := range full_quiz {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Question: ")
		fmt.Println(question.question)

		fmt.Print("Answer: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		if strings.EqualFold(text, question.answer) {
			correct = correct + 1
		} else {
			wrong = wrong + 1
		}
		fmt.Println("")
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Correct: ", correct)
	fmt.Println("Wrong: ", wrong)
	percentage := float64(correct) / float64(len(full_quiz)) * float64(100)
	fmt.Println("Percentage: ", percentage)
}

func main() {
	run()
}