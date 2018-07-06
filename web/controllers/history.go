package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *Application) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	blockData, err := app.Fabric.QueryHello()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	type Test struct {
		Value   string `json:"value"`
		
	}

	type Data struct {
		Key    string `json:"key"`
		Record Test    `json:"record"`
	}

	type RecordHistory struct {
		TxId      string `json:"TxId"`
		Value     Test    `json:"Value"`
		Timestamp string `json:"Timestamp"`
		IsDelete  string `json:"IsDelete"`
	}

	var data []Data
	json.Unmarshal([]byte(blockData), &data)

	returnData := &struct {
		ResponseData         []Data
		TransactionRequested string
		TransactionUpdated   string
		RecordHistory        []RecordHistory
	}{
		ResponseData:         data,
		TransactionRequested: "true",
	}
	// Query History Using Key
	
	if r.FormValue("requested") == "true" {
		// Retrieving Single Query
		QueryValue := r.FormValue("KeySearch")
		blockHistory, _ := app.Fabric.GetHistory(QueryValue)
		var queryResponse []RecordHistory
		json.Unmarshal([]byte(blockHistory), &queryResponse)
		returnData.RecordHistory = queryResponse
		returnData.TransactionRequested = "true"
		fmt.Println("### Response History ###")
		//fmt.Printf("%s", blockHistory)
		fmt.Println(blockHistory)
	}
	
	renderTemplate(w, r, "history.html", returnData)
}
