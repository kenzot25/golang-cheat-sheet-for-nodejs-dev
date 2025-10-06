// Package declaration - all Go files must start with this
// 'main' package indicates this is an executable program
package main

// User represents a user in our system
// This is a struct - Go's way of defining custom data types (like classes in other languages)
type User struct {
	// Field declarations: FieldName Type `json:"tag"`
	// The `json:"..."` tags tell Go how to marshal/unmarshal this struct to/from JSON
	// These tags are used by json.Marshal() and json.Unmarshal() functions
	ID    int    `json:"id"`    // Auto-generated unique identifier
	Name  string `json:"name"`  // User's display name
	Email string `json:"email"` // User's email address (must be unique)
}
