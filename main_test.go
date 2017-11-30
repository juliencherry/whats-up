package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"fmt"
	"os"
	"os/exec"
)

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
			err := os.Remove(".set")
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
				var reminders []string

				BeforeEach(func() {
					reminders = []string{"Put on my teeth", "Brush my pants"}

					for _, reminder := range reminders {
						cmd := exec.Command(bin, "add", reminder)
						_, err := cmd.Output()
						Expect(err).NotTo(HaveOccurred())
					}
				})

				It("displays those reminders", func() {
					Expect(output).To(HavePrefix("Reminders:\n"))

					for _, reminder := range reminders {
						Expect(output).To(ContainSubstring(fmt.Sprintf("• %s\n", reminder)))
					}
				})
			})
		})

		Context("with two arguments", func() {
			var reminder string

			BeforeEach(func() {
				reminder = "Do something important"
				args = []string{"add", reminder}
			})

			It("displays a message confirming that the reminder was added", func() {
				green := "\033[0;32m"
				noColor := "\033[0m"
				message := fmt.Sprintf("%sAdded reminder:%s %s", green, noColor, reminder)
				Expect(output).To(HavePrefix(message))
			})
		})
	})
})
