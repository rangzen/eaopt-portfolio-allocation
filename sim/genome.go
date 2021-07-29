package sim

import (
	"github.com/MaxHalford/eaopt"
	"math"
	"math/rand"
)

func (s Sim) Evaluate() (float64, error) {
	var totalValue float64
	var nbVariables int
	for _, s := range s.data.Shares {
		nbVariables = max(nbVariables, s)
	}
	var allocValues = make(map[Domain]map[Part]float64, nbVariables)
	for i, sh := range s.data.Shares {
		value := float64(s.vec[i]) * sh.Price * sh.CurrRatio
		totalValue += value
		for domain, allocation := range sh.Allocations {
			m := allocValues[domain]
			if m == nil {
				allocValues[domain] = make(map[Part]float64, 0)
			}
			for part, percentage := range allocation {
				allocValues[domain][part] += value * percentage / 100
			}
		}
	}

	diff := math.Pow(s.targetValue-totalValue, 2)
	for _, target := range s.data.Targets {
		for part, percentage := range target.Allocation {
			d := percentage/100*totalValue - allocValues[target.Name][part]
			diff += math.Pow(d, 2)
		}
	}
	if totalValue > s.targetValue {
		return math.Pow(diff, 2), nil
	}
	return diff, nil
}

func max(nbVariables int, s Share) int {
	return int(math.Max(float64(nbVariables), float64(len(s.Allocations))))
}

func (s Sim) Mutate(rng *rand.Rand) {
	for i := range s.vec {
		s.vec[i] += rng.Intn(501) - 250
		if s.vec[i] < s.data.Shares[i].Owned {
			s.vec[i] = s.data.Shares[i].Owned
		}
	}
}

func (s Sim) Crossover(_ eaopt.Genome, _ *rand.Rand) {}

func (s Sim) Clone() eaopt.Genome {
	newSim := Sim{
		data:         s.data,
		nbGeneration: s.nbGeneration,
		targetValue:  s.targetValue,
		maxShares:    s.maxShares,
		vec:          make([]int, len(s.vec)),
	}
	copy(newSim.vec, s.vec)
	return newSim
}
