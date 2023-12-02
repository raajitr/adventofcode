package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := inputFromFile("inputs/input")
	games := strings.Split(string(input), "\n")

	fmt.Println("Task 1:", task1(games))
	fmt.Println("Task 2:", task2(games))
}

func task1(games []string) int {
	var tot int
	for _, game := range games {
		gid, sets := parseGame(game)

		if checkGamePossible(sets) {
			tot += gid
		}
	}

	return tot
}

func task2(games []string) int {
	var tot int
	for _, game := range games {
		_, sets := parseGame(game)

		tot += minCubePower(sets)
	}

	return tot
}

func checkGamePossible(sets []map[string]int) bool {
	var cubes = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	isPossible := true

	for _, set := range sets {
		for k, v := range set {
			if v > cubes[k] {
				isPossible = false
			}
		}
	}
	return isPossible
}

func minCubePower(sets []map[string]int) int {
	min_cube := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, set := range sets {
		for k, v := range set {
			min_cube[k] = max(v, min_cube[k])
		}
	}

	power := 1

	for _, v := range min_cube {
		power *= v
	}

	return power
}

func parseSets(sets string) []map[string]int {
	var smap []map[string]int
	for _, set := range strings.Split(sets, "; ") {
		c_map := make(map[string]int)
		for _, cube := range strings.Split(set, ", ") {

			numc := strings.Split(cube, " ")
			num, _ := strconv.Atoi(numc[0])
			c := numc[1]

			c_map[c] = num
		}
		smap = append(smap, c_map)
	}

	return smap
}

func parseGame(game string) (int, []map[string]int) {
	g := strings.Split(strings.ReplaceAll(game, "Game ", ""), ":")
	gid, _ := strconv.Atoi(g[0])
	sets := parseSets(strings.TrimSpace(g[1]))

	return gid, sets
}

func inputFromFile(filename string) []byte {
	binp, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return binp
}
