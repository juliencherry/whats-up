package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Set struct{}

var statePath = ".set"

func (s *Set) Add(element interface{}) {
	f, err := os.OpenFile(statePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	elementAsJSON, err := json.Marshal(element)
	if err != nil {
		panic(err)
	}

	if _, err = f.Write(elementAsJSON); err != nil {
		panic(err)
	}
}

func (s Set) GetElements() []string {
	content, err := ioutil.ReadFile(statePath)
	if os.IsNotExist(err) {
		return []string{}
	} else if err != nil {
		panic(err)
	}

	elements := strings.TrimSpace(string(content))
	return strings.Split(elements, "\n")
}
