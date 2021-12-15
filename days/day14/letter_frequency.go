package day14

import (
	"sort"
	"strings"
)

type LetterFrequency struct {
	letter string
	count  int
}

type LetterFrequencyList []LetterFrequency

func buildLettersFrequencyMap(str string) map[string]int {
	lettersFrequencyMap := make(map[string]int)

	for _, l := range strings.Split(str, "") {
		lettersFrequencyMap[l] += 1
	}

	return lettersFrequencyMap
}

func sortByLetterFrequency(lettersFrequencyMap map[string]int) LetterFrequencyList {
	lfl := make(LetterFrequencyList, len(lettersFrequencyMap))
	i := 0

	for letter, count := range lettersFrequencyMap {
		lfl[i] = LetterFrequency{letter, count}
		i++
	}

	sort.Sort(sort.Reverse(lfl))

	return lfl
}

func (lfl LetterFrequencyList) Len() int {
	return len(lfl)
}

func (lfl LetterFrequencyList) Less(i, j int) bool {
	return lfl[i].count < lfl[j].count
}

func (lfl LetterFrequencyList) Swap(i, j int) {
	lfl[i], lfl[j] = lfl[j], lfl[i]
}
