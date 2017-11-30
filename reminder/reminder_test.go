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

	Context("some reminders have been added", func() {
		reminders := []string{"Put on my teeth", "Brush my pants"}

		BeforeEach(func() {
			for _, reminder := range reminders {
				manager.Add(reminder)
			}
		})

		It("retrieves those reminders", func() {
			retrievedReminders := manager.GetReminders()
			Expect(retrievedReminders).To(ConsistOf(reminders))
		})
	})
})
