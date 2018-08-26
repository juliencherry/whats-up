package reminder

type Reminder struct {
	Text string
}

type Map interface {
	Add(date string, reminder Reminder)
	GetElements() map[string][]Reminder
}

type Manager struct {
	Reminders Map
}

func (m *Manager) Add(date string, reminder Reminder) {
	m.Reminders.Add(date, reminder)
	return
}

func (m Manager) GetReminders() map[string][]Reminder {
	return m.Reminders.GetElements()
}
