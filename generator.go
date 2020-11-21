package randstr

import (
	"math/rand"
)

var (
	alphabetSet = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	numberSet = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

// RandStr is random string generator
type RandStr struct {
	seed            int64
	candidates      []string
	sampleSizeFixed bool
	sampleSizeMax   int
}

// NewRandStr create a new generator
func NewRandStr(candidates []string, sampleSizeFixed bool, sampleSizeMax int) *RandStr {
	return &RandStr{
		candidates:      candidates,
		sampleSizeFixed: sampleSizeFixed,
		sampleSizeMax:   sampleSizeMax,
	}
}

func (g *RandStr) genBy(candidates []string) string {
	rand.Seed(g.seed)

	randStr := ""
	for i := 0; i < g.sampleSizeMax; i++ {
		randStr += candidates[rand.Intn(len(candidates))]
		if i != 0 && !g.sampleSizeFixed && rand.Intn(2) == 1 { // at least sample once
			break
		}
	}

	return randStr
}

// Seed sets a seed to math.rand
func (g *RandStr) Seed(seed int64) {
	g.seed = seed
}

// Gen generates a random string according to candidates
func (g *RandStr) Gen() string {
	return g.genBy(g.candidates)
}

// Alphabets generates a random string with alphabets
func (g *RandStr) Alphabets() string {
	return g.genBy(alphabetSet)
}

// Numbers generates a random string with numbers
func (g *RandStr) Numbers() string {
	return g.genBy(numberSet)
}

// Alnums generates a random string with alphabets and numbers
func (g *RandStr) Alnums() string {
	return g.genBy(append(alphabetSet, numberSet...))
}
