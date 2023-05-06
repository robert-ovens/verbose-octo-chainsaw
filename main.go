package main

import (
	"log"
	"net/http"

	compute "github.com/robert-ovens/verbose-octo-chainsaw/api"
	"github.com/robert-ovens/verbose-octo-chainsaw/backend"
)

func main() {
	log.Printf("Server started")

	BackendImpl := backend.NewBackendImpl()
	DefaultApiService := compute.NewDefaultApiService(BackendImpl)
	DefaultApiController := compute.NewDefaultApiController(DefaultApiService)
	router := compute.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))

}
