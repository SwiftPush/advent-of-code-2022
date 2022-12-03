package day02

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registry.Registry[2] = main
}

type Row struct {
	opponentMove, playerMove string
}

func readInput(filename string) []Row {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputSplit := strings.Split(inputText, "\n")

	rows := lo.Map(inputSplit, func(s string, _ int) Row {
		elems := strings.Split(s, " ")
		return Row{opponentMove: elems[0], playerMove: elems[1]}
	})

	return rows
}

func main() {
	filename := utils.ParseCommandLineArguments()
	rows := readInput(filename)

	part1Scores := []int{}
	for _, row := range rows {
		part1Scores = append(part1Scores, part1CalcScore(row))
	}
	part1 := lo.Sum(part1Scores)
	fmt.Println("part1", part1)

	part2Scores := []int{}
	for _, row := range rows {
		part2Scores = append(part2Scores, part2CalcScore(row))
	}
	part2 := lo.Sum(part2Scores)
	fmt.Println("part2", part2)
}

func part1CalcScore(r Row) int {
	scoreForPlayerMove := 0
	switch r.playerMove {
	case "X":
		scoreForPlayerMove = 1
	case "Y":
		scoreForPlayerMove = 2
	case "Z":
		scoreForPlayerMove = 3
	}

	outcomeScore := 0
	switch {
	case r.opponentMove == "A" && r.playerMove == "X": // draw
		outcomeScore = 3
	case r.opponentMove == "B" && r.playerMove == "Y":
		outcomeScore = 3
	case r.opponentMove == "C" && r.playerMove == "Z":
		outcomeScore = 3

	case r.opponentMove == "A" && r.playerMove == "Y": // win
		outcomeScore = 6
	case r.opponentMove == "A" && r.playerMove == "Z": // lose

	case r.opponentMove == "B" && r.playerMove == "X": // lose
	case r.opponentMove == "B" && r.playerMove == "Z": // win
		outcomeScore = 6

	case r.opponentMove == "C" && r.playerMove == "X": // win
		outcomeScore = 6
	case r.opponentMove == "C" && r.playerMove == "Y": // lose
	}

	return scoreForPlayerMove + outcomeScore
}

func part2CalcScore(r Row) int {
	// x - we lose
	// y - we draw
	// z - we win

	switch {
	case r.opponentMove == "A" && r.playerMove == "X":
		// we play scissors to lose
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "Z"})
	case r.opponentMove == "A" && r.playerMove == "Y":
		// we play rock to draw
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "X"})
	case r.opponentMove == "A" && r.playerMove == "Z":
		// we play paper to win
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "Y"})

	case r.opponentMove == "B" && r.playerMove == "X":
		// we play rock to lose
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "X"})
	case r.opponentMove == "B" && r.playerMove == "Y":
		// we play paper to draw
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "Y"})
	case r.opponentMove == "B" && r.playerMove == "Z":
		// we play scissors to win
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "Z"})

	case r.opponentMove == "C" && r.playerMove == "X":
		// we play paper to lose
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "Y"})
	case r.opponentMove == "C" && r.playerMove == "Y":
		// we play scissors to draw
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "Z"})
	case r.opponentMove == "C" && r.playerMove == "Z":
		// we play rock to win
		return part1CalcScore(Row{opponentMove: r.opponentMove, playerMove: "X"})
	}

	return -1
}

//a x - rock
//b y - paper
//c z - scissors
