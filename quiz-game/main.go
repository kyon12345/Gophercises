package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// timeout := flag.Int("timeout", 0, "timeout for quiz")
	var timeout int
	flag.IntVar(&timeout, "t", 0, "timeout for quiz")
	flag.Parse()

	fmt.Println("Ready?")
	reader := bufio.NewReader(os.Stdin)
	_, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error occur during read from stdin", err)
	}

	timer := time.NewTimer(time.Duration(timeout) * time.Second)

	//read from csv.get the question and answer
	f, err := os.Open("./quiz-game/problems.csv")
	if err != nil {
		log.Fatal("unable to open input file", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("unable to parse file as csv", err)
	}

	//store them into map
	m := make(map[string]string, len(records))
	for _, record := range records {
		m[record[0]] = record[1]
	}
	score := 0

	//popup question
	done := make(chan struct{})
	go func() {
		for key, val := range m {
			fmt.Println("Question ", key)

			input := readInput()
			answer, _ := strconv.Atoi(val)

			if input == answer {
				score++
			}
		}
		done <- struct{}{}
	}()

	select {
	case <-timer.C:
		fmt.Println("time out")
	case <-done:
		fmt.Printf("your score is %d,total question %d ", score, len(m))
	}
}

func readInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error occur during read from stdin", err)
	}
	str := strings.TrimSuffix(input, "\r\n")
	if str == "" {
		return -1
	}
	res, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal("invalid number ", err)
	}
	return res
}
