package main

import(
	"net/http"
)

func handlerValidate (w http.ResponseWriter, req *http.Request){ 
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	max_chars := 140
	
	chirpBody, err := decode[Chirp](req)
	if err != nil{
		respondWithError(w, http.StatusInternalServerError, "Error decoding parameters", err)
		return
	}

	var status int
	var res interface{}

	if len(chirpBody.Body) > max_chars{
		err := "Chirp too long" 
		res, status = responseInvalid(err)

	}else{
		validRes, s := responseValid()

		validRes.Cleaned = censor(chirpBody.Body)
		res = validRes
		status = s
	}

	respondJSON(w, status, res)
}