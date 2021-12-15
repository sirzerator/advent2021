package day14

import (
	"strings"
)

type PairStore struct {
	lastLetter string
	rules      map[string]string
	pairsCount map[string]int
}

func NewPairStore(template string, rules []string) *PairStore {
	ps := new(PairStore)
	ps.rules = make(map[string]string)
	ps.pairsCount = make(map[string]int)

	for _, rule := range rules {
		parts := strings.Split(rule, " -> ")
		left, right := parts[0], parts[1]
		ps.rules[left] = right
	}

	ps.lastLetter = string(template[len(template)-1])

	tLetters := strings.Split(template, "")
	for i := 0; i < len(template)-1; i++ {
		ps.pairsCount[tLetters[i]+tLetters[i+1]] += 1
	}

	return ps
}

func (ps *PairStore) ApplyRules() {
	newPairsCount := make(map[string]int)
	for pair, count := range ps.pairsCount {
		newPairsCount[pair] = count
	}

	for pair, count := range ps.pairsCount {
		pairLetters := strings.Split(pair, "")

		newPairsCount[pairLetters[0]+ps.rules[pair]] += count
		newPairsCount[ps.rules[pair]+pairLetters[1]] += count
		newPairsCount[pair] -= count
	}

	ps.pairsCount = newPairsCount
}

func (ps *PairStore) CountLetters() map[string]int {
	lettersFrequencyMap := make(map[string]int)

	for pair, count := range ps.pairsCount {
		pairLetters := strings.Split(pair, "")

		lettersFrequencyMap[pairLetters[0]] += count
	}

	return lettersFrequencyMap
}
