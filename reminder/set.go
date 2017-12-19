package reminder

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type FileSet struct{}

var statePath = ".reminders"

func (f FileSet) Add(element Reminder) {
	elements := append(f.GetElements(), element)

	elementsAsJSON, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(statePath, elementsAsJSON, 0777)
	if err != nil {
		panic(err)
	}
}

func (f FileSet) GetElements() []Reminder {
	content, err := ioutil.ReadFile(statePath)
	if os.IsNotExist(err) {
		return []Reminder{}
	} else if err != nil {
		panic(err)
	}

	var elements []Reminder
	err = json.Unmarshal(content, &elements)
	if err != nil {
		panic(err)
	}

	return elements
}
