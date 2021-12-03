package main

import (
        "fmt"
        "os"
        "strconv"
        "strings"
)

type instruction struct {
        direction string
        distance int
}

func main() {
        instructions := readFile()
        fmt.Println(calculateMovement(instructions))
        fmt.Println(calculateMovementWithAim(instructions))
}

func calculateMovement(instructions []instruction) int {
        position := 0
        depth := 0
        for _, instruction := range instructions {
                if instruction.direction == "forward" {
                        position += instruction.distance
                } else if instruction.direction == "down" {
                        depth += instruction.distance
                } else if instruction.direction == "up" {
                        depth -= instruction.distance
                }
        }
        return position * depth
}

func calculateMovementWithAim(instructions []instruction) int {
        position := 0
        depth := 0
        aim := 0
        for _, instruction := range instructions {
                if instruction.direction == "forward" {
                        position += instruction.distance
                        depth += aim * instruction.distance
                } else if instruction.direction == "down" {
                        aim += instruction.distance
                } else if instruction.direction == "up" {
                        aim -= instruction.distance
                }
        }
        return position * depth
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func readFile() []instruction {
        data, err := os.ReadFile("./input.txt")
        check(err)
        trimmedData := strings.TrimSuffix(string(data), "\n")
        stringDirections := strings.Split(trimmedData,"\n")
        instructions := make([]instruction, len(stringDirections))
        for i, v := range stringDirections {
                parts := strings.Split(v, " ")
                distance, err := strconv.Atoi(parts[1])
                check(err)
                instructions[i] = instruction{direction: parts[0], distance: distance}
        }
        return instructions
}

