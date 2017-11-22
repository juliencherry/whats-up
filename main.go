package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const reminderPath = ".reminders"

func main() {
	args := os.Args[1:]

	if len(args) >= 2 {
		reminder := args[1]
		if err := addReminder(reminder); err != nil {
			log.Fatal(err)
			return
		}

		green := "\033[0;32m"
		noColor := "\033[0m"
		fmt.Printf("%sAdded reminder:%s %s\n", green, noColor, reminder)
		return
	}

	reminder, err := getReminder()
	if err != nil {
		log.Fatal(err)
		return
	}

	if reminder == "" {
		fmt.Println("No reminders!")
		return
	}

	fmt.Println("Reminders:")
	fmt.Println("•", reminder)
}

func addReminder(reminder string) error {
	return ioutil.WriteFile(reminderPath, []byte(reminder), 0666)
}

func getReminder() (string, error) {
	data, err := ioutil.ReadFile(reminderPath)
	if os.IsNotExist(err) {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", data), nil
}
