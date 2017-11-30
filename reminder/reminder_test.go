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
		var reminders = []reminder.Reminder{
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
