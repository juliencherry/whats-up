package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"os/exec"
)

var _ = Describe("Main", func() {

	Context("when the binary is executed", func() {
		var str string

		BeforeEach(func() {
			bin, err := Build("github.com/juliencherry/whats-up")
			Expect(err).NotTo(HaveOccurred())

			cmd := exec.Command(bin)
			out, err := cmd.Output()
			Expect(err).NotTo(HaveOccurred())

			str = string(out)
		})

		It("prints “What’s up?”", func() {
			Expect(str).To(HavePrefix("What’s up?"))
		})
	})
})
