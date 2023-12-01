package dayone

import (
	"testing"
)

func TestProcess(t *testing.T) {
    input := []string{"1abc2","pqr3stu8vwx","a1b2c3d4e5f","treb7uchet"}
    expected := 142
    test(input, expected, t)
}

func TestProcessStringWithNoDigits(t *testing.T) {
    input := []string{"1abc2","pqr3stu8vwx","a1b2c3d4e5f","treb7uchet","abd"}
    expected := 142
    test(input, expected, t)
}

func test(input []string, expected int, t *testing.T) {
    actual := Process(input)
    if actual != expected {
        t.Fatalf(`got %v, expected %v`, actual, expected)
    }
}