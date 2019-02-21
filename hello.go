package main

import (
	"bufio"
	"log"
	"net/http"
	"os"
)

func main() {

	PORT := ":8080"
	log.Print("Running server on " + PORT)
	http.HandleFunc("/", CompleteTaskFunc)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
	page := ReadFile("Testpage.html")
	w.Write([]byte(page))
}

func ReadFile(path string) string {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	input := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
