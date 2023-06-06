package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetShipments(w http.ResponseWriter, r *http.Request) {
	shipments := getListOfShipments()

	out, err := json.Marshal(shipments)

	if err != nil {
        fmt.Println("error encoding people:", err)
    }

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func AddShipment(w http.ResponseWriter, r *http.Request) {
	var shipment ShipmentJson
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&shipment)

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
	
	saveShipment(shipment)

	payload.Error = false
	payload.Message = "profile saved"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func EditShipment(w http.ResponseWriter, r *http.Request) {
	var shipment ShipmentJson
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&shipment)

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
	
	editShipment(shipment)

	payload.Error = false
	payload.Message = "profile edited"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func DeleteShipment(w http.ResponseWriter, r *http.Request) {
	var shipmentId ShipmentID
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&shipmentId)
	
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
	
	deleteShipment(shipmentId.Id)

	payload.Error = false
	payload.Message = "profile deleted"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}