package reminder

type FakeMap struct {
	Map map[string][]Reminder
}

func (f *FakeMap) Add(date string, reminder Reminder) {
	if reminders, ok := f.Map[date]; ok {
		f.Map[date] = append(reminders, reminder)
		return
	}

	f.Map[date] = []Reminder{reminder}
}

func (f FakeMap) GetElements() map[string][]Reminder {
	return f.Map
}
