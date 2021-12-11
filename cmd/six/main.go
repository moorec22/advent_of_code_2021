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

type Stack []int

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
        return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(e int) {
        *s = append(*s, e) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (int, bool) {
        if s.IsEmpty() {
                return 0, false
        } else {
                index := len(*s) - 1 // Get the index of the top most element.
                e := (*s)[index] // Index into the slice and obtain the element.
                *s = (*s)[:index] // Remove it from the stack by slicing it off.
                return e, true
        }
}

func main() {
        fish := readFile("sample.txt")
        fmt.Println(fishAfterDays(fish, 18))
}

func fishAfterDays(fish []int, days int) int {
        total := 0
        for _, f := range fish {
                total += offspring(f, days) + 1
        }
        return total
}

func offspring(startDay int, afterDay int) int {
        if startDay >= afterDay {
                return 0
        }
        children := (afterDay - startDay + Rate - 1) / Rate
        for i := 1; i <= children; i++ {
                children += offspring(startDay + Delay + (i * Rate), afterDay)
        }
        return children
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

