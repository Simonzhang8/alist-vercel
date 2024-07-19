package main

import (
	"net/http"
	"os"
	"os/exec"
	"io"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("./alist")
	cmd.Stdout = w
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		http.Error(w, "Error creating stdin pipe", http.StatusInternalServerError)
		return
	}

	go func() {
		defer stdin.Close()
		io.Copy(stdin, r.Body)
	}()

	if err := cmd.Run(); err != nil {
		http.Error(w, "Error running command", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5244", nil)
}
