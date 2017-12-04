package reminder

type Reminder struct {
	Text string
	Date string
}

type Set interface {
	Add(element interface{})
	GetElements() []interface{}
}

type Manager struct {
	Reminders Set
}

func (m *Manager) Add(reminder Reminder) {
	m.Reminders.Add(reminder)
	return
}

func (m Manager) GetReminders() []Reminder {
	var reminders []Reminder

	for _, element := range m.Reminders.GetElements() {
		reminder, ok := element.(Reminder)
		if !ok {
			panic("could not get reminder")
		}

		reminders = append(reminders, reminder)
	}

	return reminders
}
