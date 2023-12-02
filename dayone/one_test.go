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
    input := []string{"1abc2","pqr3stu8vwx","a1b2c3d4e5f","treb7uchet","asdfasdt3eaadsfasd","abd"}
    expected := 175
    test(input, expected, t)
}

func TestProcessStringWithWordDigits(t *testing.T) {
    input := []string{
        "two1nine",
        "eightwothree",
        "abcone2threexyz",
        "xtwone3four",
        "4nineeightseven2",
        "zoneight234",
        "7pqrstsixteen",
    }
    expected := 281
    test(input, expected, t)
}

func TestReplaceWordNumbersForDigits(t *testing.T) {
    m := make(map[string]string)
    m["one5hello7"] = "15hello7"
    m["two5hello7"] = "25hello7"
    m["three5hello7"] = "35hello7"
    m["four5hello7"] = "45hello7"
    m["five6hello7"] = "56hello7"
    m["six5hello7"] = "65hello7"
    m["seven5hello7"] = "75hello7"
    m["eight5hello7"] = "85hello7"
    m["nine5hello7"] = "95hello7"
    m["rtsevenfourfive1rqhslone"] = "rt7451rqhsl1"
    m["plckgsixeight6eight95bnrfonetwonej"] = "plckg686895bnrf121j"
    m["eighthree"] = "83"
    m["sevenine"] = "79"
    m["asdfasdthreeaadsfasd"] = "asdfasd3aadsfasd"

    for input, expected := range m {
        actual := ReplaceWordNumbersForDigits(input)
        if actual != expected {
            t.Fatalf(`got %v, expected %v`, actual, expected)
        }
    }
}

func test(input []string, expected int, t *testing.T) {
    actual := Process(input)
    if actual != expected {
        t.Fatalf(`got %v, expected %v`, actual, expected)
    }
}