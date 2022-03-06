package course

type Course struct {
	Name string
}

func New() Course {
	return Course{
		Name: "Golang For Beginer",
	}
}
