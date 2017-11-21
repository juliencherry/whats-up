package main_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

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

		JustBeforeEach(func() {
			cmd := exec.Command(bin, args...)
			out, err := cmd.Output()
			Expect(err).NotTo(HaveOccurred())

			output = string(out)
		})

		It("prints a welcome message", func() {
			Expect(output).To(HavePrefix("What’s up?"))
		})

		Context("when there are at least two arguments", func() {
			var reminder string

			BeforeEach(func() {
				reminder = "Do something important"
				args = []string{"add", reminder}
			})

			It("prints a message confirming that the reminder was added", func() {
				green := "\033[0;32m"
				noColor := "\033[0m"
				message := fmt.Sprintf("%sAdded reminder: %s%s", green, noColor, reminder)
				Expect(output).To(HavePrefix(message))
			})
		})
	})
})
