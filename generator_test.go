package randstr

import "testing"

func TestGenerators(t *testing.T) {
	alphaSet := map[string]bool{}
	numSet := map[string]bool{}
	alnumSet := map[string]bool{}

	for _, ch := range alphabetSet {
		alphaSet[ch] = true
		alnumSet[ch] = true
	}
	for _, ch := range numberSet {
		numSet[ch] = true
		alnumSet[ch] = true
	}

	maxLen := 10

	t.Run("test fixed size string generating", func(t *testing.T) {
		alnumStr := NewAlnumStr(true, maxLen)
		for i := 0; i < 100; i++ {
			str := alnumStr.Gen()
			if len(str) != maxLen {
				t.Errorf("'%s' length is not %d", str, maxLen)
			}
			for _, ch := range str {
				if !alnumSet[string(ch)] {
					t.Errorf("'%s' is not alnum", str)
				}
			}
		}
	})

	t.Run("test unfixed size string generating", func(t *testing.T) {
		alnumStr := NewAlnumStr(false, maxLen)
		for i := 0; i < 100; i++ {
			str := alnumStr.Gen()
			if len(str) > maxLen {
				t.Errorf("'%s' length is longer than max length %d", str, maxLen)
			}
			for _, ch := range str {
				if !alnumSet[string(ch)] {
					t.Errorf("'%s' is not alnum", str)
				}
			}
		}
	})
}
