package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	jobs "nightcitylabbackend/Jobs"
	"os"
)

func main() {

	http.HandleFunc("/", GetMainPage)
	http.HandleFunc("/Update", SendSchedule)

	MyErrors := http.ListenAndServe(":8080", nil)

	if errors.Is(MyErrors, http.ErrServerClosed) {
		fmt.Printf("HomeLab is close")
	} else if MyErrors != nil {
		fmt.Printf("error starting server: %s\n", MyErrors)
		os.Exit(1)
	}

	jobs.RunCronJobs()
}

func GetMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request Successful / \n")
	io.WriteString(w, "NightCityLab is working \n")
}

func SendSchedule(w http.ResponseWriter, r *http.Request) {

	//CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	json.NewEncoder(w).Encode(jobs.Loadshedding{Schedule: jobs.GetGwede})

	fmt.Printf("Request Successful / \n")
}
