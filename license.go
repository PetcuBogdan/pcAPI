package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AddLicense(w http.ResponseWriter, r *http.Request) {
	var license LicenseVerification
	var payload jsonResponse
	fmt.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&license)
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

	payload.Error = false
	payload.Message = saveLicense(license.License_key, license.User_id)
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)

}

func DeleteLicense(w http.ResponseWriter, r *http.Request) {
	var license LicenseKey
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&license)
	
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
	
	deleteLicense(license.License_key)

	payload.Error = false
	payload.Message = "license deleted"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}