package jobs

import (
	"fmt"
	"net/http"
	"os/exec"
)

func jobs() {

}

// EskomSePush
func LoadSheddingSchedule(Area string, License string) string {
	cmd, MyErrors := exec.Command("./Scripts/GetLoadSheddingSchedule.sh", Area, License).Output()

	fmt.Printf("Request is running... / \n")

	if MyErrors != nil {
		fmt.Printf("error %s", MyErrors)
	}

	output := string(cmd)
	return output
}

func GetLoadSheddingUpdates(w http.ResponseWriter, r *http.Request) {
	cmd, MyResponses := exec.Command("cat CurrentLoadSheddingSchedule.txt").Output()

	fmt.Printf("Request is running... / \n")

	if MyResponses != nil {
		fmt.Printf("error %s", MyResponses)
	}

	EskomUpdates := string(cmd)
	fmt.Printf(EskomUpdates)
}
