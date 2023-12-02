package dayone

import (
    "regexp"
    "strconv"
    "strings"
)

func Process(input []string) int {
    regex, _ := regexp.Compile("\\d")

    var sum int

    for _, line := range input {
        i := ReplaceWordNumbersForDigits(line)
        digits := regex.FindAllString(i, -1)
        if len(digits) > 0 {
            first, _ := strconv.Atoi(digits[0])
            last, _ := strconv.Atoi(digits[len(digits)-1])
            // fmt.Printf("line: %v; number: %v%v; \n", line, first, last)
            // lastSum := sum
            sum = sum + (first * 10) + last
            // fmt.Printf("sum: %v + %v = %v;\n", lastSum, (first * 10) + last, sum)
        }
    }

    return sum
}

func ReplaceWordNumbersForDigits(input string) string {
    return strings.NewReplacer(
        "oneight", "18",
        "twone", "21",
        "threeight", "38",
        "eighthree", "83",
        "eightwo", "82",
        "fiveight", "58",
        "sevenine", "79",
        "nineight", "98",
        "one", "1",
        "two", "2",
        "three", "3",
        "four", "4",
        "five", "5",
        "six", "6",
        "seven", "7",
        "eight", "8",
        "nine", "9").Replace(input)
}