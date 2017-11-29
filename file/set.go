package file

import (
	"io/ioutil"
	"os"
)

type Set struct{}

var statePath = ".set"

func (s *Set) Add(element string) {
	ioutil.WriteFile(statePath, []byte(element), 0600)
}

func (s Set) GetElements() []string {
	data, err := ioutil.ReadFile(statePath)
	if os.IsNotExist(err) {
		return []string{}
	} else if err != nil {
		panic(err)
	}
	return []string{string(data)}
}
