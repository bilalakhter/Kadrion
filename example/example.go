package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	server := &http.Server{
		Addr: ":3000",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{"message": "Hello from server"}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
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
