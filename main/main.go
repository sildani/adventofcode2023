package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	dayone "github.com/sildani/adventofcode2023/dayone"
)

func main() {
    inputReader := bufio.NewReader(os.Stdin)

    fmt.Println("Welcome to Advent of Code 2023 Solutions by Daniel!")

    fmt.Println("Please enter day [1-1]:")
    dayInput, err := inputReader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }

    day, _ := strconv.Atoi(strings.Trim(dayInput, "\n"))
    var fileName string
    if day == 1 {
        fileName = "./inputs/dayone/input.txt"
    } else {
        fmt.Println("Value must be [1-1]")
        os.Exit(1)
    }

    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    var input []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input = append(input, scanner.Text())
    }

    var results []int
    if (day == 1) {
        results = append(results, dayone.Process(input))
    } 
    fmt.Printf("Answer: %v\n\n", results)
}