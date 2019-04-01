package controllers

import (
	"net/http"
)

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	//call fabric function named QueryHello
	helloValue, err := app.Fabric.QueryHello()

	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	// Struct does not need coma(,)
	data := &struct {
		Hello string
		Greet string
	}{
		Hello: helloValue,
		Greet: "Huu Hien welcome you:",
	}
	renderTemplate(w, r, "home.html", data)
}
