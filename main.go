package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Poll is a struct that represents a poll with multiple choices.
type Poll struct {
	ID       int      `json:"id"`
	Question string   `json:"question"`
	Choices  []string `json:"choices"`
	Results  []int    `json:"results"`
}

var polls []Poll

func main() {
	http.HandleFunc("/api/createPoll/", createPoll)
	http.HandleFunc("/api/poll/", makeVote)
	http.HandleFunc("/api/getResult/", getResult)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func createPoll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body and unmarshal it into a Poll struct.
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var p Poll
	err = json.Unmarshal(b, &p)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Assign the poll an ID and add it to the list of polls.
	p.ID = len(polls) + 1
	p.Results = make([]int, len(p.Choices))
	polls = append(polls, p)

	// Marshal the poll and write it to the response.
	b, err = json.Marshal(p)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
