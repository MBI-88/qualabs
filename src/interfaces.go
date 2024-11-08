package src



type UsersInt interface {
	LoadData(path string) bool
	SolutionA() ([]byte, error)
	SolutionB() []string
}



func NewUser() UsersInt {
	return new(managerUsers)
}