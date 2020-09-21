package test

type Person struct {
	Name string 	`airport-data:"allen, max=10"`
	Age int64  		`airport-data:"max=50"`
	Sex string		`airport-data:"male,female"`
	Married bool	`airport-data:"true,false"`
	It interface{}  `airport-data:""`
}
