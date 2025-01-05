package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Value string `json:"value"`
}

func Stage1(data *Data) error {
	if data.Value == "" {
		return fmt.Errorf("value is empty")
	}
	fmt.Println("Stage 1: Processing data")
	return nil
}

func Stage2(data *Data) error {
	if data.Value == "error" {
		return fmt.Errorf("error in stage 2")
	}
	fmt.Println("Stage 2: Processing data")
	return nil
}

func Stage3(data *Data) error {
	fmt.Println("Stage 3: Processing data")
	return nil
}

func Pipeline(data *Data) error {
	if err := Stage1(data); err != nil {
		return err
	}

	if err := Stage2(data); err != nil {
		return err
	}

	if err := Stage3(data); err != nil {
		return err
	}

	return nil
}

func ProcessData(w http.ResponseWriter, r *http.Request) {
	var data Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := Pipeline(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/process", ProcessData)

	log.Fatal(http.ListenAndServe(":8080", nil))
}