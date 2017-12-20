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
				var reminderMap map[string][]Reminder

				BeforeEach(func() {
					reminderMap = map[string][]Reminder{
						"2007-02-19": []Reminder{
							{
								text: "Put on my teeth",
								date: "2007-02-19",
							},
							{
								text: "Brush my pants",
								date: "2007-02-19",
							},
						},
						"2004-12-01": []Reminder{{
							text: "Get ready",
							date: "2004-12-01",
						}},
					}

					for _, reminders := range reminderMap {
						for _, reminder := range reminders {
							cmd := exec.Command(bin, "add", reminder.text, reminder.date)
							_, err := cmd.Output()
							Expect(err).NotTo(HaveOccurred())
						}
					}
				})

				It("displays those reminders grouped by date", func() {
					Expect(output).To(HavePrefix("Reminders:\n\n"))

					for date, reminders := range reminderMap {
						substring := fmt.Sprintf("%s\n", date)

						for _, reminder := range reminders {
							substring += fmt.Sprintf("• %s\n", reminder.text)
						}

						substring += "\n"

						Expect(output).To(ContainSubstring(substring))
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
