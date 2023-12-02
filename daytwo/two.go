package daytwo

import (
    "strconv"
    "strings"
)

type cubes struct {
    blue  int
    red   int
    green int
}

type game struct {
    id int
    cubes cubes
}

func Process(input []string) int {
    var sum int
    for _, line := range input {
        game := ParseGame(line)
        if game.cubes.red <= 12 && game.cubes.blue <= 14 && game.cubes.green <= 13 {
            sum = sum + game.id
        }
    }
    return sum
}

func ParseGame(input string) game {
    x := strings.Index(input, ":")
    gameId, _ := strconv.Atoi(input[5:x])
    var redCubes int
    var blueCubes int
    var greenCubes int

    cubesSets := strings.Split(input[x+1:], ";")
    for _, cubesSet := range cubesSets {
        cubesSet = strings.TrimSpace(cubesSet)
        cubes := strings.Split(cubesSet, ", ")
        for _, cube := range cubes {
            currentCount, _ := strconv.Atoi(cube[0:strings.Index(cube, " ")])
            currentColor := cube[strings.Index(cube, " ")+1:]
            switch currentColor {
            case "red":
                redCubes = redCubes + currentCount
            case "blue":
                blueCubes = blueCubes + currentCount
            case "green":
                greenCubes = greenCubes + currentCount
            default:
                // ignore
            }
        }
    }

    return game{
        id:    gameId,
        cubes: cubes{
            blue: blueCubes,
            red: redCubes,
            green: greenCubes,
        },
    }
}