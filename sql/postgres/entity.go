package postgres

type Teacher struct {
	Name string
}

type Student struct {
	Name    string
	Teacher string
}
