package main

func main() {
	c := NewChromosome(10000)
	c.GeneratePopulation()
	c.Evolve()

	c.Elitism().PrettyPrint()
}
