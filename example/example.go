package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/mux"
)

type BodyJson struct {
	Health string `json:"Health"`
	Really struct {
		Check string `json:"Check"`
	} `json:"Really"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", CreateItem).Methods("POST")

	http.Handle("/", router)

	go func() {
		http.ListenAndServe(":3000", nil)
	}()

	exec.Command("go", "build", "-o", "bin/kadrion", "cmd/kadrion/main.go")
	terminal := exec.Command("./bin/kadrion", "apply", "tconfig.yaml")

	result, err := terminal.CombinedOutput()

	if err != nil {
		fmt.Println("Error executing command:", err)
	} else {
		fmt.Println()
		fmt.Println(" ------Test Completed successfully-----")
		fmt.Println()
		fmt.Println("Command output:")
		fmt.Println(string(result))
	}

	os.Exit(0)
}
func CreateItem(w http.ResponseWriter, r *http.Request) {
	var requestBody BodyJson
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := "Hello from server"
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(response))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
