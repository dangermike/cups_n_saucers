package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/dangermike/cups_n_saucers/board"
	"github.com/dangermike/cups_n_saucers/cupsaucer"
)

type boardPair struct {
	next  int
	board board.B
}

type result struct {
	solutions     []board.B
	solutionCount int
	backtracks    int
	checks        int
}

func main() {
	start := time.Now()
	res := backtracking()
	stop := time.Now()

	// _, w, _ := termSize()
	// printSolutions(res.solutions, w/(3*len(color.Colors)+1))

	fmt.Printf(
		"Found %d solutions with %d checks, %d backtracks in %0.0f μsecs\n",
		res.solutionCount,
		res.checks,
		res.backtracks,
		stop.Sub(start).Seconds()*100000,
	)
}

func backtracking() result {
	solutions := []board.B{}
	solutionCount := 0
	backtracks := 0
	checks := 0
	// unlike n-queens, we know that _something_ has to go in the top-left
	// corner. since all potential solutions are equivalent when the board
	// is empty, let's throw the first item in the first square to save time
	// startBoard := board.B{}
	// startBoard[0] = cupsaucer.All[0]
	// workQueue := []boardPair{boardPair{1, startBoard}}

	workQueue := []boardPair{boardPair{}}

	for len(workQueue) > 0 {
		bp := workQueue[len(workQueue)-1]
		workQueue = workQueue[:len(workQueue)-1]
		any := false
		for i := 0; i < len(bp.board); i++ {
			next, ok := bp.board.TryPlace(cupsaucer.All[bp.next], i)
			if ok {
				checks++
				any = true
				if bp.next == len(bp.board)-1 {
					solutions = append(solutions, next)
					solutionCount++
				} else {
					workQueue = append(workQueue, boardPair{bp.next + 1, next})
				}
			}
		}
		if !any {
			backtracks++
		}
	}
	return result{
		solutions,
		solutionCount,
		backtracks,
		checks,
	}
}

func printSolutions(solutions []board.B, width int) {
	w := bufio.NewWriter(os.Stdout)
	if width < 1 {
		width = 1
	}
	for i := 0; i < len(solutions); i += width {
		cnt := len(solutions) - i
		if cnt > width {
			cnt = width
		}
		for l := 0; l < 4; l++ {
			for s := 0; s < cnt; s++ {
				board := solutions[i+s]
				for c := 0; c < 4; c++ {
					if c > 0 {
						w.WriteString(" ")
					}
					w.WriteString(board[(4*l)+c].Cup().String())
					w.WriteString(board[(4*l)+c].Saucer().String())
				}
				w.WriteString("  ")
			}
			w.WriteString("\n")
		}
		w.WriteString("\n")
	}
	w.Flush()
}

func termSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	parts := strings.Split(string(out), " ")

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	y, err := strconv.Atoi(parts[1][:len(parts[1])-1])
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}
