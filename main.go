package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	jobs "nightcitylabbackend/Jobs"
	"os"
	"time"

	//Custom Addons
	"github.com/robfig/cron"
)

var Area string = os.Args[1]
var License string = os.Args[2]

func main() {

	http.HandleFunc("/", GetMainPage)
	http.HandleFunc("/hello", GetHTTPTestRequest)
	http.HandleFunc("/eskom", Eskom)
	MyErrors := http.ListenAndServe(":3333", nil)

	if errors.Is(MyErrors, http.ErrServerClosed) {
		fmt.Printf("HomeLab is close")
	} else if MyErrors != nil {
		fmt.Printf("error starting server: %s\n", MyErrors)
		os.Exit(1)
	}
}

func GetMainPage(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Request Successful / \n")
	io.WriteString(w, "NightCityLab is working \n")

}

func GetHTTPTestRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Request Successful / \n")
	io.WriteString(w, "NightCityLab is working (with HTTP) \n")
}

// EskomSePush
func Eskom(w http.ResponseWriter, r *http.Request) {

	//CronJobs
	UpdateLoadSheddingSchedule := cron.New()

	UpdateLoadSheddingSchedule.AddFunc("0 0 * * 1-6", func() {

		if Area != "" && License != "" {
			PhoneMbalula := jobs.LoadSheddingSchedule(Area, License)
			fmt.Printf(PhoneMbalula)
		} else if Area == "" && License == "" {
			fmt.Printf("Loadshedding schedule update failed.")
			io.WriteString(w, "Loadshedding schedule update failed. \n")
		}

	})

	UpdateLoadSheddingSchedule.Start()

	time.Sleep(1 * time.Minute)

	UpdateLoadSheddingSchedule.Stop()

}
