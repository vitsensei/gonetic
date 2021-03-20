package main

import (
	"fmt"
	"github.com/vitsensei/gonetic/examples"
	"github.com/vitsensei/gonetic/genetic"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	chromosomes := examples.GenerateRandomPopulation(25, 500)

	population, _ := genetic.CreatePopulation(10, chromosomes, 100, 0.5)
	population.Populate()

	array := population.GetSample(1)

	fmt.Println("The best sample is: ", population.GetSample(1))
	fmt.Println("The score of that sample is: ", array.Evaluate())
}
