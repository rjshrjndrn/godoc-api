package main

import (
	"godoc/pkg/router"
	"log"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

func main() {
	// Getting env variables
	routerImp, err := router.NewRouter()
	if err != nil {
		log.Fatal("error creating router", err)
	}
	defer routerImp.DB.Close()
	httpRouter := http.NewServeMux()
	api := humago.New(httpRouter, huma.DefaultConfig("GoDoc", "1.0.0"))

	huma.Register(api, huma.Operation{
		OperationID: "Adding patient",
		Method:      http.MethodPost,
		Path:        "/patient",
		Summary:     "Onboard a patient.",
		Description: "Onboard a patient.",
		Tags:        []string{"patient"},
	}, routerImp.AddPatient)

	log.Println("Starting api at :8080")
	http.ListenAndServe(":8080", httpRouter)
}
