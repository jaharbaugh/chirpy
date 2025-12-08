package main

import(
	"strings"
)

func censor(unclean string) string{

	bannedWords :=[]string{"kerfuffle", "sharbert", "fornax"}

	replacementString := "****"

	words := strings.Split(unclean, " ")

	for i, word := range words{
		lower := strings.ToLower(word)
		for _, banned := range bannedWords {
			if lower == banned{
				words[i] = replacementString
			}
		}
	}

	cleaned := strings.Join(words, " ")
	return cleaned
}