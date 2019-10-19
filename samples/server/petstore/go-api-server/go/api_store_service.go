/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
    "errors"
)

// StoreApiService is a service that implents the logic for the StoreApiServicer
// This service should implement the business logic for every endpoint for the StoreApi API. 
// Include any external packages or services that will be required by this service.
type StoreApiService struct {
}

// NewStoreApiService creates a default api service
func NewStoreApiService() StoreApiServicer {
    return &StoreApiService{}
}

// DeleteOrder - Delete purchase order by ID
func (s *StoreApiService) DeleteOrder(orderId string) (interface{}, error) {
    // TODO - update DeleteOrder with the required logic for this service method.
    // Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
    return nil, errors.New("service method 'DeleteOrder' not implemented")
}

// GetInventory - Returns pet inventories by status
func (s *StoreApiService) GetInventory() (interface{}, error) {
    // TODO - update GetInventory with the required logic for this service method.
    // Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
    return nil, errors.New("service method 'GetInventory' not implemented")
}

// GetOrderById - Find purchase order by ID
func (s *StoreApiService) GetOrderById(orderId int64) (interface{}, error) {
    // TODO - update GetOrderById with the required logic for this service method.
    // Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
    return nil, errors.New("service method 'GetOrderById' not implemented")
}

// PlaceOrder - Place an order for a pet
func (s *StoreApiService) PlaceOrder(body Order) (interface{}, error) {
    // TODO - update PlaceOrder with the required logic for this service method.
    // Add api_store_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
    return nil, errors.New("service method 'PlaceOrder' not implemented")
}