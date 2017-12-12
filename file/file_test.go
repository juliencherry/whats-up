package file_test

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/juliencherry/whats-up/file"
	"github.com/juliencherry/whats-up/reminder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Set", func() {
	set := file.Set{}
	statePath := ".reminders"

	AfterEach(func() {
		os.Remove(statePath)
	})

	Context("no elements have been added", func() {
		It("gets no elements", func() {
			elements := set.GetElements()
			Expect(elements).To(BeEmpty())
		})

		It("does not create a state file", func() {
			_, err := os.Stat(statePath)
			Expect(os.IsNotExist(err)).To(BeTrue())
		})
	})

	Context("some elements have been added", func() {
		var elements = []reminder.Reminder{
			{
				Text: "Put on my teeth",
				Date: "1988-06-05",
			},
			{
				Text: "Brush my pants",
				Date: "2004-12-01",
			},
		}

		BeforeEach(func() {
			for _, element := range elements {
				set.Add(element)
			}
		})

		It("gets those elements", func() {
			gottenElements := set.GetElements()
			Expect(gottenElements).To(ConsistOf(elements))
		})

		It("creates a state file whose contents contain that element", func() {
			out, err := ioutil.ReadFile(statePath)
			Expect(err).NotTo(HaveOccurred())

			for _, element := range elements {
				expectedContents, err := json.Marshal(element)
				Expect(err).NotTo(HaveOccurred())

				Expect(string(out)).To(ContainSubstring(string(expectedContents)))
			}
		})
	})
})
