package reminder

type FakeSet struct {
	Set []string
}

func (f *FakeSet) Add(element string) {
	f.Set = append(f.Set, element)
}

func (f FakeSet) GetElements() []string {
	return f.Set
}
