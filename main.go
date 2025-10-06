// Package main indicates this is an executable program
// Every Go program starts with a main package and main() function
package main

// Import the net/http package for HTTP server functionality
// Single import for one package (could also use parentheses for multiple imports)
import "net/http"

// main() is the entry point of our program - like index.js in Node.js
// It takes no parameters and returns nothing
func main() {
	// Create an instance of our api struct using struct literal syntax
	// &api{} creates a pointer to a new api struct
	// We use a pointer because we might want to modify the struct later
	api := &api{addr: ":8080"} // addr: ":8080" means listen on port 8080

	// http.NewServeMux() creates a new HTTP request multiplexer (router)
	// It's like Express.js router - decides which handler function to call for each URL
	mux := http.NewServeMux()

	// Create an HTTP server configuration
	// &http.Server{} creates a pointer to a new Server struct
	// We configure it with our address and router
	srv := &http.Server{
		Addr:    api.addr, // Server address (":8080" means localhost:8080)
		Handler: mux,      // Router that will handle incoming requests
	}

	// Register route handlers - similar to app.get() and app.post() in Express.js
	// "GET /users" means this handler only responds to GET requests to /users
	// api.getUsersHandler is a method of our api struct
	mux.HandleFunc("GET /users", api.getUsersHandler)

	// "POST /users" means this handler only responds to POST requests to /users
	mux.HandleFunc("POST /users", api.createUserHandler)

	// Start the HTTP server and listen for incoming requests
	// ListenAndServe() blocks the program and keeps the server running
	// It returns an error if the server fails to start
	err := srv.ListenAndServe()
	if err != nil {
		// panic() is like throwing an exception - it stops the program immediately
		// In production code, you'd want more graceful error handling
		panic(err)
	}
}
