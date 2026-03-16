package person

type Person struct {
	id int // unexported field, only accessible within the package
	Name string // exported field, accessible from other packages
	Age int // exported field, accessible from other packages
}

func NewPerson(id int, name string, age int) Person {
	return Person{
		id:   id,
		Name: name,
		Age:  age,
	}
}

// method with pointer receiver can modify the original struct
func (p *Person) UpdatePerson(newName string, newAge int) {
	p.id++ // You can access and modify the unexported field 'id' within the package
	p.Name = newName
	p.Age = newAge
}