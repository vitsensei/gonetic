package genetic

import (
	"errors"
	"math"
	"math/rand"
	"sort"
)

/* Copied from https://towardsdatascience.com/introduction-to-genetic-algorithms-including-example-code-e396e98d8bf3
START
Generate the initial population
Compute fitness
REPEAT
	Selection
	Crossover
	Mutation
	Compute fitness
UNTIL population has converged
STOP
*/

type Gene interface {
	Mutation()
}

/*
	- Chromosome stores an array of genes.
	- Evaluate() will calculate the score of the Chromosome
	- Mutation() will mutate the Chromosome. The expected behaviours is to
	iterate through the stored genes and calling Mutation() function
	- Len() will tell how many Genes this Chromosome supposed to have
*/
type Chromosome interface {
	Evaluate() float64
	Len() int
	Mutation()
	Genes() *[]Gene
	CreateCopy() Chromosome
}

/*
	Population create an easy interface for the user.
	To run GA, the end user should only call 2 function: CreatePopulation()
	and population.Run().
*/
type Population struct {
	nGene   int
	samples []Chromosome // this should not be a pointer since we don't
	// want to modify the original samples
	nSamples   int     // keep track of the original number of samples
	ratio      float64 // the percentage of population will be considered for reproduction
	nIteration int
}

func CreatePopulation(nGene int, samples []Chromosome, nIteration int, ratio float64) (Population, error) {
	notValidGeneLength := errors.New("not valid number of gene in chromosome")
	for _, s := range samples {
		if s.Len() != nGene {
			return Population{}, notValidGeneLength
		}
	}

	newPopulation := Population{
		nGene:      nGene,
		samples:    samples,
		nIteration: nIteration,
		ratio:      ratio,
		nSamples:   len(samples),
	}

	return newPopulation, nil
}

type IndPair struct {
	first  int
	second int
}

func (p *Population) Populate() {
	for i := 0; i < p.nIteration; i++ {
		indPairs := p.selection()

		for j := range indPairs {
			p.crossOver(indPairs[j].first, indPairs[j].second)
		}

		p.mutation()
		p.clean()
	}
}

func (p *Population) GetSample(ind int) Chromosome {
	sort.Sort(ByScore(p.samples))
	return p.samples[ind]
}

func (p *Population) selection() []IndPair {
	// Sorting the Chromosome array
	sort.Sort(ByScore(p.samples))

	// Calculate number of pairs of sample for reproduction
	nPair := int(math.Round(p.ratio * float64(p.nSamples) / 2))

	// Since only a subset of the population (the elitists, lol), it makes
	// sense to keep the population to be smaller than a certain number.
	// For the record, this is not the way I view the world.
	indPairArr := make([]IndPair, nPair)
	for i := 0; i < nPair; i++ {
		indPairArr[i] = IndPair{first: i * 2, second: i*2 + 1}
	}

	return indPairArr
}

func (p *Population) crossOver(indOne, indTwo int) {
	// Calculate the random cutoff point
	cutoffPoint := rand.Intn(p.nGene)

	sampleOneCopy := p.samples[indOne].CreateCopy()
	sampleTwoCopy := p.samples[indTwo].CreateCopy()

	for i := 0; i <= cutoffPoint; i++ {
		sampleOneGenes := sampleOneCopy.Genes()
		sampleTwoGenes := sampleTwoCopy.Genes()

		(*sampleTwoGenes)[i], (*sampleOneGenes)[i] = (*sampleOneGenes)[i], (*sampleTwoGenes)[i]
	}

	if sampleOneCopy.Evaluate() < sampleTwoCopy.Evaluate() {
		p.samples = append(p.samples, sampleTwoCopy)
	} else {
		p.samples = append(p.samples, sampleOneCopy)
	}
}

func (p *Population) mutation() {
	for i := range p.samples {
		p.samples[i].Mutation()
	}
}

func (p *Population) clean() {
	sort.Sort(ByScore(p.samples))
	p.samples = p.samples[:p.nSamples]
}

type ByScore []Chromosome

func (c ByScore) Len() int           { return len(c) }
func (c ByScore) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c ByScore) Less(i, j int) bool { return c[i].Evaluate() > c[j].Evaluate() }
