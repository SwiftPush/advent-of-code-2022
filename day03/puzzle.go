package day03

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registry.Registry[3] = main
}

func readInput(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	return strings.Split(inputText, "\n")
}

func main() {
	filename := utils.ParseCommandLineArguments()
	lines := readInput(filename)

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	priorties := []int{}
	for _, line := range lines {
		lineLen := len(line)
		compartment1, compartment2 := []rune(line[:lineLen/2]), []rune(line[lineLen/2:])
		intersections := lo.Intersect(compartment1, compartment2)
		intersect := intersections[0]
		priority := getPriority(intersect)
		priorties = append(priorties, priority)
	}

	part1 := lo.Sum(priorties)
	fmt.Println("part1", part1)
}

func part2(lines []string) {
	priorites := []int{}
	groups := lo.Chunk(lines, 3)
	for _, group := range groups {
		foo := []rune(group[0])
		for _, line := range group[1:] {
			foo = lo.Intersect(foo, []rune(line))
		}
		priority := getPriority(foo[0])
		priorites = append(priorites, priority)
	}

	part2 := lo.Sum(priorites)
	fmt.Println("part2", part2)
}

func getPriority(r rune) int {
	if r >= 'a' {
		return int(r-'a') + 1
	}
	return int(r-'A') + 27
}
