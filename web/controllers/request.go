package controllers

import (
	"net/http"
)
type Data struct {
	Key   string `json:"key"`
	Value  string `json:"value"`
}

func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		TransactionId string
		Success       bool
		Response      bool
	}{
		TransactionId: "",
		Success:       false,
		Response:      false,
	}
	if r.FormValue("submitted") == "true" {
		helloValue := r.FormValue("hello")
		Val := r.FormValue("value")
		txid, err := app.Fabric.InvokeHello(helloValue,Val)
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "request.html", data)
}
