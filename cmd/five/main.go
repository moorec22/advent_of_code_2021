package main

import (
        "fmt"
        "os"
        "strconv"
        "strings"
)

type coordinate struct {
        x int
        y int
}

type line struct {
        start coordinate
        end coordinate
}

func main() {
        lines := readFile("input.txt")
        fmt.Println(overlapCount(lines, horizontalAndVerticalLinesToPoints))
        fmt.Println(overlapCount(
                lines,
                func(l line) []coordinate {
                        return append(
                                horizontalAndVerticalLinesToPoints(l),
                                diagonalLinesToPoints(l)...,
                        )
                },
        ))
}

func overlapCount(lines []line, linesToPoints func(line) []coordinate) int {
        points := make(map[coordinate]int)
        for _, l := range lines {
                coordinates := linesToPoints(l)
                for _, coordinate := range coordinates {
                        if _, ok := points[coordinate]; !ok {
                                points[coordinate] = 0
                        }
                        points[coordinate]++
                }
        }
        overlaps := 0
        for _, count := range points {
                if count > 1 {
                        overlaps++
                }
        }
        return overlaps
}

func horizontalAndVerticalLinesToPoints(l line) []coordinate {
        coordinates := make([]coordinate, 0)
        if isHorizontal(l) || isVertical(l) {
                minX := min(l.start.x, l.end.x)
                maxX := max(l.start.x, l.end.x)
                minY := min(l.start.y, l.end.y)
                maxY := max(l.start.y, l.end.y)
                for i := minX; i <= maxX; i++ {
                        for j := minY; j <= maxY; j++ {
                                coordinates = append(coordinates, coordinate{x: i, y: j})
                        }
                }
        }
        return coordinates
}

func diagonalLinesToPoints(l line) []coordinate {
        coordinates := make([]coordinate, 0)
        if isDiagonal(l) {
                minX := min(l.start.x, l.end.x)
                maxX := max(l.start.x, l.end.x)
                minY := min(l.start.y, l.end.y)
                for i := 0; i <= maxX - minX; i++ {
                        if positiveSlope(l) {
                                coordinates = append(coordinates, coordinate{x: minX + i, y: minY + i})
                        } else {
                                coordinates = append(coordinates, coordinate{x: maxX - i, y: minY + i})
                        }
                }
        }
        return coordinates
}

func isHorizontal(l line) bool {
        return l.start.y == l.end.y
}

func isVertical(l line) bool {
        return l.start.x == l.end.x
}

func isDiagonal(l line) bool {
        return abs(l.start.x - l.end.x) == abs(l.start.y - l.end.y)
}

func positiveSlope(l line) bool {
        return (l.end.y - l.start.y) / (l.end.x - l.start.x) > 0
}

func getCoordinate(stringCoordinate string) coordinate {
        parts := strings.Split(stringCoordinate, ",")
        x, e := strconv.Atoi(parts[0])
        check(e)
        y, e := strconv.Atoi(parts[1])
        check(e)
        return coordinate{x: x, y: y}
}

func min(one int, two int) int {
        if one < two {
                return one
        } else {
                return two
        }
}

func max(one int, two int) int {
        if one > two {
                return one
        } else {
                return two
        }
}

func abs(num int) int {
        if num < 0 {
                return -num
        } else {
                return num
        }
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func readFile(fileName string) []line {
        data, err := os.ReadFile(fileName)
        check(err)
        trimmedData := strings.TrimSuffix(string(data), "\n")
        stringData := strings.Split(trimmedData,"\n")
        lines := make([]line, len(stringData))
        for i, stringDatum := range stringData {
                stringLine := strings.Fields(stringDatum)
                start := getCoordinate(stringLine[0])
                end := getCoordinate(stringLine[2])
                lines[i] = line{start: start, end: end}
        }
        return lines
}

