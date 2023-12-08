package daythree

import (
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
        t.Errorf(`Test '%v': expected '%v', got '%v'`, "TestProcess", expected, actual)
    }
}

func TestProcess_TopLeftCorner(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_TopLeftCorner symbol right before": {
            input: []string{
                "+627.....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_TopLeftCorner symbol right after": {
            input: []string{
                "627+....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_TopLeftCorner symbol below on edge": {
            input: []string{
                "627.....",
                "+.......",
            },
            expected: 627,
        },
        "TestProcess_TopLeftCorner symbol below": {
            input: []string{
                "627.....",
                ".+......",
            },
            expected: 627,
        },
        "TestProcess_TopLeftCorner symbol below at edge": {
            input: []string{
                "627.....",
                "...+....",
            },
            expected: 627,
        },
        "TestProcess_TopLeftCorner not a part number": {
            input: []string{
                "627.....",
                "........",
            },
            expected: 0,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_TopRightCorner(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_TopRightCorner symbol right before": {
            input: []string{
                "....+627",
                "........",
            },
            expected: 627,
        },
        "TestProcess_TopRightCorner symbol right after": {
            input: []string{
                "....627+",
                "........",
            },
            expected: 627,
        },
        "TestProcess_TopRightCorner symbol below on edge": {
            input: []string{
                ".....627",
                ".......+",
            },
            expected: 627,
        },
        "TestProcess_TopRightCorner symbol below": {
            input: []string{
                ".....627",
                "......+.",
            },
            expected: 627,
        },
        "TestProcess_TopRightCorner symbol below at edge": {
            input: []string{
                ".....627",
                "....+...",
            },
            expected: 627,
        },
        "TestProcess_TopRightCorner not a part number": {
            input: []string{
                ".....627",
                "........",
            },
            expected: 0,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_BottomLeftCorner(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_BottomLeftCorner symbol right before": {
            input: []string{
                "........",
                "+627....",
            },
            expected: 627,
        },
        "TestProcess_BottomLeftCorner symbol right after": {
            input: []string{
                "........",
                "627+....",
            },
            expected: 627,
        },
        "TestProcess_BottomLeftCorner symbol above on edge": {
            input: []string{
                "+.......",
                "627.....",
            },
            expected: 627,
        },
        "TestProcess_BottomLeftCorner symbol above": {
            input: []string{
                ".+......",
                "627.....",
            },
            expected: 627,
        },
        "TestProcess_BottomLeftCorner symbol above at edge": {
            input: []string{
                "...+....",
                "627.....",
            },
            expected: 627,
        },
        "TestProcess_BottomLeftCorner not a part number": {
            input: []string{
                "........",
                "627.....",
            },
            expected: 0,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_BottomRightCorner(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_BottomRightCorner symbol right before": {
            input: []string{
                "........",
                "....+627",
            },
            expected: 627,
        },
        "TestProcess_BottomRightCorner symbol right after": {
            input: []string{
                "........",
                "....627+",
            },
            expected: 627,
        },
        "TestProcess_BottomRightCorner symbol above": {
            input: []string{
                "......+.",
                ".....627",
            },
            expected: 627,
        },
        "TestProcess_BottomRightCorner symbol above on edge": {
            input: []string{
                ".......+",
                ".....627",
            },
            expected: 627,
        },
        "TestProcess_BottomRightCorner symbol above at edge": {
            input: []string{
                "....+...",
                ".....627",
            },
            expected: 627,
        },
        "TestProcess_BottomRightCorner not a part number": {
            input: []string{
                "........",
                ".....627",
            },
            expected: 0,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_LeftEdge(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_LeftEdge symbol right before": {
            input: []string{
                "........",
                "+627....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge symbol right after": {
            input: []string{
                "........",
                "627+....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge symbol below on edge": {
            input: []string{
                "........",
                "627.....",
                "+.......",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge symbol below": {
            input: []string{
                "........",
                "627.....",
                ".+......",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge symbol below at edge": {
            input: []string{
                "........",
                "627.....",
                "...+....",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge symbol above on edge": {
            input: []string{
                "+.......",
                "627.....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge symbol above": {
            input: []string{
                ".+......",
                "627.....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge symbol above at edge": {
            input: []string{
                "...+....",
                "627.....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge not a part number": {
            input: []string{
                "........",
                "627.....",
                "........",
            },
            expected: 0,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_RightEdge(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_RightEdge symbol right before": {
            input: []string{
                "........",
                "....627+",
                "........",
            },
            expected: 627,
        },
        "TestProcess_RightEdge symbol right after": {
            input: []string{
                "........",
                "....+627",
                "........",
            },
            expected: 627,
        },
        "TestProcess_RightEdge symbol below on edge": {
            input: []string{
                "........",
                ".....627",
                ".......+",
            },
            expected: 627,
        },
        "TestProcess_RightEdge symbol below": {
            input: []string{
                "........",
                ".....627",
                "......+.",
            },
            expected: 627,
        },
        "TestProcess_RightEdge symbol below at edge": {
            input: []string{
                "........",
                ".....627",
                "....+...",
            },
            expected: 627,
        },
        "TestProcess_RightEdge symbol above on edge": {
            input: []string{
                ".......+",
                ".....627",
                "........",
            },
            expected: 627,
        },
        "TestProcess_RightEdge symbol above": {
            input: []string{
                "......+.",
                ".....627",
                "........",
            },
            expected: 627,
        },
        "TestProcess_RightEdge symbol above at edge": {
            input: []string{
                "....+...",
                ".....627",
                "........",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge not a part number": {
            input: []string{
                "........",
                ".....627",
                "........",
            },
            expected: 0,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_TopEdge(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_TopEdge symbol right before": {
            input: []string{
                "..+627..",
                "........",
                "........",
            },
            expected: 627,
        },
        "TestProcess_TopEdge symbol right after": {
            input: []string{
                "..627+..",
                "........",
                "........",
            },
            expected: 627,
        },
        "TestProcess_TopEdge symbol below on edge": {
            input: []string{
                "..627...",
                "..+.....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_TopEdge symbol below": {
            input: []string{
                "..627...",
                "...+....",
                "........",
            },
            expected: 627,
        },
        "TestProcess_TopEdge symbol below at edge": {
            input: []string{
                "..627...",
                ".....+..",
                "........",
            },
            expected: 627,
        },
        "TestProcess_LeftEdge not a part number": {
            input: []string{
                "........",
                "627.....",
                "........",
            },
            expected: 0,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_BottomEdge(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_BottomEdge symbol right before": {
            input: []string{
                "........",
                "........",
                "..+627..",
            },
            expected: 627,
        },
        "TestProcess_BottomEdge symbol right after": {
            input: []string{
                "........",
                "........",
                "..627+..",
            },
            expected: 627,
        },
        "TestProcess_BottomEdge symbol below on edge": {
            input: []string{
                "........",
                "..+.....",
                "..627...",
            },
            expected: 627,
        },
        "TestProcess_BottomEdge symbol below": {
            input: []string{
                "........",
                "...+....",
                "..627...",
            },
            expected: 627,
        },
        "TestProcess_BottomEdge symbol below at edge": {
            input: []string{
                "........",
                ".....+..",
                "..627...",
            },
            expected: 627,
        },
        "TestProcess_BottomEdge not a part number": {
            input: []string{
                "........",
                "627.....",
                "........",
            },
            expected: 0,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_Middle(t *testing.T) {
    tests := map[string]struct {
        input    []string
        expected int
    }{
        "TestProcess_Middle symbol right before": {
            input: []string{
                "........",
                "..627+..",
                "........",
            },
            expected: 627,
        },
        "TestProcess_Middle symbol right after": {
            input: []string{
                "........",
                "..+627..",
                "........",
            },
            expected: 627,
        },
        "TestProcess_Middle symbol below on edge": {
            input: []string{
                "........",
                "...627..",
                ".....+..",
            },
            expected: 627,
        },
        "TestProcess_Middle symbol below": {
            input: []string{
                "........",
                "...627..",
                "....+...",
            },
            expected: 627,
        },
        "TestProcess_Middle symbol below at edge": {
            input: []string{
                "........",
                "...627..",
                "..+.....",
            },
            expected: 627,
        },
        "TestProcess_Middle symbol above on edge": {
            input: []string{
                ".....+..",
                "...627..",
                "........",
            },
            expected: 627,
        },
        "TestProcess_Middle symbol above": {
            input: []string{
                "....+...",
                "...627..",
                "........",
            },
            expected: 627,
        },
        "TestProcess_Middle symbol above at edge": {
            input: []string{
                "..+.....",
                "...627..",
                "........",
            },
            expected: 627,
        },
        "TestProcess_Middle not a part number": {
            input: []string{
                "........",
                "...627..",
                "........",
            },
            expected: 0,
        },
        "TestProcess_Middle sample": {
            input: []string{
                "...*......",
                "..35..633.",
                "......#...",
            },
            expected: 633 + 35,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := Process(test.input)
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}

func TestProcess_isSymbol(t *testing.T) {
    tests := map[string]struct {
        input    []rune
        expected bool
    }{
        "isSymbol a number": {
            input:    []rune("1"),
            expected: false,
        },
        "isSymbol a dot": {
            input:    []rune("."),
            expected: false,
        },
        "isSymbol a symbol": {
            input:    []rune("+"),
            expected: true,
        },
    }
    for name, test := range tests {
        t.Run(name, func(t *testing.T) {
            actual := isSymbol(test.input[0])
            if actual != test.expected {
                t.Errorf(`Test '%v': expected '%v', got '%v'`, name, test.expected, actual)
            }
        })
    }
}
