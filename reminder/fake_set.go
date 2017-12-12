package reminder

type FakeSet struct {
	Set []Reminder
}

func (f *FakeSet) Add(reminder Reminder) {
	f.Set = append(f.Set, reminder)
}

func (f FakeSet) GetElements() []Reminder {
	return f.Set
}
