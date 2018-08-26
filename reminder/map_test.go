package reminder_test

import (
	"encoding/json"
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/juliencherry/whats-up/reminder"
)

var _ = Describe("Map", func() {
	fileMap := &reminder.FileMap{}
	statePath := ".reminders"

	AfterEach(func() {
		os.Remove(statePath)
	})

	Context("no elements have been added", func() {
		It("gets no elements", func() {
			elements := fileMap.GetElements()
			Expect(elements).To(BeEmpty())
		})

		It("does not create a state file", func() {
			_, err := os.Stat(statePath)
			Expect(os.IsNotExist(err)).To(BeTrue())
		})
	})

	Context("some elements have been added", func() {
		var elements = map[string][]reminder.Reminder{
			"1988-06-05": []reminder.Reminder{
				{Text: "Do something important"},
			},

			"2004-12-01": []reminder.Reminder{
				{Text: "Put on my teeth"},
				{Text: "Brush my pants"},
			},
		}

		BeforeEach(func() {
			for date, reminders := range elements {
				for _, reminder := range reminders {
					fileMap.Add(date, reminder)
				}
			}
		})

		It("gets those elements", func() {
			gottenElements := fileMap.GetElements()
			Expect(gottenElements).To(Equal(elements))
		})

		It("creates a state file whose contents contain that element", func() {
			out, err := ioutil.ReadFile(statePath)
			Expect(err).NotTo(HaveOccurred())

			for _, reminders := range elements {
				expectedContents, err := json.Marshal(reminders)
				Expect(err).NotTo(HaveOccurred())

				Expect(string(out)).To(ContainSubstring(string(expectedContents)))
			}
		})
	})
})
