package main

import (
	"log"
	"net"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", AuthenticationAPI),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8088", api.MakeHandler()))
}

func AuthenticationAPI(w rest.ResponseWriter, req *rest.Request) {
	ip, _ := net.LookupIP(req.PathParam("host"))
	// rest.Error(w, err.Error(), http.StatusInternalServerError)
	w.WriteJson(&ip)
	w.WriteJson(map[string]interface{}{"Body": ip})
}
