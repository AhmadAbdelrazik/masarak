package entity

type Person struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

func NewPerson(id, firstName, lastName, Email string) (*Person, error) {
	return &Person{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     Email,
	}, nil
}
