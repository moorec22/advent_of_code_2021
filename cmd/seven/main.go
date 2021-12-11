package main

import (
        "fmt"
        "math"
        "os"
        "sort"
        "strconv"
        "strings"
)

// the delay for internal clock start
const Delay = 2
// a fish will have an offspring after Rate days
const Rate = 7

func main() {
        positions := readFile("input.txt")
        fmt.Println(minLinearFuelConsumption(positions))
        fmt.Println(minPolyFuelConsumption(positions))
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func minLinearFuelConsumption(positions []int) int {
        m := median(positions)
        cost := 0
        for _, position := range positions {
                cost += abs(position - m)
        }
        return cost
}

func minPolyFuelConsumption(positions []int) int {
        a := average(positions)
        cost := 0
        for _, position := range positions {
                travel := abs(position - a)
                cost += travel * (travel + 1) / 2
        }
        return cost
}

func median(numbers []int) int {
        sort.Ints(numbers)
        return numbers[len(numbers) / 2]
}

func average(numbers []int) int {
        sum := 0
        for _, number := range numbers {
                sum += number
        }
        average := int(math.Round(float64(sum) / float64(len(numbers)))) - 1
        return average
}

func abs(number int) int {
        if number < 0 {
                return -number
        } else {
                return number
        }
}

func readFile(fileName string) []int {
        data, err := os.ReadFile(fileName)
        check(err)
        trimmedData := strings.TrimSuffix(string(data), "\n")
        stringData := strings.Split(trimmedData, ",")
        fish := make([]int, len(stringData))
        for i, v := range stringData {
                num, err := strconv.Atoi(v)
                check(err)
                fish[i] = num
        }
        return fish
}

