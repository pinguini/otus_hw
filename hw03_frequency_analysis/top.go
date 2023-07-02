package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// filter input text for additional task.
func filterAsterisk(input string) string {
	// remove punctuation and do not forget about " -"
	re := regexp.MustCompile(`[.,!?\'\"]|(\s+\-)|(\-\s+)`)
	input = strings.ToLower(re.ReplaceAllString(input, " "))
	return input
}

func Top10(text string) []string {
	// just comment string below for standart task
	text = filterAsterisk(text)

	words := map[string]int{}
	keys := []string{}

	// fill slice of Word
	for _, wordText := range strings.Fields(text) {
		_, ok := words[wordText]
		if ok {
			words[wordText]++
		} else {
			words[wordText] = 1
			keys = append(keys, wordText)
		}
	}

	// sort by count and alphabetical
	sort.SliceStable(keys, func(i, j int) bool {
		wordI, wordJ := keys[i], keys[j]
		return words[wordI] > words[wordJ] ||
			(words[wordI] == words[wordJ] && strings.Compare(wordJ, wordI) > 0)
	})

	// if slice more than 10 - cut it
	if len(keys) > 10 {
		keys = keys[:10]
	}
	return keys
}
