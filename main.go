package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	user string = "victoryotaghogho"
)

type Post struct {
	Operation_type string `json:"operation_type"`
	X              int    `json:"x"`
	Y              int    `json:"y"`
}

type report struct {
	SlackUsername  string      `json:"slackUsername"`
	Result         int         `json:"result"`
	Operation_type interface{} `json:"operation_type"`
}

func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := Post{}
	res := report{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		fmt.Fprintln(w, "Error decoding json request")
	}
	res.SlackUsername = user
	res.Operation_type = req.Operation_type

	if strings.Contains(req.Operation_type, "add") {
		res.Result = req.X + req.Y
		json.NewEncoder(w).Encode(res)
		return
	} else if strings.Contains(req.Operation_type, "sub") {
		res.Result = req.X - req.Y
		json.NewEncoder(w).Encode(res)
		return
	} else if strings.Contains(req.Operation_type, "mul") {
		res.Result = req.X * req.Y
		json.NewEncoder(w).Encode(res)
		return
	}

}

func main() {
	http.HandleFunc("/", post)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
