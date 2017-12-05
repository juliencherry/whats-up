package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Set struct{}

var statePath = ".set"

func (s *Set) Add(element interface{}) {
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

func (s Set) GetElements() []interface{} {
	content, err := ioutil.ReadFile(statePath)
	if os.IsNotExist(err) {
		return []interface{}{}
	} else if err != nil {
		panic(err)
	}

	var elements []interface{}
	err = json.Unmarshal(content, &elements)
	if err != nil {
		panic(err)
	}

	return elements
}
