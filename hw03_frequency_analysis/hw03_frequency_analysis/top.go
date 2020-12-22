package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
)

var topWordsCount = 10

func Top10(rawText string) []string {
	if len(rawText) == 0 {
		return nil
	}

	// split text by space symbol
	r := regexp.MustCompile(`\n|\s|\t|\z`)
	textSlice := r.Split(rawText, -1)
	entryes := map[string]int{}
	for _, word := range textSlice {
		if word == "" {
			continue
		}
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
	return sortedWords[:topWordsCount]
}
