package reminder

type Reminder struct {
	Text string
	Date string
}

type Set interface {
	Add(reminder Reminder)
	GetElements() []Reminder
}

type Manager struct {
	Reminders Set
}

func (m *Manager) Add(reminder Reminder) {
	m.Reminders.Add(reminder)
	return
}

func (m Manager) GetReminders() []Reminder {
	return m.Reminders.GetElements()
}
