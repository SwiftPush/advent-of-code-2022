package day01

import (
	"aoc/registry"
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func init() {
	registry.Registry[1] = main
}

func readInput(filename string) [][]int {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	inputText := string(data)
	inputSplit := strings.Split(inputText, "\n")

	var res [][]int
	var currentElf []int

	for _, line := range inputSplit {
		if line == "" {
			res = append(res, currentElf)
			currentElf = []int{}
			continue
		}

		num, _ := strconv.Atoi(line)
		currentElf = append(currentElf, num)
	}

	return res
}

func main() {
	filename := utils.ParseCommandLineArguments()
	input := readInput(filename)

	totalSums := []int{}
	for _, elf := range input {
		totalSum := lo.Sum(elf)
		totalSums = append(totalSums, totalSum)
	}

	sort.Slice(totalSums, func(i, j int) bool {
		return totalSums[i] > totalSums[j]
	})

	part1 := totalSums[0]
	part2 := lo.Sum(totalSums[:3])

	fmt.Println("part1", part1)
	fmt.Println("part2", part2)
}
