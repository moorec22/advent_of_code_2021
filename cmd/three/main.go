package main

import (
        "fmt"
        "math"
        "os"
        "strconv"
        "strings"
)

func main() {
        binaries := readFile("input.txt")
        fmt.Println(calculatePowerConsumption(binaries))
        fmt.Println(calculateLifeSupportRating(binaries))
}

func calculatePowerConsumption(binaries []string) int {
        if len(binaries) == 0 {
                return 0
        }
        frequencies := frequencies(binaries)
        gamma := float64(0)
        epsilon := float64(0)
        for i, v := range frequencies {
                if v > 0 {
                        gamma += math.Pow(2, float64(len(frequencies) - i - 1))
                } else if v < 0 {
                        epsilon += math.Pow(2, float64(len(frequencies) - i - 1))
                }
        }
        return int(gamma * epsilon)
}

func calculateLifeSupportRating(binaries []string) int {
        oxygenGeneratorBinaries := make([]string, len(binaries))
        copy(oxygenGeneratorBinaries, binaries)

        for i := range binaries[0] {
                frequencies := frequencies(oxygenGeneratorBinaries)
                if len(oxygenGeneratorBinaries) == 1 {
                        break
                }
                newOxygen := make([]string, 0)
                for _, v := range oxygenGeneratorBinaries {
                        if (frequencies[i] < 0 && v[i] == '0') || (frequencies[i] >= 0 && v[i] == '1') {
                                newOxygen = append(newOxygen, v)
                        }
                }
                oxygenGeneratorBinaries = newOxygen
        }

        co2ScrubberBinaries := make([]string, len(binaries))
        copy(co2ScrubberBinaries, binaries)

        for i := range binaries[0] {
                frequencies := frequencies(co2ScrubberBinaries)
                if len(co2ScrubberBinaries) == 1 {
                        break
                }
                newCo2 := make([]string, 0)
                for _, v := range co2ScrubberBinaries {
                        if (frequencies[i] < 0 && v[i] == '1') || (frequencies[i] >= 0 && v[i] == '0') {
                                newCo2 = append(newCo2, v)
                        }
                }
                co2ScrubberBinaries = newCo2
        }

        oxygenGeneratorRating, e := strconv.ParseInt(oxygenGeneratorBinaries[0], 2, 64)
        check(e)
        co2ScrubberRating, e := strconv.ParseInt(co2ScrubberBinaries[0], 2, 64)
        check(e)
        return int(oxygenGeneratorRating * co2ScrubberRating)
}

func frequencies(binaries []string) []int {
        frequencies := make([]int, len(binaries[0]))
        for _, v := range binaries {
                for i, c := range v {
                        if c == '0' {
                                frequencies[i]--
                        } else {
                                frequencies[i]++
                        }
                }
        }
        return frequencies
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func readFile(fileName string) []string {
        data, err := os.ReadFile(fileName)
        check(err)
        trimmedData := strings.TrimSuffix(string(data), "\n")
        return strings.Split(trimmedData,"\n")
}

