package file

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/juliencherry/whats-up/reminder"
)

type Set struct{}

var statePath = ".reminders"

func (s Set) Add(element reminder.Reminder) {
	elements := append(s.GetElements(), element)

	elementsAsJSON, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(statePath, elementsAsJSON, 0777)
	if err != nil {
		panic(err)
	}
}

func (s Set) GetElements() []reminder.Reminder {
	content, err := ioutil.ReadFile(statePath)
	if os.IsNotExist(err) {
		return []reminder.Reminder{}
	} else if err != nil {
		panic(err)
	}

	var elements []reminder.Reminder
	err = json.Unmarshal(content, &elements)
	if err != nil {
		panic(err)
	}

	return elements
}
