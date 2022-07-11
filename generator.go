package randstr

import (
	"math/rand"
)

var (
	alphabetSet = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	numberSet         = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	defaultMaxStrLen  = 8
	defaultIsLenFixed = false
)

// RandStr is random string generator
type RandStr struct {
	candidates      []string
	sampleSizeFixed bool
	sampleSizeMax   int
}

func New(candidates []string) *RandStr {
	return NewRandStr(candidates, defaultIsLenFixed, defaultMaxStrLen)
}

// NewRandStr create a new generator
func NewRandStr(candidates []string, sampleSizeFixed bool, sampleSizeMax int) *RandStr {
	return &RandStr{
		candidates:      candidates,
		sampleSizeFixed: sampleSizeFixed,
		sampleSizeMax:   sampleSizeMax,
	}
}

func (g *RandStr) sample(candidates []string) string {
	sampleSize := g.sampleSizeMax
	if !g.sampleSizeFixed {
		sampleSize = rand.Intn(g.sampleSizeMax) + 1
	}

	randStr := ""
	for i := 0; i < sampleSize; i++ {
		randStr += candidates[rand.Intn(len(candidates))]
	}

	return randStr
}

// SetLenFixed sets if randstr generates fixed size random strings
func (g *RandStr) SetLenFixed(isFixed bool) {
	g.sampleSizeFixed = isFixed
}

// SetLenMax sets the max length of random strings.
// If sample size is fixed, it is the sample size.
func (g *RandStr) SetLenMax(sampleSizeMax int) {
	g.sampleSizeMax = sampleSizeMax
}

// Seed sets a seed to math.rand
func (g *RandStr) Seed(seed int64) {
	rand.Seed(seed)
}

// Gen generates a random string according to candidates
func (g *RandStr) Gen() string {
	return g.sample(g.candidates)
}

// GenSlice generates a slice of random strings according to candidates
func (g *RandStr) GenSlice(length int) []string {
	return g.genSlice(length, g.candidates)
}

func (g *RandStr) genSlice(length int, candidates []string) []string {
	strs := make([]string, length, length)
	for i := range strs {
		strs[i] = g.sample(candidates)
	}
	return strs
}

// Alphabets generates a random string with alphabets
func (g *RandStr) Alphabets() string {
	return g.sample(alphabetSet)
}

// AlphabetSlice generates a slice of random strings with alphabets
func (g *RandStr) AlphabetSlice(length int) []string {
	return g.genSlice(length, alphabetSet)
}

// Numbers generates a random string with numbers
func (g *RandStr) Numbers() string {
	return g.sample(numberSet)
}

// NumberSlice generates a slice of random strings with numbers
func (g *RandStr) NumberSlice(length int) []string {
	return g.genSlice(length, alphabetSet)
}

// Alnums generates a random string with alphabets and numbers
func (g *RandStr) Alnums() string {
	return g.sample(append(alphabetSet, numberSet...))
}

// AlnumSlice generates a slice of random strings with alphabets and numbers
func (g *RandStr) AlnumSlice(length int) []string {
	return g.genSlice(length, append(alphabetSet, numberSet...))
}
