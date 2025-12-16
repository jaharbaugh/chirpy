package main

import(
	"errors"
)

func Validate (body string)(string, error){ 
	const maxChars = 140
	
	if len(body) > maxChars{
		 
		return "", errors.New("Chirp too long")
		

	}else{
		cleaned := censor(body)
		return cleaned, nil
	} 

}