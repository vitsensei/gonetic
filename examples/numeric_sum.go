package examples

import (
	"github.com/vitsensei/gonetic/genetic"
	"math"
	"math/rand"
)

type Number struct {
	number int
}

func (n *Number) Mutation() {
	if rand.Intn(2) == 0 {
		n.number += rand.Intn(25)
	} else {
		n.number -= rand.Intn(25)
	}
}

type Array struct {
	array []Number
	sum   int
}

func (a *Array) Evaluate() float64 {
	sum := 0

	for i := range a.array {
		sum += a.array[i].number
	}

	return -math.Abs(float64(a.sum) - float64(sum))
}

func (a *Array) Len() int {
	return len(a.array)
}

func (a *Array) Mutation() {
	a.array[rand.Intn(len(a.array))].Mutation()
}

func (a Array) CreateCopy() genetic.Chromosome {
	var anotherArray Array
	anotherArray.array = make([]Number, len(a.array))
	anotherArray.sum = a.sum
	for i := range a.array {
		anotherArray.array[i] = Number{number: a.array[i].number}
	}

	return &anotherArray
}

func (a *Array) Genes() *[]genetic.Gene {
	genes := make([]genetic.Gene, len(a.array))

	for i := 0; i < len(a.array); i++ {
		genes[i] = &a.array[i]
	}

	return &genes
}

func GenerateRandomPopulation(nPopulation, sum int) []genetic.Chromosome {
	// Generate 100 array of 10 random integer ranging from -50 to 50
	chromosomes := make([]genetic.Chromosome, nPopulation)

	for i := 0; i < nPopulation; i++ {
		var array Array
		for j := 0; j < 10; j++ {
			array.array = append(array.array, Number{number: rand.Intn(101) - 50})
		}
		array.sum = sum

		chromosomes[i] = &array
	}

	return chromosomes
}
