package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"order/models"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /user", createHandler)
	mux.HandleFunc("PUT /user", updateHandler)
	mux.HandleFunc("DELETE /user", deleteHandler)

	log.Println("Starting server on :9000")
	err := http.ListenAndServe(":9000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var userReq models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	jsonData, err := json.Marshal(userReq)
	if err != nil {
		log.Println("Error marshaling UserRequest to JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	targetURL := url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort("localhost", "9001"),
		Path:   "user",
	}

	req, err := http.NewRequestWithContext(r.Context(), "POST", targetURL.String(), bytes.NewReader(jsonData))
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending HTTP request:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Unexpected response status code: %s", resp.Status)
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "The result is %s\n", body)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var userReq models.UserRequest
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	jsonData, err := json.Marshal(userReq)
	if err != nil {
		log.Println("Error marshaling UserRequest to JSON:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	targetURL := url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort("localhost", "9001"),
		Path:   "user/" + id,
	}

	req, err := http.NewRequestWithContext(r.Context(), "POST", targetURL.String(), bytes.NewReader(jsonData))
	if err != nil {
		log.Println("Error creating HTTP request:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending HTTP request:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("Unexpected response status code: %s", resp.Status)
		http.Error(w, resp.Status, resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "The result is %s\n", body)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {

}
