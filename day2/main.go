package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cubes = map[string]int{ 
	"red": 12,
	"green": 13,
	"blue": 14,
}

func main() {
	input := inputFromFile("input")
	games := strings.Split(string(input), "\n")

	task1(games)
}

func task1(games []string) {
	var tot int
	for _, game := range games {
		g := strings.Split(strings.ReplaceAll(game, "Game ", ""), ":")
		gid, _ := strconv.Atoi(g[0])
		sets := parseSets(strings.TrimSpace(g[1]))

		if checkGamePossible(sets) {
			tot += gid
		}
	}

	fmt.Println(tot)
}

func checkGamePossible(sets []map[string]int) bool {
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

func inputFromFile(filename string) []byte {
	binp, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	return binp
}