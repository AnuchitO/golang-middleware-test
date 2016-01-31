package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type IPMiddleware struct {
}

func (ip *IPMiddleware) MiddlewareFunc(handler rest.HandlerFunc) rest.HandlerFunc {
	fmt.Println(":: IPMiddleware ::")
	return func(w rest.ResponseWriter, r *rest.Request) {
		fmt.Println(":: Handler ::")

		handler(w, r)
	}
}

func main() {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&IPMiddleware{})
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
