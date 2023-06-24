package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// filter input text for additional task.
func filterAsterisk(input string) string {
	// remove punctuation and do not forget about " -"
	re := regexp.MustCompile(`[.,!?\'\"]|(\s+\-)`)
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
		return words[keys[i]] > words[keys[j]] || (words[keys[i]] == words[keys[j]] && strings.Compare(keys[j], keys[i]) > 0)
	})

	// if slice more than 10 - cut it
	if len(keys) > 10 {
		keys = keys[0:10]
	}
	return keys
}
