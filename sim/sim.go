package sim

import (
	"encoding/json"
	"fmt"
	"github.com/MaxHalford/eaopt"
	"io/ioutil"
	"math"
	"math/rand"
)

func (s Sim) String() string {
	return fmt.Sprint(s.vec)
}

func New(filename string, nbGeneration uint, targetValue float64, maxShares uint) (Sim, error) {
	data, err := loadData(filename)
	if err != nil {
		return Sim{}, err
	}
	return Sim{data: &data, nbGeneration: nbGeneration, targetValue: targetValue, maxShares: maxShares}, nil
}

func loadData(filename string) (data Data, err error) {
	var rawData []byte
	if rawData, err = ioutil.ReadFile(filename); err != nil {
		return Data{}, err
	}
	if err = json.Unmarshal(rawData, &data); err != nil {
		return Data{}, err
	}
	return
}

func (s *Sim) Run() {
	ga, err := eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	ga.HofSize = 3
	ga.NGenerations = s.nbGeneration

	generationBetweenPrints := s.nbGeneration / 100
	ga.Callback = func(ga *eaopt.GA) {
		if ga.Generations%generationBetweenPrints == 0 {
			fmt.Printf("Best fitness at generation %d(%s): %s\n",
				ga.Generations,
				ga.Age,
				ga.HallOfFame[0],
			)
		}
	}

	err = ga.Minimize(
		func(rng *rand.Rand) eaopt.Genome {
			s.vec = initInt(len(s.data.Shares), s.maxShares, rng)
			return s
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	printResult(ga.HallOfFame[0].Genome.(Sim))
}

// initInt generates random ints x such that x < upper.
func initInt(n int, upper uint, rng *rand.Rand) (ints []int) {
	ints = make([]int, n)
	for i := range ints {
		ints[i] = int(math.Floor(rng.Float64() * (float64(upper))))
	}
	return
}

func printResult(s Sim) {
	fmt.Println("Result:")
	const sizeCode = "6"
	const sizeNb = "5"
	const sizeValue = "10"
	fmt.Printf("%"+sizeCode+"s %"+sizeNb+"s %"+sizeValue+"s %"+sizeValue+"s\n", "Code", "Nb", "Value", "Wo. ratio")
	for i := 0; i < len(s.vec); i++ {
		fmt.Printf("%"+sizeCode+"s %"+sizeNb+"d %"+sizeValue+".2f",
			s.data.Shares[i].Code,
			s.vec[i],
			float64(s.vec[i])*s.data.Shares[i].Price*s.data.Shares[i].CurrRatio,
		)
		if s.data.Shares[i].CurrRatio != 1 {
			fmt.Printf(" %"+sizeValue+".2f", float64(s.vec[i])*s.data.Shares[i].Price)
		}
		fmt.Println("")
	}
}
