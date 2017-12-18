package main

import (
	"fmt"
	"os"

	"github.com/juliencherry/whats-up/file"
	"github.com/juliencherry/whats-up/reminder"
)

var reminderManager *reminder.Manager

func main() {
	reminderManager = &reminder.Manager{
		Reminders: file.Set{},
	}

	args := os.Args[1:]

	if len(args) >= 3 {
		reminder := reminder.Reminder{
			Text: args[1],
			Date: args[2],
		}

		reminderManager.Add(reminder)

		green := "\033[0;32m"
		noColor := "\033[0m"
		fmt.Printf("%sAdded reminder for %s:%s %s\n", green, reminder.Date, noColor, reminder.Text)
		return
	}

	reminders := reminderManager.GetReminders()

	if len(reminders) == 0 {
		fmt.Println("No reminders!")
		return
	}

	fmt.Println("Reminders:")
	for _, reminder := range reminders {
		fmt.Println(reminder.Date)
		fmt.Printf("• %s\n", reminder.Text)
	}
}
