package reminder

type Set interface {
	Add(element string)
	GetElements() []string
}

type Manager struct {
	Reminders Set
}

func (m *Manager) Add(reminder string) {
	m.Reminders.Add(reminder)
	return
}

func (m Manager) GetReminders() []string {
	return m.Reminders.GetElements()
}
