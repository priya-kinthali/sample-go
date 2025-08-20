package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Welcome to the Home Page!"))
}

func main() {
    // Load environment variables from a .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Get a value from the environment
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Create a new router
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)

    // Start the server
    fmt.Printf("Server is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}