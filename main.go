package main

import (
	"fmt"
	"os"
	"time"

	"github.com/juliencherry/whats-up/reminder"
)

var reminderManager *reminder.Manager

func main() {
	reminderManager = &reminder.Manager{
		Reminders: reminder.FileMap{},
	}

	args := os.Args[1:]

	if len(args) >= 3 {
		addReminder(args[2], args[1])
		return
	}

	datesWithReminders := reminderManager.GetReminders()

	if len(datesWithReminders) == 0 {
		fmt.Println("No reminders!")
		return
	}

	if len(args) == 1 {
		fmt.Println("Here’s what’s on your plate for today:")
		printReminders(datesWithReminders[time.Now().Format("2006-01-02")])
		return
	}

	fmt.Print("Reminders:\n\n")
	for date, reminders := range datesWithReminders {
		fmt.Println(date)
		printReminders(reminders)
		fmt.Print("\n\n")
	}
}

func addReminder(date string, text string) {
	reminder := reminder.Reminder{
		Text: text,
	}

	reminderManager.Add(date, reminder)

	green := "\033[0;32m"
	noColor := "\033[0m"
	fmt.Printf("%sAdded reminder for %s:%s %s\n", green, date, noColor, reminder.Text)
}

func printReminders(reminders []reminder.Reminder) {
	for _, reminder := range reminders {
		fmt.Printf("• %s\n", reminder.Text)
	}
}
