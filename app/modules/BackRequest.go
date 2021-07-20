package modules

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"

	"ecom_frontend/app/dto"
)

//Configuration is a class for get all configuration from config folder
type BackRequest struct{}


//Modules to handle GET request and Auth Request
func (c BackRequest) BackendRequestGet(action string, username string, pass string) string {
	config := new(Configuration).GetConfig()
	url := config.GetString("BACKEND_URL") + action
	reqType := "POST"
	if username == "" {
		reqType = "GET"
	}

	dataPayload := dto.LoginPayload {
		Username: username,
		Password: pass,
	}
	fmt.Println(dataPayload)
	payloadIncentive, err := json.Marshal(dataPayload)
	body := bytes.NewReader(payloadIncentive)
	req, err := http.NewRequest(reqType, url, body)
	if err != nil {
	  fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+pass)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("request ", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Println("Error http req", err)
		return "failed"
	}
	defer resp.Body.Close()
	fmt.Println("resp body", resp.Body)

	dataCurl, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println("saat read", erro)
	}

	// fmt.Println(string(dataCurl))
	if string(dataCurl) != "" {
		return string(dataCurl)
	}else {
		return "failed"
	}

}

// modules to handle POST request
func (c BackRequest) BackendRequestPost(action string, token string, data dto.UpdateCart) string {
	config := new(Configuration).GetConfig()
	url := config.GetString("BACKEND_URL") + action
	reqType := "POST"

	fmt.Println(data)
	payloadIncentive, err := json.Marshal(data)
	body := bytes.NewReader(payloadIncentive)
	req, err := http.NewRequest(reqType, url, body)
	if err != nil {
	  fmt.Println(err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("request ", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		fmt.Println("Error http req", err)
		return "failed"
	}
	defer resp.Body.Close()
	fmt.Println("resp body", resp.Body)

	dataCurl, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		fmt.Println("saat read", erro)
	}

	// fmt.Println(string(dataCurl))
	if string(dataCurl) != "" {
		return string(dataCurl)
	}else {
		return "failed"
	}

}
