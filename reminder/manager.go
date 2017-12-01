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
	var reminder []Reminder

	for _, element := range m.Reminders.GetElements() {
		reminder = append(reminder, Reminder(element))
	}

	return reminders
}
