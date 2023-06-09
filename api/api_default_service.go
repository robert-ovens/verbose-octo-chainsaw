/*
 * compute
 *
 * Manage instances in cloud.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package compute

import (
	"context"
	"errors"
	"net/http"

	"github.com/robert-ovens/verbose-octo-chainsaw/backend"
)

// DefaultApiService is a service that implements the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
	backend backend.Backend
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService(backend backend.Backend) DefaultApiServicer {
	return &DefaultApiService{backend: backend}
}

// Create - create instance
func (s *DefaultApiService) Create(ctx context.Context, instance Instance) (ImplResponse, error) {
	// TODO - update Create with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, string{}) or use other options such as http.Ok ...
	//return Response(200, string{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("Create method not implemented")
}

func (s *DefaultApiService) Status(ctx context.Context) (ImplResponse, error) {

	return Response(http.StatusOK, nil), nil
}

// List - list instances
func (s *DefaultApiService) List(ctx context.Context) (ImplResponse, error) {

	instances, error := s.backend.GetInstances()
	if error != nil {
		return Response(http.StatusInternalServerError, nil), error
	}
	response := GetInstanceResponse{}
	response.Instances = make([]CreatedInstance, len(instances))
	for i := 0; i < len(instances); i++ {
		response.Instances[i] = CreatedInstance{
			Id:       instances[i].Id,
			Label:    instances[i].Label,
			Image:    instances[i].Image,
			SwapSize: instances[i].SwapSize,
			Type:     instances[i].Type,
		}
	}

	return Response(http.StatusOK, response), nil
}
