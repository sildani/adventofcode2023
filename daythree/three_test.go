package daythree

import (
    "reflect"
    "testing"
)

func TestProcess(t *testing.T) {
    input := []string{
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
    }
    expected := 4361
    actual := Process(input)
    if actual != expected {
        t.Fatalf(`got %v, expected %v`, actual, expected)
    }
}

func TestParseLine(t *testing.T) {
    inputs := []string{
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
        "...*44....",
    }
    expected := [][]token{
        {
            { value: "467", index: 0, isDigit: true  },
            { value: "114", index: 5, isDigit: true  },
        },
        {
            { value: "*",   index: 3, isDigit: false },
        },
        {
            { value: "35",  index: 2, isDigit: true  },
            { value: "633", index: 6, isDigit: true  },
        },
        {
            { value: "#",   index: 6, isDigit: false },
        },
        {
            { value: "617", index: 0, isDigit: true  },
            { value: "*",   index: 3, isDigit: false },
        },
        {
            { value: "+",   index: 5, isDigit: false },
            { value: "58",  index: 7, isDigit: true  },
        },
        {
            { value: "592", index: 2, isDigit: true  },
        },
        {
            { value: "755", index: 6, isDigit: true  },
        },
        {
            { value: "$",   index: 3, isDigit: false },
            { value: "*",   index: 5, isDigit: false },
        },
        {
            { value: "664", index: 1, isDigit: true  },
            { value: "598", index: 5, isDigit: true  },
        },
        {
            { value: "*",   index: 3, isDigit: false },
            { value: "44",  index: 4, isDigit: true  },
        },
    }
    for i := range inputs {
        actual := ParseLine(inputs[i])
        if !reflect.DeepEqual(actual, expected[i]) {
            t.Fatalf(`got %v, expected %v`, actual, expected[i])
        }
    }
}

// write a test for a function that takes two lines (current and previous)
// and checks for valid numbers in the current line. previous will be nil
// for first line. test for not marking the same number valid TWICE
func TestParsePartNumbers(t *testing.T) {
    inputs := [][][]token{
        {
            {
                {},
            },
            {
                { value: "467", index: 0, isDigit: true  },
                { value: "114", index: 5, isDigit: true  },
            },
            {},
        },
        {
            {
                { value: "467", index: 0, isDigit: true  },
                { value: "114", index: 5, isDigit: true  },
            },
            {
                { value: "*",   index: 3, isDigit: false },
            },
            {},
        },
        {
            {
                { value: "*",   index: 3, isDigit: false },
            },
            {
                { value: "35",  index: 2, isDigit: true  },
                { value: "633", index: 6, isDigit: true  },
            },
            {
                { value: "467", index: 0, isDigit: true  },
            },
        },
        {
            {
                { value: "35",  index: 2, isDigit: true  },
                { value: "633", index: 6, isDigit: true  },
            },
            {
                { value: "#",   index: 6, isDigit: false },
            },
            {
                { value: "35",  index: 2, isDigit: true  },
            },
        },
        {
            {
                { value: "#",   index: 6, isDigit: false },
            },
            {
                { value: "617", index: 0, isDigit: true  },
                { value: "*",   index: 3, isDigit: false },
            },
            {
                { value: "633", index: 6, isDigit: true  },
            },
        },
        {
            {
                { value: "617", index: 0, isDigit: true  },
                { value: "*",   index: 3, isDigit: false },
            },
            {
                { value: "+",   index: 5, isDigit: false },
                { value: "58",  index: 7, isDigit: true  },
            },
            {
                { value: "617", index: 0, isDigit: true  },
            },
        },
        {
            {
                { value: "+",   index: 5, isDigit: false },
                { value: "58",  index: 7, isDigit: true  },
            },
            {
                { value: "592", index: 2, isDigit: true  },
            },
            {},
        },
        {
            {
                { value: "592", index: 2, isDigit: true  },
            },
            {
                { value: "755", index: 6, isDigit: true  },
            },
            {
                { value: "592", index: 2, isDigit: true  },
            },
        },
        {
            {
                { value: "755", index: 6, isDigit: true  },
            },
            {
                { value: "$",   index: 3, isDigit: false },
                { value: "*",   index: 5, isDigit: false },
            },
            {},
        },
        {
            {
                { value: "$",   index: 3, isDigit: false },
                { value: "*",   index: 5, isDigit: false },
            },
            {
                { value: "664", index: 1, isDigit: true  },
                { value: "598", index: 5, isDigit: true  },
            },
            {
                { value: "755", index: 6, isDigit: true  },
            },
        },
        {
            {
                { value: "664", index: 1, isDigit: true  },
                { value: "598", index: 5, isDigit: true  },
            },
            {
                { value: "*",   index: 3, isDigit: false },
                { value: "44",  index: 4, isDigit: true  },
            },
            {
                { value: "664", index: 1, isDigit: true  },
                { value: "598", index: 5, isDigit: true  },
            },
        },
    }
    expected := [][]token{
        {},
        {
            { value: "467", index: 0, isDigit: true  },},
        {
            { value: "35",  index: 2, isDigit: true  },},
        {
            { value: "633", index: 6, isDigit: true  },},
        {
            { value: "617", index: 0, isDigit: true  },},
        {},
        {
            { value: "592", index: 2, isDigit: true  },},
        {},
        {
            { value: "755", index: 6, isDigit: true  },},
        {
            { value: "664", index: 1, isDigit: true  },
            { value: "598", index: 5, isDigit: true  },
        },
        {
            { value: "44",  index: 4, isDigit: true  },
        },
    }
    for i := range inputs {
        actual := ParsePartNumbers(inputs[i][0], inputs[i][1], inputs[i][2])
        if !reflect.DeepEqual(actual, expected[i]) {
            t.Fatalf(`got %v, expected %v`, actual, expected[i])
        }
    }
}
