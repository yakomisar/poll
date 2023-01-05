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

func makeVote(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body and unmarshal it into a struct with the poll ID and choice ID.
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var vote struct {
		PollID int `json:"poll_id"`
		Choice int `json:"choice_id"`
	}
	err = json.Unmarshal(b, &vote)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Check if the poll ID is valid.
	if vote.PollID < 0 || vote.PollID >= len(polls) {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}

	// Check if the choice ID is valid.
	if vote.Choice < 0 || vote.Choice >= len(polls[vote.PollID].Choices) {
		http.Error(w, "Invalid choice ID", http.StatusBadRequest)
		return
	}

	// Increment the result count for the chosen poll and choice.
	polls[vote.PollID].Results[vote.Choice]++

	// Send a success response.
	w.WriteHeader(http.StatusOK)
}

func getResult(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read the request body and get the poll ID.
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var pollID struct {
		PollID int `json:"poll_id"`
	}
	err = json.Unmarshal(b, &pollID)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	// Check if the poll ID is valid.
	if pollID.PollID < 0 || pollID.PollID >= len(polls) {
		http.Error(w, "Invalid poll ID", http.StatusBadRequest)
		return
	}

	// Marshal the poll and write it to the response.
	b, err = json.Marshal(polls[pollID.PollID])
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
