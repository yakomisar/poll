package main

import (
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
