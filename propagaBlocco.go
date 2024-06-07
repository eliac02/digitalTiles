package main

import "fmt"

func propagaBlocco(p piano, x, y int) {
	tile := piastrella{x: x, y: y}
	seen := make(map[piastrella]bool)
	block := esploraBlocco(p, tile, seen)

	originalBlock := make(map[piastrella]*properties)
	for t, props := range block {
		originalBlock[t] = &properties{
			color:          props.color,
			intensity:      props.intensity,
			parent:         props.parent,
			rank:           props.rank,
			blockIntensity: props.blockIntensity,
		}
	}

	for t := range block {
		colorCount := make(map[string]int)
		adiacenti := getAdiacenti(t.x, t.y)
		for _, adj := range adiacenti {
			if props, exists := originalBlock[adj]; exists {
				colorCount[props.color]++
			}
		}

		for index, reg := range *p.rules {
			if ruleOk(reg, colorCount) {
				fmt.Println(reg)
				p.tiles[t].color = reg.color
				(*p.rules)[index].usage++
				break
			}
		}

	}
}
