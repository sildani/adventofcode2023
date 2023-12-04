package daythree

import (
    "reflect"
    "slices"
    "strconv"
    "strings"
    "unicode"
)

type token struct {
    value   string
    index   int
    isDigit bool
}

func Process(input []string) int {
    var sum int

    var previousTokens []token
    var previousPartTokens []token
    for i := range input {
        currentTokens := ParseLine(input[i])
        currentPartTokens := ParsePartNumbers(previousTokens, currentTokens, previousPartTokens)
        for i := range currentPartTokens {
            x, _ := strconv.Atoi(currentPartTokens[i].value)
            sum = sum + x
        }
        previousTokens = currentTokens
        previousPartTokens = currentPartTokens
    }
    return sum
}

func ParseLine(input string) []token {
    splitLine := strings.Split(input, ".")
    var tokens []token
    var inputIndexCounter int
    var currentDigitSequence string
    var currentDigitIndex int
    var currentDigitIndexSet bool
    for _, t := range splitLine {
        if t != "" {
            if _, err := strconv.Atoi(t); err == nil {
                token := token{value: t, index: inputIndexCounter, isDigit: true}
                tokens = append(tokens, token)
                inputIndexCounter = inputIndexCounter + len(t)
            } else {
                for _, rune := range t {
                    if unicode.IsDigit(rune) {
                        currentDigitSequence += string(rune)
                        if !currentDigitIndexSet {
                            currentDigitIndex = inputIndexCounter
                            currentDigitIndexSet = true
                        }
                        inputIndexCounter = inputIndexCounter + 1
                    } else {
                        if currentDigitSequence != "" {
                            tokens = append(tokens, token{value: currentDigitSequence, index: currentDigitIndex, isDigit: true})
                            currentDigitSequence = ""
                            currentDigitIndex = 0
                            currentDigitIndexSet = false
                        }
                        tokens = append(tokens, token{value: string(rune), index: inputIndexCounter, isDigit: false})
                        inputIndexCounter = inputIndexCounter + 1
                    }
                }
            }
        }

        if t == "" && currentDigitSequence != "" {
            tokens = append(tokens, token{value: currentDigitSequence, index: currentDigitIndex, isDigit: true})
            currentDigitSequence = ""
            currentDigitIndex = 0
            currentDigitIndexSet = false
        }
        inputIndexCounter = inputIndexCounter + 1
    }
    return tokens
}

func ParsePartNumbers(previousTokens, currentTokens, ignoreTokens []token) []token {
    partNumberTokens := []token{}

    for _, previousToken := range previousTokens {
        var storedCurrentToken token
        for _, currentToken := range currentTokens {
            if previousToken.isDigit {
                if !currentToken.isDigit {
                    if previousToken.index >= (currentToken.index-len(previousToken.value)) &&
                        previousToken.index <= (currentToken.index+1) &&
                        !slices.Contains(ignoreTokens, previousToken) {
                        partNumberTokens = append(partNumberTokens, previousToken)
                    }
                }
            }

            if !previousToken.isDigit && previousToken.value != "" {
                if currentToken.index >= (previousToken.index-len(currentToken.value)) &&
                    currentToken.index <= (previousToken.index+1) {
                    partNumberTokens = append(partNumberTokens, currentToken)
                    ignoreTokens = append(ignoreTokens, currentToken)
                }
            }

            if currentToken.isDigit {
                if !slices.Contains(ignoreTokens, currentToken) &&
                    !storedCurrentToken.isDigit &&
                    (currentToken.index == storedCurrentToken.index-1 ||
                        currentToken.index == storedCurrentToken.index+1) {
                    partNumberTokens = append(partNumberTokens, currentToken)
                    ignoreTokens = append(ignoreTokens, currentToken)
                    storedCurrentToken = token{}
                } else {
                    storedCurrentToken = currentToken
                }
            } else {
                if !reflect.DeepEqual(storedCurrentToken, token{}) {
                    if storedCurrentToken.index >= (currentToken.index-len(storedCurrentToken.value)) &&
                        storedCurrentToken.index <= (currentToken.index+1) {
                        partNumberTokens = append(partNumberTokens, storedCurrentToken)
                        storedCurrentToken = token{}
                    }
                }
                storedCurrentToken = currentToken
            }
        }
    }

    return partNumberTokens
}
