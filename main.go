package main

import (
	"flag"
	"fmt"

	"github.com/boodyvo/tendermint/world"
)

const (
	MaxIterationsNumber = 10000
)

func main() {
	aliensNumber := flag.Int("n", 0, "number of aliens")
	inputPath := flag.String("input_path", "input.txt", "path to input file")
	flag.Parse()

	fmt.Printf("Creating world from file %s with %d aliens\n", *inputPath, *aliensNumber)

	worldMap, err := world.ReadMapFromFile(*inputPath)
	if err != nil {
		//log.Fatal("cannot read file that describe city map", err)
		fmt.Printf("cannot read file that describe city map: %v\n", err)
		return
	}

	worldX := world.NewWorld(worldMap, *aliensNumber)

	for i := 0; i < MaxIterationsNumber; i++ {
		if worldX.HasAliveAliens() {
			worldX.MakeIteration()

			continue
		}

		break
	}

	worldX.PrintCities()
}
