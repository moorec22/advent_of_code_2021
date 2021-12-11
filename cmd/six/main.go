package main

import (
        "fmt"
        "os"
        "strconv"
        "strings"
)

// the delay for internal clock start
const Delay = 2
// a fish will have an offspring after Rate days
const Rate = 7

func main() {
        fish := readFile("input.txt")
        fmt.Println(fishAfterDays(fish, 256))
}

func fishAfterDays(fish []int, days int) int {
        buckets := make([]int, Rate + Delay)
        total := len(fish)
        for _, f := range fish {
                buckets[f]++
        }
        for i := 0; i < days; i++ {
                generation := buckets[i % len(buckets)]
                total += generation
                buckets[(i + Rate) % len(buckets)] += generation
        }
        return total
}

func check(e error) {
        if e != nil {
                panic(e)
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

