package handlers

import (
	"io"
	"net/http"
)

// LeetCodeProxyHandler forwards requests to LeetCode GraphQL API
func LeetCodeProxyHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create a new request to the LeetCode API
	proxyReq, err := http.NewRequest(
		http.MethodPost,
		"https://leetcode.com/graphql/",
		r.Body,
	)
	if err != nil {
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	// Copy headers from original request
	proxyReq.Header = r.Header

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Error making proxy request: "+err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Set content type header
	w.Header().Set("Content-Type", "application/json")

	// Copy response status code
	w.WriteHeader(resp.StatusCode)

	// Copy response body
	io.Copy(w, resp.Body)
}

