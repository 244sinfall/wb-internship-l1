package main

func task20(words string) string {
	var managedWords []string
	inRunes := []rune(words)
	var currentWord []rune
	for idx, thisRune := range inRunes {
		if thisRune != ' ' {
			currentWord = append(currentWord, thisRune)
		}
		if (thisRune == ' ' || idx == len(inRunes)-1) && len(currentWord) != 0 {
			managedWords = append(managedWords, string(currentWord))
			currentWord = nil
		}
	}
	var output string
	for i := len(managedWords) - 1; i >= 0; i-- {
		output += managedWords[i]
		if i != 0 {
			output += " "
		}
	}
	return output
}
