package randstr

import (
	"math/rand"
	"time"
)

var (
	alphabetSet = []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}
	numberSet = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
)

// Generator generates random strings
type RandStr struct {
	strSet     []string
	isLenFixed bool
	maxLen     int
}

func NewRandStr(strSet []string, isLenFixed bool, maxLen int) *RandStr {
	return &RandStr{
		strSet:     strSet,
		isLenFixed: isLenFixed,
		maxLen:     maxLen,
	}
}

func (g *RandStr) Gen() string {
	rand.Seed(time.Now().UnixNano())

	randStr := ""
	for i := 0; i < g.maxLen; i++ {
		randStr += g.strSet[rand.Intn(len(g.strSet))]
		if !g.isLenFixed && rand.Intn(2) == 1 {
			break
		}
	}

	return randStr
}

// AlphabetStr is random alphabetic string generator
type AlphabetStr struct {
	*RandStr
}

func NewAlphabetStr(isLenFixed bool, maxLen int) *AlphabetStr {
	return &AlphabetStr{
		RandStr: NewRandStr(alphabetSet, isLenFixed, maxLen),
	}
}

// NumberStr is random numeric string generator
type NumberStr struct {
	*RandStr
}

func NewNumberStr(isLenFixed bool, maxLen int) *NumberStr {
	return &NumberStr{
		RandStr: NewRandStr(numberSet, isLenFixed, maxLen),
	}
}

// AlnumStr is random alphanumeric string generator
type AlnumStr struct {
	*RandStr
}

func NewAlnumStr(isLenFixed bool, maxLen int) *AlnumStr {
	return &AlnumStr{
		RandStr: NewRandStr(append(numberSet, alphabetSet...), isLenFixed, maxLen),
	}
}
