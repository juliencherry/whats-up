package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"fmt"
	"os"
	"os/exec"
)

type Reminder struct {
	text string
	date string
}

var _ = Describe("Main", func() {
	var args []string

	Context("when the binary runs", func() {
		var (
			bin    string
			output string
		)

		BeforeEach(func() {
			var err error
			bin, err = Build("github.com/juliencherry/whats-up")
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			err := os.Remove(".reminders")
			if !os.IsNotExist(err) {
				Expect(err).ToNot(HaveOccurred())
			}
		})

		JustBeforeEach(func() {
			cmd := exec.Command(bin, args...)
			out, err := cmd.Output()
			Expect(err).NotTo(HaveOccurred())

			output = string(out)
		})

		Context("with no arguments", func() {
			BeforeEach(func() {
				args = []string{}
			})

			Context("when no reminders have been added", func() {
				It("displays a message indicating that there are no reminders", func() {
					Expect(output).To(HavePrefix("No reminders!"))
				})
			})

			Context("when some reminders have been added", func() {
				var reminders []Reminder

				BeforeEach(func() {
					reminders = []Reminder{
						{
							text: "Put on my teeth",
							date: "1988-06-05",
						},
						{
							text: "Brush my pants",
							date: "2004-12-01",
						},
					}

					for _, reminder := range reminders {
						cmd := exec.Command(bin, "add", reminder.text, reminder.date)
						_, err := cmd.Output()
						Expect(err).NotTo(HaveOccurred())
					}
				})

				It("displays those reminders by date", func() {
					Expect(output).To(HavePrefix("Reminders:\n\n"))

					for _, reminder := range reminders {
						Expect(output).To(ContainSubstring(fmt.Sprintf("%s\n• %s\n\n", reminder.date, reminder.text)))
					}
				})
			})
		})

		Context("with three arguments", func() {
			var reminder string
			var date string

			BeforeEach(func() {
				reminder = "Do something important"
				date = "2013-02-27"
				args = []string{"add", reminder, date}
			})

			It("displays a message confirming that the reminder was added", func() {
				green := "\033[0;32m"
				noColor := "\033[0m"
				message := fmt.Sprintf("%sAdded reminder for %s:%s %s", green, date, noColor, reminder)
				Expect(output).To(HavePrefix(message))
			})
		})
	})
})
