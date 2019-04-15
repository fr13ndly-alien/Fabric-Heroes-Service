package controllers

import (
	"fmt"
	"net/http"
)

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

	// Check form submitted
	if r.FormValue("submitted") == "true" {
		fmt.Printf("\t[HuuHien] Handling request.html\n")
		helloValue := r.FormValue("hello")
		pinValue := r.FormValue("pin")
		fmt.Println("\t- [HuuHien] pin value: ", pinValue)
		txid, err := app.Fabric.InvokeHello(helloValue)
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		data.TransactionId = txid
		data.Success = true
		data.Response = true
	}
	// Render to HTML and write to HTTP response
	renderTemplate(w, r, "request.html", data)
}
