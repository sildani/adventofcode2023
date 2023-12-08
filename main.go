package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"

    dayone "github.com/sildani/adventofcode2023/dayone"
    daytwo "github.com/sildani/adventofcode2023/daytwo"
    daythree "github.com/sildani/adventofcode2023/daythree"
)

func main() {
    inputReader := bufio.NewReader(os.Stdin)

    fmt.Println("Welcome to Advent of Code 2023 Solutions by Daniel!")

    fmt.Println("Please enter day [1-3]:")
    dayInput, err := inputReader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }

    day, _ := strconv.Atoi(strings.Trim(dayInput, "\n"))
    var fileName string
    switch day {
    case 1:
        fileName = "./dayone/input.txt"
    case 2:
        fileName = "./daytwo/input.txt"
    case 3:
        fileName = "./daythree/input.txt"
    default:
        fmt.Println("Value must be [1-3]")
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
    switch day {
    case 1:
        results = append(results, dayone.Process(input))
    case 2:
        partOneResult, partTwoResult := daytwo.Process(input)
        results = append(results, partOneResult, partTwoResult)
    case 3:
        results = append(results, daythree.Process(input))
    default:
        // ignore
    }
    fmt.Printf("Answer: %v\n\n", results)
}