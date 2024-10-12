package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	jobs "nightcitylabbackend/Jobs"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var FinalSchedule []string
var DC []string

// This code was adapted from Dev.To
// Author: Enda
// Year: 2023
// Link: https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	} else {
		log.Print((".env loaded successfully"))

		//Find Location and Keys
		//os.LookupEnv("APP_ENV_Eskom_Location")
		//os.LookupEnv("APP_ENV_Eskom_Key")

		//Get Location and Keys
		Location := os.Getenv("APP_ENV_Eskom_Location")
		API := os.Getenv("APP_ENV_Eskom_Key")

		jobs.RunCronJobs(Location, API)
	}

}

func main() {

	http.HandleFunc("/", GetMainPage)
	http.HandleFunc("/Update", SendSchedule)

	MyErrors := http.ListenAndServe(":8080", nil)

	if errors.Is(MyErrors, http.ErrServerClosed) {
		fmt.Printf("HomeLab is closed")
	} else if MyErrors != nil {
		fmt.Printf("error starting server: %s\n", MyErrors)
		os.Exit(1)
	}

}

func GetMainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request Successful / \n")
	io.WriteString(w, "NightCityLab is working \n")
}

func SendSchedule(w http.ResponseWriter, r *http.Request) {

	FinalSchedule := jobs.GetLoadSheddingUpdates("Solidus")
	DF := FinalSchedule[0]
	DD := strings.Split(DF, "Current")
	//DC := strings.Join(DD, "")
	//ReadySchedule := strings.Join(strings.Split(DC, "\n"), "")
	//DDD := strconv.Itoa(int(len(ReadySchedule)))

	//CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")

	json.NewEncoder(w).Encode(jobs.Loadshedding{Schedule: strings.Join(DD, "")})

	fmt.Printf("Request Successful / \n")
}
