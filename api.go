// Package declaration - must match across all files in the same package
package main

// Import statements - bringing in external packages we need
import (
	"encoding/json" // For JSON encoding/decoding (marshal/unmarshal)
	"errors"        // For creating custom error messages
	"net/http"      // For HTTP server functionality
)

// api struct holds configuration for our API server
// In Go, we use structs instead of classes for data organization
type api struct {
	addr string // Server address (e.g., ":8080")
}

// Package-level variable to store our users in memory
// []User means "slice of User" - like an array but more flexible
// var declares a variable at package level (accessible throughout the package)
var users = []User{}

// Method definition: (receiver) functionName(parameters) returnType
// (a *api) is the receiver - this function "belongs to" the api struct
// *api means "pointer to api" - allows us to modify the original struct
func (a *api) getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Set HTTP response header to indicate we're returning JSON
	// w.Header() returns a map-like structure for HTTP headers
	w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w) creates a JSON encoder that writes directly to the response
	// .Encode(users) converts our users slice to JSON and writes it to the response
	// This is more efficient than json.Marshal() for HTTP responses
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		// http.Error sends an HTTP error response with the specified message and status code
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return // Exit the function early if there's an error
	}

	// Set HTTP status code to 200 OK
	// Note: when using json.NewEncoder(w).Encode(), this should come after encoding
	w.WriteHeader(http.StatusOK)
}

// Handler for creating new users via POST requests
func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// var payload User declares a variable of type User with zero values
	// Zero values: int=0, string="", bool=false, pointers=nil
	var payload User

	// json.NewDecoder(r.Body) creates a decoder that reads from the request body
	// .Decode(&payload) parses JSON and fills our payload struct
	// &payload gives the memory address of payload (required for modification)
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		// Return 400 Bad Request if JSON is malformed
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new User struct using struct literal syntax
	// User{field: value, field: value} creates and initializes a struct
	u := User{
		ID:    len(users) + 1, // Simple ID generation
		Name:  payload.Name,   // Copy name from the request
		Email: payload.Email,  // Copy email from the request
	}

	// Call our validation function
	// Functions can return multiple values - here we only care about the error
	err = insertUser(u)
	if err != nil {
		// Return 400 Bad Request if validation fails
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set HTTP status code to 201 Created (indicates successful creation)
	w.WriteHeader(http.StatusCreated)
}

// insertUser validates and adds a user to our in-memory storage
// Returns an error if validation fails, nil if successful
// This demonstrates Go's error handling pattern: return error as last value
func insertUser(u User) error {
	// Validation: check required fields
	// In Go, empty string is "", not null/undefined like in JavaScript
	if u.Email == "" {
		// errors.New() creates a new error with the given message
		return errors.New("email is required")
	}
	if u.Name == "" {
		return errors.New("name is required")
	}

	// Check for duplicate emails
	// for range loops over slices, arrays, maps, channels, strings
	// _ discards the index, user gets each User in the slice
	for _, user := range users {
		if user.Email == u.Email {
			return errors.New("email already exists")
		}
	}

	// append() adds elements to a slice and returns a new slice
	// In Go, slices can grow dynamically (unlike arrays which have fixed size)
	users = append(users, u)

	// Return nil (no error) to indicate success
	// nil is Go's equivalent to null/undefined for pointers, slices, maps, channels, interfaces
	return nil
}
