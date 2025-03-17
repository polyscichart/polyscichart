package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TweetRequest mimics the X API v2 payload for creating a tweet
type TweetRequest struct {
	Text string `json:"text"`
}

// TweetResponse mimics the X API v2 response structure
type TweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

func main() {
	// Set up the HTTP server and route
	http.HandleFunc("/2/tweets", createTweetHandler)

	// Start the server on port 8080
	fmt.Println("Mock X API server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}

// createTweetHandler handles POST requests to /2/tweets
func createTweetHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests (X API uses POST for creating tweets)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming JSON payload
	var req TweetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Basic validation (e.g., mimic X's 280-character limit)
	if len(req.Text) == 0 || len(req.Text) > 280 {
		http.Error(w, "Text must be between 1 and 280 characters", http.StatusBadRequest)
		return
	}

	// Create a mock response
	resp := TweetResponse{}
	resp.Data.ID = "mock_tweet_123456789" // Static ID for simplicity
	resp.Data.Text = req.Text

	// Set response headers (mimicking X API)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 status, like X API for successful post

	// Encode and send the response
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
