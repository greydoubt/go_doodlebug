// grid.go
package main

import (
    "fmt"
)

type Grid struct {
    Width, Height int
    cells         [][]Organism
}

func NewGrid(width, height int) *Grid {
    cells := make([][]Organism, height)
    for i := range cells {
        cells[i] = make([]Organism, width)
    }
    return &Grid{Width: width, Height: height, cells: cells}
}

func (g *Grid) Print() {
    for y := 0; y < g.Height; y++ {
        for x := 0; x < g.Width; x++ {
            fmt.Print(g.cells[y][x].Symbol())
        }
        fmt.Println()
    }
}

func (g *Grid) AddOrganism(x, y int, org Organism) {
    g.cells[y][x] = org
}

func (g *Grid) GetOrganism(x, y int) Organism {
    return g.cells[y][x]
}
