package notify

import (
	"fmt"
	"os/exec"
)

func SendNotification(message string) {
	// Execute the shell script to send a notification
	cmd := exec.Command("sh", "-c", fmt.Sprintf("osascript -e 'display notification \"%s\" with title \"Nike Alert\"'", message))
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error sending notification: %v\n", err)
		return
	}
}
