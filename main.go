package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Post("/", AuthenticationAPI),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	fmt.Println(" starting ")
	log.Fatal(http.ListenAndServe(":8088", api.MakeHandler()))
}

func AuthenticationAPI(w rest.ResponseWriter, req *rest.Request) {
	fmt.Println(":: AuthenticationAPI ::")
	ip, _ := net.LookupIP(req.PathParam("host"))
	// rest.Error(w, err.Error(), http.StatusInternalServerError)
	w.WriteJson(&ip)
	w.WriteJson(map[string]interface{}{"Body": ip})
}
