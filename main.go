package main

import (
	"fmt"
	"os"

	"github.com/juliencherry/whats-up/reminder"
)

var reminderManager *reminder.Manager

func main() {
	reminderManager = &reminder.Manager{
		Reminders: reminder.FileMap{},
	}

	args := os.Args[1:]

	if len(args) >= 3 {

		reminder := reminder.Reminder{
			Text: args[1],
		}
		date := args[2]

		reminderManager.Add(date, reminder)

		green := "\033[0;32m"
		noColor := "\033[0m"
		fmt.Printf("%sAdded reminder for %s:%s %s\n", green, date, noColor, reminder.Text)
		return
	}

	datesWithReminders := reminderManager.GetReminders()

	if len(datesWithReminders) == 0 {
		fmt.Println("No reminders!")
		return
	}

	fmt.Print("Reminders:\n\n")
	for date, reminders := range datesWithReminders {
		fmt.Println(date)
		for _, reminder := range reminders {
			fmt.Printf("• %s\n", reminder.Text)
		}
		fmt.Print("\n\n")
	}
}
