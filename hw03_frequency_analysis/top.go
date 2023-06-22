package hw03frequencyanalysis

import (
	"strings"
)

// filter input text for additional task.
func filterAsterisk(input string) string {
	// regular expression if prefer for replace symbols.
	input = strings.ReplaceAll(input,".","")
	input = strings.ReplaceAll(input,"\"","")
	input = strings.ReplaceAll(input,",","")
	input = strings.ReplaceAll(input,"?","")
	input = strings.ReplaceAll(input,"!","")
	input = strings.ReplaceAll(input," - "," ")

	// to lower all text now
	input = strings.ToLower(input)

	return input
}

// word struct
type Word struct{
	Text string
	Counter int
}

// search word index in slice of Word, return index
func searchWord(text string, words []Word) int {
	for i, woritem :=range(words){
		if woritem.Text==text {
			return i
		}
	}
	return -1
}

// search word index with max counter in slice of Word, return index
func searchMax(words []Word) int {
	index:=-1
	for i, wordItem := range(words){
		if wordItem.Counter>0 && index==-1 {
			// if word has reseted counter do not use it
			index=i
		}
		if index>=0 && wordItem.Counter > words[index].Counter {
			index=i
		}
		if index>=0 && wordItem.Counter == words[index].Counter && strings.Compare(wordItem.Text,words[index].Text) <0 {
			index=i
		}
	}
	return index
}

func Top10(text string) []string {
	// just comment string below for standart task
	text = filterAsterisk(text)

	words := []Word{}
	// fill slice of Word
	for _, wordText := range (strings.Fields(text)) {
		if wordIndex:=searchWord(wordText,words); wordIndex>=0 {
			words[wordIndex].Counter ++
		}else{
			words=append(words, Word{wordText,1})
		}
	}

	result := []string {}

	for  i := 0; i < 10; i++{
		index:=searchMax(words)
		// search index and reset its counter
		if index>=0 {
			words[index].Counter=0
			result=append(result,words[index].Text)
		}
	}
	return result
}
