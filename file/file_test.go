package file_test

import (
	"io/ioutil"
	"os"

	"github.com/juliencherry/whats-up/file"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("File", func() {
	set := file.Set{}
	statePath := ".set"

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

	Context("one element has been added", func() {
		element := "hydrogen"

		BeforeEach(func() {
			set.Add(element)
		})

		It("gets that element", func() {
			elements := set.GetElements()
			Expect(elements).To(ConsistOf(element))
		})

		It("creates a state file whose contents contain that element", func() {
			out, err := ioutil.ReadFile(statePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(out).To(HavePrefix(element))
		})
	})
})
