package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type quiz struct {
	questions string
	answers   string
}

func main() {
	myslice := []quiz{
		{questions: "What is the capital of India?", answers: "New Delhi"},
		{questions: "What is 2 + 2?", answers: "4"},
		{questions: "What is the color of the sky?", answers: "blue"},
		{questions: "How many days are there in a week?", answers: "7"},
		{questions: "What is 2 + 3?", answers: "5"},
		{questions: "Which animal says 'meow'?", answers: "cat"},
		{questions: "Which planet do we live on?", answers: "Earth"},
		{questions: "What is the opposite of hot?", answers: "cold"},
		{questions: "What is the opposite of sad?", answers: "happy"},
		{questions: "What is the color of grass?", answers: "green"},
	}

	var score = 0
	var userAnswer string

	fmt.Println("******************      Quiz App    *************************")
	fmt.Println(" ")

	for i, qns := range myslice {

		fmt.Println("Q", i+1, ": ", qns.questions)
		//fmt.Scanf("%s", &userAnswer)

		reader := bufio.NewReader(os.Stdin)
		userAnswer, _ = reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)
		fmt.Println("A: ", userAnswer)
		//string manipulation
		formattedanswer := strings.ToLower(strings.ReplaceAll(myslice[i].answers, " ", ""))
		formatteduseranswer := strings.ToLower(strings.ReplaceAll(userAnswer, " ", ""))
		//fmt.Println(userAnswer, " its from buffer tim", formattedanswer, formatteduseranswer)
		//if myslice[i].answers == userAnswer {
		if formattedanswer == formatteduseranswer {
			fmt.Println("Correct")
			score++
		} else {
			fmt.Println("Wrong")
		}
		fmt.Println(" ")
	}
	fmt.Println(" ")
	fmt.Println("Total Score:", score)
	fmt.Println(" ")
}
