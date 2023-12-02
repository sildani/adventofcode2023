package daytwo

import (
    "testing"
    "reflect"
)

func TestProcess(t *testing.T) {
    input := []string{
        "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
        "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
        "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
        "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
        "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
    }
    expectedPartOne, expectedPartTwo := 8, 2286
    actualPartOne, actualPartTwo := Process(input)
    if actualPartOne != expectedPartOne {
        t.Fatalf(`got %v, expected %v`, actualPartOne, expectedPartOne)
    }
    if actualPartTwo != expectedPartTwo {
        t.Fatalf(`got %v, expected %v`, actualPartTwo, expectedPartTwo)
    }
}

func TestParseGame(t *testing.T) {
    input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
    // my first iteration assumed new cubes were pulled out of the bag each
    // time but that is not correct!! Let's fix that.
    // old expectation
    // expected := game{id: 1, cubes: cubes{blue: 9, red: 5, green: 4}}
    // new expectation:
    expected := game{id: 1, cubes: cubes{blue: 6, red: 4, green: 2}}
    actual := ParseGame(input)
    if !reflect.DeepEqual(actual, expected) {
        t.Fatalf(`got %v, expected %v`, actual, expected)
    }
}