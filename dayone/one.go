package dayone

import (
	"regexp"
	"strconv"
)

func Process(input []string) int {
    regex, _ := regexp.Compile("\\d")

    var sum int

    for _, i := range input {
        digits := regex.FindAllString(i, -1)
        if len(digits) > 0 {
            first, _ := strconv.Atoi(digits[0])
            last, _ := strconv.Atoi(digits[len(digits)-1])
            sum = sum + (first * 10) + last
        }
    }

    return sum
}