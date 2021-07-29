package sim

import (
	"github.com/magiconair/properties/assert"
	"math/rand"
	"testing"
)

func TestSim_Evaluate(t *testing.T) {
	type fields struct {
		data         *Data
		nbGeneration uint
		targetValue  float64
		maxShares    uint
		vec          []int
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "50/50",
			fields: fields{
				data: &Data{
					Targets: []Target{
						{Name: "domain1",
							Allocation: map[Part]Percentage{
								"part1": 50,
								"part2": 50,
							},
						},
					},
					Shares: []Share{
						{
							Code:      "share1",
							Owned:     0,
							Price:     10,
							CurrRatio: 1,
							Allocations: map[Domain]Allocation{
								"domain1": {
									"part1": 50,
									"part2": 50,
								},
							},
						},
						{
							Code:      "share2",
							Owned:     0,
							Price:     5,
							CurrRatio: 1,
							Allocations: map[Domain]Allocation{
								"domain1": {
									"part1": 50,
									"part2": 50,
								},
							},
						},
					},
				},
				nbGeneration: 0,
				targetValue:  10000,
				maxShares:    0,
				vec:          []int{500, 1000},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "excess total value, diff^2^2",
			fields: fields{
				data: &Data{
					Targets: []Target{
						{Name: "domain1",
							Allocation: map[Part]Percentage{
								"part1": 50,
								"part2": 50,
							},
						},
					},
					Shares: []Share{
						{
							Code:      "share1",
							Owned:     0,
							Price:     10,
							CurrRatio: 1,
							Allocations: map[Domain]Allocation{
								"domain1": {
									"part1": 50,
									"part2": 50,
								},
							},
						},
						{
							Code:      "share2",
							Owned:     0,
							Price:     5,
							CurrRatio: 1,
							Allocations: map[Domain]Allocation{
								"domain1": {
									"part1": 50,
									"part2": 50,
								},
							},
						},
					},
				},
				nbGeneration: 0,
				targetValue:  10000,
				maxShares:    0,
				vec:          []int{502, 1004},
			},
			want:    40 * 40 * 40 * 40, // 40^2 diff with target. >target so ^2 again
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sim{
				data:         tt.fields.data,
				nbGeneration: tt.fields.nbGeneration,
				targetValue:  tt.fields.targetValue,
				maxShares:    tt.fields.maxShares,
				vec:          tt.fields.vec,
			}
			got, err := s.Evaluate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Evaluate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Evaluate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimMutateAndControlsIfLessThanOwned(t *testing.T) {
	sim := Sim{
		data: &Data{
			Targets: []Target{
				{Name: "domain1",
					Allocation: map[Part]Percentage{
						"part1": 50,
						"part2": 50,
					},
				},
			},
			Shares: []Share{
				{
					Code:      "share1",
					Owned:     2,
					Price:     10,
					CurrRatio: 1,
					Allocations: map[Domain]Allocation{
						"domain1": {
							"part1": 50,
							"part2": 50,
						},
					},
				},
				{
					Code:      "share2",
					Owned:     2,
					Price:     5,
					CurrRatio: 1,
					Allocations: map[Domain]Allocation{
						"domain1": {
							"part1": 50,
							"part2": 50,
						},
					},
				},
			},
		},
		nbGeneration: 333,
		targetValue:  765,
		maxShares:    12,
		vec:          []int{2, 2},
	}

	// seed of 2 will produce less than owned shares
	rng := rand.New(rand.NewSource(2))
	sim.Mutate(rng)

	assert.Equal(t, 2, sim.vec[0])
	assert.Equal(t, 2, sim.vec[1])

}

func TestSimCloneClonesAreEquals(t *testing.T) {
	sim := Sim{
		data: &Data{
			Targets: []Target{
				{Name: "domain1",
					Allocation: map[Part]Percentage{
						"part1": 50,
						"part2": 50,
					},
				},
			},
			Shares: []Share{
				{
					Code:      "share1",
					Owned:     0,
					Price:     10,
					CurrRatio: 1,
					Allocations: map[Domain]Allocation{
						"domain1": {
							"part1": 50,
							"part2": 50,
						},
					},
				},
				{
					Code:      "share2",
					Owned:     0,
					Price:     5,
					CurrRatio: 1,
					Allocations: map[Domain]Allocation{
						"domain1": {
							"part1": 50,
							"part2": 50,
						},
					},
				},
			},
		},
		nbGeneration: 333,
		targetValue:  765,
		maxShares:    12,
		vec:          []int{502, 1004},
	}

	clone := sim.Clone()

	assert.Equal(t, clone, sim)
}
