// organism.go
package main

type Organism interface {
    Symbol() string
    Move(g *Grid)
    Breed(g *Grid)
}

const (
    EmptyCell = " "
    AntCell   = "o"
    DoodlebugCell = "X"
)
