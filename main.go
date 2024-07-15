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
	cfg := huma.DefaultConfig("GoDoc", "v1.0.0")
	// Disabling the url schema from output
	// https://github.com/danielgtaylor/huma/issues/230#issuecomment-1927998004
	cfg.CreateHooks = []func(huma.Config) huma.Config{}

	api := humago.New(httpRouter, cfg)

	huma.Register(api, huma.Operation{
		OperationID: "Adding patient",
		Method:      http.MethodPost,
		Path:        "/patient",
		Summary:     "Onboard a patient.",
		Description: "Onboard a patient.",
		Tags:        []string{"patient"},
	}, routerImp.AddPatient)

	huma.Register(api, huma.Operation{
		OperationID: "Search patient",
		Method:      http.MethodGet,
		Path:        "/patient/{firstName}",
		Summary:     "Search a patient.",
		Description: "Search a patient.",
		Tags:        []string{"patient"},
	}, routerImp.GetPatient)

	huma.Register(api, huma.Operation{
		OperationID: "List all users with pagination",
		Method:      http.MethodGet,
		Path:        "/patients/{page}",
		Summary:     "List all patients.",
		Description: "List all patients.",
		Tags:        []string{"patient"},
	}, routerImp.ListPatients)

	log.Println("Starting api at :8080")
	http.ListenAndServe(":8000", httpRouter)
}
