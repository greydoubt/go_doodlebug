// main.go
package main

import (
    "fmt"
    "math/rand"
    "time"
    "./doodlebug"
    "./grid"
    "./ant"
)



func main() {
    rand.Seed(time.Now().UnixNano())

    width, height := 10, 10
    grid := NewGrid(width, height)

    // Initialize the grid with organisms (ants and doodlebugs)

    // Create ants
    for i := 0; i < 10; i++ {
        x := rand.Intn(width)
        y := rand.Intn(height)
        ant := NewAnt()
        grid.AddOrganism(x, y, ant)
    }

    // Create doodlebugs
    for i := 0; i < 5; i++ {
        x := rand.Intn(width)
        y := rand.Intn(height)
        doodlebug := NewDoodlebug()
        grid.AddOrganism(x, y, doodlebug)
    }

    for {
        // Simulate a step of the ecosystem
        // Move organisms
        for y := 0; y < height; y++ {
            for x := 0; x < width; x++ {
                organism := grid.GetOrganism(x, y)
                if organism != nil {
                    organism.Move(grid)
                }
            }
        }

        // Breed organisms
        for y := 0; y < height; y++ {
            for x := 0; x < width; x++ {
                organism := grid.GetOrganism(x, y)
                if organism != nil {
                    organism.Breed(grid)
                }
            }
        }

        // Update the grid
        fmt.Println("**************************************")
        grid.Print()
        fmt.Println("**************************************")
        time.Sleep(500 * time.Millisecond)
    }
}
