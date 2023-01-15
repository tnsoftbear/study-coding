package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func countMatchChars(checkWordRev, dictionaryWordRev string) int {
	var count = 0
	var dicB = []byte(dictionaryWordRev)
	for i, c := range []byte(checkWordRev) {
		if len(dicB) < i+1 {
			return count
		}
		if c == dicB[i] {
			count++
		} else {
			return count
		}
	}
	return count
}

func findInDictionary(checkWord string, dictionaryRev map[byte][]string) string {
	var checkWordRev = reverseWord(checkWord)
	var maxSize = 0
	var maxDicWordRev string
	var matchCount = 0
	var key = checkWordRev[0]
	for i := len(dictionaryRev[key]) - 1; i >= 0; i-- {
		// for _, dicWordRev := range dictionaryRev[key] {
		var dicWordRev = dictionaryRev[key][i]
		if checkWordRev == dicWordRev {
			continue // don't find equal words
		}
		matchCount = countMatchChars(checkWordRev, dicWordRev)
		if maxSize < matchCount {
			maxSize = matchCount
			maxDicWordRev = dicWordRev
		}
		if maxSize == len(checkWord) {
			break
		}
		if maxSize > matchCount {
			break // reach the max word before
		}
	}
	return reverseWord(maxDicWordRev)
}

func reverseWord(word string) string {
	reversed := []rune(word)
	for i, j := 0, len(reversed)-1; i < len(reversed)/2; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	return string(reversed)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var dictionaryCount int
	fmt.Fscanf(in, "%d\n", &dictionaryCount)
	//var dictionaryMapRev = make(map[string]string, dictionaryCount)

	var firstWord string
	var inputWord string

	var dictionaryMap = make(map[string]string, dictionaryCount)
	for inputWordIdx := 0; inputWordIdx < dictionaryCount; inputWordIdx++ {
		fmt.Fscanf(in, "%s\n", &inputWord)
		if len(firstWord) == 0 {
			firstWord = inputWord
		}
		dictionaryMap[inputWord] = inputWord
	}

	var dictionaryRev = make(map[byte][]string, 26)
	var i = 0
	var key byte
	for _, v := range dictionaryMap {
		var vRev = reverseWord(v)
		key = vRev[0]
		dictionaryRev[key] = append(dictionaryRev[key], vRev)
		i++
	}

	for k, _ := range dictionaryRev {
		sort.Strings(dictionaryRev[k])
	}

	var checkWordCount int
	fmt.Fscanf(in, "%d\n", &checkWordCount)

	var checkWord string
	var rhymeWords = make(map[string]string, checkWordCount)
	for checkWordIdx := 0; checkWordIdx < checkWordCount; checkWordIdx++ {
		fmt.Fscanf(in, "%s\n", &checkWord)
		if rhymeWord, ok := rhymeWords[checkWord]; ok {
			fmt.Fprintln(out, rhymeWord)
			continue
		}
		var foundRhyme = findInDictionary(checkWord, dictionaryRev)
		rhymeWords[checkWord] = foundRhyme
		if len(rhymeWords[checkWord]) == 0 {
			rhymeWords[checkWord] = firstWord
		}
		fmt.Fprintln(out, rhymeWords[checkWord])
	}

	fmt.Fprintln(out)
}
