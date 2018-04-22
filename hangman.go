package main

import "strings"

type Word struct {
	word string
	lettersGuessed []string
	lettersLeft []string
	encryptedWord string
}

func makeWord(word string) *Word {
	makeEncrypt := func(_ rune) rune { return '?'}
	var letters = unique(strings.Split(strings.ToLower(word), ""))
	var encrypt = strings.Map(makeEncrypt, word)

	return &Word{word: word, lettersLeft: letters, encryptedWord: encrypt}
}

func guessLetter(myWord *Word, letterGuess string) bool {
	for index, letter := range myWord.lettersLeft {
		if letter == letterGuess {
			myWord.lettersGuessed = append(myWord.lettersGuessed, letterGuess)
			myWord.lettersLeft = remove(myWord.lettersLeft, index)
			revealLetter(myWord, letterGuess)
			return true
		}
	}
	return false
}

func revealLetter(myWord *Word, letter string) {
	var splitQuestions = strings.Split(myWord.encryptedWord, "")
	for i := 0; i < len(myWord.word); i ++ {
		if myWord.word[i] == letter[0] {
			splitQuestions[i] = letter
		}
	}
	myWord.encryptedWord = strings.Join(splitQuestions, "")
}

// Function is a modified version of http://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Function modified from https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-array-in-golang/37335777
func remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}