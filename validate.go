package main

import(
	"net/http"
	"encoding/json"
	"log"
)

func handlerValidate (w http.ResponseWriter, req *http.Request){ 
	if req.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

	max_chars := 140
	
	chirpBody, err := decode(req)
	if err != nil{
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	var status int
	var res interface{}

	if len(chirpBody.Body) > max_chars{
		res, status = responseInvalid()

	}else{
		validRes, s := responseValid()

		validRes.Cleaned = censor(chirpBody.Body)
		res = validRes
		status = s
	}

	dat, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
	w.Write(dat)

}