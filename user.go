package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user UserJson
	var payload jsonResponse
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if err != nil{
		fmt.Println("invalid json")
		payload.Error = true
		payload.Message = "invalid json"
		
		out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}
	

	saveUser(user)

	payload.Error = false
	payload.Message = "user saved"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}