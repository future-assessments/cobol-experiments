package mortgage

type Mortgagee struct {
	ID            string
	LastName      string
	MiddleInitial string
	FirstName     string
}

func NewMortgagee(id string, lastName string, middleInitial string, firstName string) *Mortgagee {
	return &Mortgagee{
		ID:            id,
		LastName:      lastName,
		MiddleInitial: middleInitial,
		FirstName:     firstName,
	}
}

func (mortgagee *Mortgagee) GetFullName() string {
	return mortgagee.LastName + " " + mortgagee.MiddleInitial + " " + mortgagee.FirstName
}
