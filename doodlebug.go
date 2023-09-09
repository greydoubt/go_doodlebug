// doodlebug.go
package main

import (
    "math/rand"
)

type Doodlebug struct{}

func NewDoodlebug() *Doodlebug {
    return &Doodlebug{}
}

func (d *Doodlebug) Symbol() string {
    return DoodlebugCell
}

func (d *Doodlebug) Move(g *Grid) {
	x, y := g.FindOrganism(d)
	adjX, adjY := g.GetRandomAdjacentCell(x, y)

	if adjX != -1 && adjY != -1 {
		// Move the doodlebug to the adjacent cell and eat the ant
		g.MoveOrganism(x, y, adjX, adjY)
		g.RemoveOrganism(adjX, adjY)
	} else {
		// Move the doodlebug randomly to an empty cell
		dx := rand.Intn(3) - 1
		dy := rand.Intn(3) - 1

		newX := x + dx
		newY := y + dy

		if g.IsValidCoordinate(newX, newY) && g.IsEmpty(newX, newY) {
			g.MoveOrganism(x, y, newX, newY)
		}
	}
}

func (d *Doodlebug) Breed(g *Grid) {
	x, y := g.FindOrganism(d)

	if g.HasAdjacentEmptyCell(x, y) {
		adjX, adjY := g.GetRandomAdjacentEmptyCell(x, y)

		// Create a new doodlebug in the adjacent empty cell
		newDoodlebug := NewDoodlebug()
		g.AddOrganism(adjX, adjY, newDoodlebug)
	}
}
