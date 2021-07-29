package sim

type Sim struct {
	data         *Data
	nbGeneration uint
	targetValue  float64
	maxShares    uint
	vec          []int
}

type Data struct {
	Targets []Target `json:"targets"`
	Shares  []Share  `json:"shares"`
}

type Target struct {
	Name       Domain     `json:"name"`
	Allocation Allocation `json:"allocation"`
}

type Domain = string
type Part = string
type Percentage = float64
type Allocation map[Part]Percentage

type Allocations map[Domain]Allocation

type Share struct {
	Code        string      `json:"code"`
	Owned       int         `json:"owned"`
	Price       float64     `json:"price"`
	CurrRatio   float64     `json:"curr_ratio"`
	Allocations Allocations `json:"allocations"`
}
