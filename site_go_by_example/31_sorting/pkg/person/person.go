package person

type Person struct {
	Name    string
	Will    int
	Courage int
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Persons) Less(i, j int) bool {
	person := p[i]
	anotherPerson := p[j]

	if person.Will == anotherPerson.Will {
		return person.Courage < anotherPerson.Courage
	}

	return person.Will < anotherPerson.Will
}
