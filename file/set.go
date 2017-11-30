package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Set struct{}

var statePath = ".set"

func (s *Set) Add(element string) {
	f, err := os.OpenFile(statePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintln(element)); err != nil {
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
