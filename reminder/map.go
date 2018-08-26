package reminder

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type FileMap struct{}

var statePath = ".reminders"

func (f FileMap) Add(key string, value Reminder) {
	elements := f.GetElements()

	elements[key] = append(elements[key], value)

	elementsAsJSON, err := json.Marshal(elements)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(statePath, elementsAsJSON, 0777)
	if err != nil {
		panic(err)
	}
}

func (f FileMap) Get(key string) []Reminder {
	return nil
}

func (f FileMap) GetElements() map[string][]Reminder {
	content, err := ioutil.ReadFile(statePath)
	if os.IsNotExist(err) {
		return map[string][]Reminder{}
	} else if err != nil {
		panic(err)
	}

	var elements map[string][]Reminder
	err = json.Unmarshal(content, &elements)
	if err != nil {
		panic(err)
	}

	return elements
}
