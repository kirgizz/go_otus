package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"sort"
	"strings"
)

var topWordsCount = 10

func Top10(rawText string) []string {
	if len(rawText) == 0 {
		return nil
	}

	// split text by space symbol
	textSlice := strings.Fields(rawText)
	entryes := make(map[string]int)
	for _, word := range textSlice {
		entryes[word]++
	}

	// create list with words and sort it by entryes count
	sortedWords := make([]string, 0)
	for entry := range entryes {
		sortedWords = append(sortedWords, entry)
	}
	sort.Slice(sortedWords, func(i, j int) bool {
		return entryes[sortedWords[i]] > entryes[sortedWords[j]]
	})
	if len(sortedWords) <= topWordsCount {
		return sortedWords
	}
	return sortedWords[:topWordsCount]
}
