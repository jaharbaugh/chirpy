package auth

import(
	"net/http"
	"strings"
	"errors"
)

func GetAPIKey (headers http.Header) (string, error) {
	headerString := headers.Get("Authorization")
	if headerString == ""{
		return "", errors.New("Header not set")
	}

	pref := "ApiKey "
	keyString := ""

	if strings.HasPrefix(headerString, pref){
		keyString = strings.TrimPrefix(headerString, pref)
		keyString = strings.TrimSpace(keyString)
			if keyString == ""{
				return "", errors.New("Missing ApiKey string")
			}
		}else{
			return "", errors.New("Incorrect prefix")
		}
		
	return keyString, nil
}