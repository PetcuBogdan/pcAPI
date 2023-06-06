package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCards(w http.ResponseWriter, r *http.Request) {
	cards := getListOfCards()
	out, err := json.Marshal(cards)

	if err != nil {
        fmt.Println("error encoding people:", err)
    }

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func AddCard(w http.ResponseWriter, r *http.Request) {
	var card CardJson
	var payload jsonResponse
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&card)
	fmt.Println(card)
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
	

	saveCard(card)

	payload.Error = false
	payload.Message = "card saved"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
	
	saveCard(card)

}

func EditCard(w http.ResponseWriter, r *http.Request){
	var card CardJson
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&card)
	
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
	
	editCard(card)

	payload.Error = false
	payload.Message = "card edited"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func DeleteCard(w http.ResponseWriter, r *http.Request) {
	var cardId CardID
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&cardId)
	
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
	
	deleteCard(cardId.Id)

	payload.Error = false
	payload.Message = "card deleted"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}