package randstr

import (
	"testing"
)

func TestGenerators(t *testing.T) {
	// rely on dynamic programming to check if str can be composed of candidates
	validate := func(str string, candidates []string) {
		match := make([]bool, len(str))
		for j := range str {
			for _, candidate := range candidates {
				if j-len(candidate) >= 0 && match[j-len(candidate)] {
					if str[j+1-len(candidate):j+1] == candidate {
						match[j] = true
						break
					}
				} else if j == 0 {
					if string(str[0:1]) == candidate {
						match[j] = true
						break
					}
				} else if j-len(candidate) == -1 && str[0:j+1] == candidate {
					match[j] = true
					break
				}
			}
		}
		if !match[len(match)-1] {
			t.Fatalf("no match for random str %s candidates(%v) match(%v)", str, candidates, match)
		}
	}

	t.Run("test Gen()", func(t *testing.T) {
		candidates := []string{"a", "bc", "def"}
		randStr := NewRandStr(candidates, true, 4)
		for i := 0; i < 10; i++ {
			str := randStr.Gen()
			validate(str, candidates)
		}
	})

	t.Run("test Alnums()", func(t *testing.T) {
		randStr := NewRandStr([]string{}, true, 4)
		for i := 0; i < 10; i++ {
			str := randStr.Alnums()
			validate(str, append(numberSet, alphabetSet...))
		}
	})

	t.Run("test Alphabets()", func(t *testing.T) {
		randStr := NewRandStr([]string{}, true, 4)
		for i := 0; i < 10; i++ {
			str := randStr.Alphabets()
			validate(str, alphabetSet)
		}
	})

	t.Run("test Numbers()", func(t *testing.T) {
		randStr := NewRandStr([]string{}, true, 4)
		for i := 0; i < 10; i++ {
			str := randStr.Numbers()
			validate(str, numberSet)
		}
	})
}
