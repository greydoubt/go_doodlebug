// ant.go
package main

import (
	"math/rand"
)

type Ant struct{}

func NewAnt() *Ant {
	return &Ant{}
}

func (a *Ant) Symbol() string {
	return AntCell
}

func (a *Ant) Move(g *Grid) {
	dx := rand.Intn(3) - 1
	dy := rand.Intn(3) - 1

	x, y := g.FindOrganism(a)
	newX := x + dx
	newY := y + dy

	if g.IsValidCoordinate(newX, newY) && g.IsEmpty(newX, newY) {
		// Move the ant to the new position
		g.MoveOrganism(x, y, newX, newY)
	}
}

func (a *Ant) Breed(g *Grid) {
	x, y := g.FindOrganism(a)

	if g.HasAdjacentEmptyCell(x, y) {
		// Choose a random adjacent empty cell
		adjX, adjY := g.GetRandomAdjacentEmptyCell(x, y)

		// Create a new ant in the chosen cell
		newAnt := NewAnt()
		g.AddOrganism(adjX, adjY, newAnt)
	}
}
