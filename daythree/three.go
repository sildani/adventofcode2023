package daythree

import (
    "strconv"
    "unicode"
)

func Process(input []string) int {
    var sum int
    var candidatePartNumber string
    for lineIndex, line := range input {
        for runeIndex, rune := range line {
            isPartNumber := false
            if isDigit(rune) {
                candidatePartNumber = candidatePartNumber + string(rune)
            } else if candidatePartNumber != "" {
                // if current rune is symbol, it's a part number
                if isSymbol(rune) {
                    sum = addPartNumber(candidatePartNumber, sum)
                    candidatePartNumber = ""
                    continue
                }
                // if the rune right before candidate is a symbol, it's a part number
                if runeIndex-len(candidatePartNumber) > 0 {
                    rangeToCheck := line[runeIndex-len(candidatePartNumber)-1 : runeIndex-len(candidatePartNumber)]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // top left corner
                if lineIndex == 0 && runeIndex-len(candidatePartNumber) == 0 {
                    rangeToCheck := input[lineIndex+1][runeIndex-len(candidatePartNumber) : len(candidatePartNumber)+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // top right corner
                if lineIndex == 1 && runeIndex == 0 {
                    isPartNumber := false
                    rangeToCheck := input[lineIndex-1][len(line)-len(candidatePartNumber)-1 : len(line)-len(candidatePartNumber)]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                    rangeToCheck = input[lineIndex][len(line)-len(candidatePartNumber)-1 : len(line)]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // bottom left corner
                if lineIndex == len(input)-1 && runeIndex == len(candidatePartNumber) {
                    rangeToCheck := input[lineIndex-1][runeIndex-len(candidatePartNumber) : len(candidatePartNumber)+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // left edge
                if lineIndex > 0 && lineIndex < len(input)-1 && runeIndex == len(candidatePartNumber) {
                    rangeToCheck := input[lineIndex-1][runeIndex-len(candidatePartNumber) : len(candidatePartNumber)+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                    rangeToCheck = input[lineIndex-1][runeIndex-len(candidatePartNumber) : len(candidatePartNumber)+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                    rangeToCheck = input[lineIndex+1][runeIndex-len(candidatePartNumber) : len(candidatePartNumber)+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // right edge
                if lineIndex > 1 && lineIndex <= len(input)-1 && runeIndex == 0 {
                    rangeToCheck := input[lineIndex-1][len(line)-len(candidatePartNumber)-1 : len(line)-len(candidatePartNumber)]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                    rangeToCheck = input[lineIndex-2][len(line)-len(candidatePartNumber)-1 : len(line)]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                    rangeToCheck = input[lineIndex][len(line)-len(candidatePartNumber)-1 : len(line)]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // top edge
                if lineIndex == 0 && lineIndex <= len(input)-1 && runeIndex-len(candidatePartNumber) > 0 {
                    rangeToCheck := input[lineIndex+1][runeIndex-len(candidatePartNumber)-1 : runeIndex+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // bottom edge
                if lineIndex == len(input)-1 && runeIndex-len(candidatePartNumber) > 0 {
                    rangeToCheck := input[lineIndex-1][runeIndex-len(candidatePartNumber)-1 : runeIndex+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // middle
                if lineIndex > 0 && lineIndex < len(input)-1 && runeIndex-len(candidatePartNumber) > 0 {
                    rangeToCheck := input[lineIndex-1][runeIndex-len(candidatePartNumber)-1 : runeIndex+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                    rangeToCheck = input[lineIndex+1][runeIndex-len(candidatePartNumber)-1 : runeIndex+1]
                    sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                    if isPartNumber {
                        isPartNumber = false
                        continue
                    }
                }
                // not a part number, reset
                candidatePartNumber = ""
            }
            // bottom right corner
            if candidatePartNumber != "" && lineIndex == len(input)-1 && runeIndex == len(line)-1 {
                rangeToCheck := line[runeIndex-len(candidatePartNumber) : runeIndex-len(candidatePartNumber)+1]
                sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                if isPartNumber {
                    isPartNumber = false
                    continue
                }
                rangeToCheck = input[lineIndex-1][runeIndex-len(candidatePartNumber) : runeIndex+1]
                sum, candidatePartNumber, isPartNumber = checkCandidatePartNumberWithRange(rangeToCheck, sum, candidatePartNumber, isPartNumber)
                if isPartNumber {
                    isPartNumber = false
                    continue
                }
            }
        }
    }

    return sum
}

func checkCandidatePartNumberWithRange(rangeToCheck string, sum int, candidatePartNumber string, isContinue bool) (int, string, bool) {
    if checkCandidatePartNumber(rangeToCheck, sum, candidatePartNumber) {
        sum = addPartNumber(candidatePartNumber, sum)
        candidatePartNumber = ""
        isContinue = true
    }
    return sum, candidatePartNumber, isContinue
}

func checkCandidatePartNumber(rangeToCheck string, sum int, candidatePartNumber string) bool {
    for _, r := range rangeToCheck {
        if isSymbol(r) {
            return true
        }
    }
    return false
}

func isDigit(rune rune) bool {
    return unicode.IsDigit(rune)
}

func isSpace(rune rune) bool {
    return string(rune) == "."
}

func isSymbol(rune rune) bool {
    return !unicode.IsDigit(rune) && !isSpace(rune)
}

func addPartNumber(candidatePartNumber string, sum int) int {
    x, _ := strconv.Atoi(candidatePartNumber)
    sum += x
    return sum
}
