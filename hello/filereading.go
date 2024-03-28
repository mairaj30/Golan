package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/question", readQuestions)
	http.ListenAndServe(":8080", nil)
}

func readQuestions(w http.ResponseWriter, r *http.Request) {

	flag := r.URL.Query().Get("flag")
	fmt.Println(w, "flag: %s", flag)
	filePath := "tandc.txt"
	if flag == "question" {
		filePath = "question.txt"
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Read the contents of the file
	byteContent, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert the contents of the file to a string and print it
	fileContent := string(byteContent)
	//fmt.Println(fileContent)

	json.NewEncoder(w).Encode(fileContent)
}
