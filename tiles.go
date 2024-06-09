//Elia Cortesi 01911A

package main

import (
	"bufio"
	"os"
)

type piastrella struct {
	x int
	y int
}

type properties struct {
	color          string
	intensity      int
	parent         piastrella
	rank           int
	blockIntensity int
}

type ruleset struct {
    color string
    count int
}

type rule struct {
	raw     string
	ruleset []ruleset
	color   string
	usage   int
}

type piano struct {
	tiles map[piastrella]*properties
	rules *[]rule
}

func main() {
	p := makeSet()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		esegui(p, scanner.Text())
	}
}
