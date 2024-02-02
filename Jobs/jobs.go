package jobs

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
)

var s gocron.Scheduler
var PhoneMbalula string

func EskomSePush(NewArea string, NewToken string) {
	if NewArea != "" && NewToken != "" {
		PhoneMbalula := LoadSheddingSchedule(NewArea, NewToken)
		fmt.Print(PhoneMbalula)
	} else if NewArea == "" && NewToken == "" {
		fmt.Printf("Loadshedding schedule update failed.")
	}
}

func JobUpdates() {
	fmt.Printf("Jobs are running, please wait...  \n")
}

func RunCronJobs() {
	s := gocron.NewScheduler(time.Local)

	s.Every(1).Hour().Do(func() {
		EskomSePush("tshwane-13-thereeds", "ADA190D5-7B76409E-BF774100-B7D024CD")

	})

	s.StartAsync()

	JobUpdates()

}

func jobs() {
	RunCronJobs()
}

type Loadshedding struct {
	Schedule string `json:"schedule"`
}

// EskomSePush
func LoadSheddingSchedule(Area string, License string) string {
	Cmd, MyErrors := exec.Command("./Scripts/GetLoadSheddingSchedule.sh", Area, License).Output()

	fmt.Printf("EskomSePush is running... / \n")

	if MyErrors != nil {
		fmt.Printf("error %s", MyErrors)
	}

	Output := string(Cmd)
	return Output
}

func GetLoadSheddingUpdates(Code string) []string {

	CurrentReading, MyErrors := exec.Command("./Scripts/ReadCurrentSchedule.sh").Output()
	if MyErrors != nil {
		fmt.Printf("error %s", MyErrors)
	}

	EskomUpdates := strings.Split(string(CurrentReading[:]), "Script has been executed successfully.")
	return EskomUpdates
}
