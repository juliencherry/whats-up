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
			Reminders: &reminder.FakeMap{Map: make(map[string][]reminder.Reminder)},
		}
	})

	Context("no reminders have been added", func() {
		It("retrieves no reminders", func() {
			reminders := manager.GetReminders()
			Expect(reminders).To(BeEmpty())
		})
	})

	Context("some reminders have been added", func() {
		var datesWithReminders = map[string][]reminder.Reminder{
			"1988-06-05": []reminder.Reminder{{Text: "Put on my teeth"}},
			"2004-12-01": []reminder.Reminder{{Text: "Brush my pants"}},
		}

		BeforeEach(func() {
			for date, reminders := range datesWithReminders {
				for _, reminder := range reminders {
					manager.Add(date, reminder)
				}
			}
		})

		It("retrieves those reminders", func() {
			retrievedReminders := manager.GetReminders()
			Expect(retrievedReminders).To(Equal(datesWithReminders))
		})
	})
})
