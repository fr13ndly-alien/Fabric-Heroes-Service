package web

import (
	"fmt"
	"net/http"

	"github.com/chainHero/heroes-service/web/controllers"
)

func Serve(app *controllers.Application) {
	//[Huu Hien] set web assetds: css, js, font, img
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//[Huu Hien] Handle URL, test send id via url in following line
	//[Huu Hien] heroes-services/web/controllers is imported, so can call func in there
	http.HandleFunc("/home.html/", app.HomeHandler) //call HomeHandler function to resolve url /home.html
	http.HandleFunc("/request.html", app.RequestHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("[web/app.go] Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)
}
