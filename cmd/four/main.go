package main

import (
        "fmt"
        "os"
        "strconv"
        "strings"
)

const mark_symbol string = "-1"

func main() {
        numbers, boards := readFile("input.txt")
        fmt.Println(winningBoardScore(boards, numbers))
        fmt.Println(losingBoardScore(boards, numbers))
}

func winningBoardScore(boards [][][]string, numbers []string) int {
        for _, v := range numbers[:5] {
                boards = mark(boards, v)
        }
        for _, v := range numbers[5:] {
                boards = mark(boards, v)
                for _, board := range boards {
                        if winningBoard(board) {
                                val, e := strconv.Atoi(v)
                                check(e)
                                return boardScore(board) * val
                        }
                }
        }
        return 0
}

func losingBoardScore(boards [][][]string, numbers []string) int {
        for _, v := range numbers[:5] {
                boards = mark(boards, v)
        }
        for _, v := range numbers[5:] {
                boards = mark(boards, v)
                newBoards := make([][][]string, 0)
                for _, board := range boards {
                        if !winningBoard(board) {
                                newBoards = append(newBoards, board)
                        }
                }
                if len(newBoards) == 0 {
                        val, e := strconv.Atoi(v)
                        check(e)
                        return boardScore(boards[0]) * val
                }
                boards = newBoards
        }
        return 0
}

// returns true if and only if this board has won bingo
func winningBoard(board [][]string) bool {
        // horizontal check
        for i := 0; i < 5; i++ {
                horizontal := true
                vertical := true
                for j := 0; j < 5; j++ {
                        // horizontal check
                        if board[i][j] != mark_symbol {
                                horizontal = false
                        }
                        if board[j][i] != mark_symbol {
                                vertical = false
                        }
                }
                if horizontal || vertical {
                        return true
                }
                horizontal = true
                vertical = true
        }
        return false
}

func boardScore(board [][]string) int {
        score := 0
        for _, line := range board {
                for _, v := range line {
                        if v != mark_symbol {
                                val, e := strconv.Atoi(v)
                                check(e)
                                score += val
                        }
                }
        }
        return score
}

func mark(boards [][][]string, number string) [][][]string {
        for _, board := range boards {
                for _, line := range board {
                        for i, v := range line {
                                if v == number {
                                        line[i] = mark_symbol
                                }
                        }
                }
        }
        return boards
}

func check(e error) {
        if e != nil {
                panic(e)
        }
}

func readFile(fileName string) ([]string, [][][]string) {
        data, err := os.ReadFile(fileName)
        check(err)
        trimmedData := strings.TrimSuffix(string(data), "\n")
        stringData := strings.Split(trimmedData,"\n")
        boards := make([][][]string, 0)
        board := make([][]string, 0)
        for _, line := range stringData[2:] {
                if line == "" {
                        boards = append(boards, board)
                        board = make([][]string, 0)
                } else {
                        board = append(board, strings.Fields(line))
                }
        }
        boards = append(boards, board)
        return strings.Split(stringData[0], ","), boards
}

