package reminder_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/juliencherry/whats-up/reminder"
)

var _ = Describe("Manager", func() {
	var manager reminder.Manager

	BeforeEach(func() {
		manager = reminder.Manager{
			Reminders: &reminder.FakeSet{},
		}
	})

	Context("no reminders have been added", func() {
		It("retrieves no reminders", func() {
			reminders := manager.GetReminders()
			Expect(reminders).To(BeEmpty())
		})
	})

	Context("one reminder has been added", func() {
		reminder := "Put on my teeth and brush my pants"

		BeforeEach(func() {
			manager.Add(reminder)
		})

		It("retrieves that reminder", func() {
			reminders := manager.GetReminders()
			Expect(reminders).To(ConsistOf(reminder))
		})
	})
})
