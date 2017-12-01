package reminder

type FakeSet struct {
	Set []interface{}
}

func (f *FakeSet) Add(element interface{}) {
	f.Set = append(f.Set, element)
}

func (f FakeSet) GetElements() []interface{} {
	return f.Set
}
