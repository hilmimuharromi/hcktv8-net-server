package main

import (
	"fmt"
	"net/http"
	"net_server/service"
)

var UserSvc = service.NewUserService()

func user(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		UserSvc.GetUser(w)
	case "POST":
		UserSvc.Register(r, &w)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

const (
	port = "8090"
)

func main() {
	http.HandleFunc("/user", user)

	fmt.Println("=================Running at port " + port + "=================")
	http.ListenAndServe(":"+port, nil)
}
